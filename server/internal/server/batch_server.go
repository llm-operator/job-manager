package server

import (
	"context"
	"errors"
	"time"

	"github.com/llm-operator/common/pkg/id"
	v1 "github.com/llm-operator/job-manager/api/v1"
	"github.com/llm-operator/job-manager/server/internal/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

// CreateBatchJob creates a batch job.
func (s *S) CreateBatchJob(ctx context.Context, req *v1.CreateBatchJobRequest) (*v1.BatchJob, error) {
	userInfo, err := s.extractUserInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Image == "" {
		return nil, status.Error(codes.InvalidArgument, "image is required")
	}
	image, ok := s.batchJobImages[req.Image]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "invalid image: %s", req.Image)
	}

	if req.Command == "" {
		return nil, status.Error(codes.InvalidArgument, "command is required")
	}

	if len(req.Scripts) == 0 {
		return nil, status.Error(codes.InvalidArgument, "scripts are required")
	}

	for _, fileID := range req.DataFiles {
		if err := s.validateFile(ctx, fileID); err != nil {
			return nil, err
		}
	}

	jobID, err := id.GenerateIDForK8SResource("bj-")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "generate notebook id: %s", err)
	}

	if len(userInfo.AssignedKubernetesEnvs) == 0 {
		return nil, status.Errorf(codes.Internal, "no kuberentes cluster/namespace for a job")
	}
	// TODO(kenji): Revisit. We might want dispatcher to pick up a cluster/namespace.
	kenv := userInfo.AssignedKubernetesEnvs[0]

	jobProto := &v1.BatchJob{
		Id:        jobID,
		CreatedAt: time.Now().UTC().Unix(),
		Status:    string(store.BatchJobStateQueued),
		Image:     image,
		Command:   req.Command,
		Resources: req.Resources,
		Envs:      req.Envs,
		DataFiles: req.DataFiles,
	}
	msg, err := proto.Marshal(jobProto)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "marshal batch job: %s", err)
	}

	apikey, err := s.extractTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}
	kclient, err := s.k8sClientFactory.NewClient(kenv, apikey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create k8s client: %s", err)
	}
	if err := kclient.CreateSecret(ctx, jobID, kenv.Namespace, map[string][]byte{
		"OPENAI_API_KEY": []byte(apikey),
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "create secret: %s", err)
	}
	if err := kclient.CreateConfigMap(ctx, jobID, kenv.Namespace, req.Scripts); err != nil {
		return nil, status.Errorf(codes.Internal, "create configmap for scripts: %s", err)
	}

	job := &store.BatchJob{
		JobID:               jobID,
		Image:               image,
		Message:             msg,
		State:               store.BatchJobStateQueued,
		QueuedAction:        store.BatchJobQueuedActionCreate,
		TenantID:            userInfo.TenantID,
		OrganizationID:      userInfo.OrganizationID,
		ProjectID:           userInfo.ProjectID,
		KubernetesNamespace: kenv.Namespace,
	}
	if err := s.store.CreateBatchJob(job); err != nil {
		return nil, status.Errorf(codes.Internal, "create batch job: %s", err)
	}
	return jobProto, nil
}

// ListBatchJobs lists batch jobs.
func (s *S) ListBatchJobs(ctx context.Context, req *v1.ListBatchJobsRequest) (*v1.ListBatchJobsResponse, error) {
	userInfo, err := s.extractUserInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Limit < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "limit must be non-negative")
	}
	limit := req.Limit
	if limit == 0 {
		limit = defaultPageSize
	}
	if limit > maxPageSize {
		limit = maxPageSize
	}

	var after uint
	if req.After != "" {
		job, err := s.store.GetBatchJobByIDAndProjectID(req.After, userInfo.ProjectID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, status.Errorf(codes.InvalidArgument, "invalid after: %s", err)
			}
			return nil, status.Errorf(codes.Internal, "get batch job: %s", err)
		}
		after = job.ID
	}

	jobs, hasMore, err := s.store.ListBatchJobsByProjectIDWithPagination(userInfo.ProjectID, after, int(limit))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list batch jobs: %s", err)
	}

	var jobProtos []*v1.BatchJob
	for _, job := range jobs {
		jobProto, err := job.V1BatchJob()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "convert batch job to proto: %s", err)
		}
		jobProtos = append(jobProtos, jobProto)
	}
	return &v1.ListBatchJobsResponse{
		Jobs:    jobProtos,
		HasMore: hasMore,
	}, nil
}

// GetBatchJob gets a batch job.
func (s *S) GetBatchJob(ctx context.Context, req *v1.GetBatchJobRequest) (*v1.BatchJob, error) {
	userInfo, err := s.extractUserInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	job, err := s.store.GetBatchJobByIDAndProjectID(req.Id, userInfo.ProjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "get batch job: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get batch job: %s", err)
	}

	jobProto, err := job.V1BatchJob()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "convert batch job to proto: %s", err)
	}
	return jobProto, nil
}

// CancelBatchJob cancels a batch job.
func (s *S) CancelBatchJob(ctx context.Context, req *v1.CancelBatchJobRequest) (*v1.BatchJob, error) {
	userInfo, err := s.extractUserInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	job, err := s.store.GetBatchJobByIDAndProjectID(req.Id, userInfo.ProjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "get batch job: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get batch job: %s", err)
	}

	jobProto, err := job.V1BatchJob()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "convert batch job to proto: %s", err)
	}

	switch job.State {
	case store.BatchJobStateFailed,
		store.BatchJobStateCanceled,
		store.BatchJobStateSucceeded:
		return jobProto, nil
	case store.BatchJobStateRunning:
	case store.BatchJobStateQueued:
		if job.QueuedAction == store.BatchJobQueuedActionCancel {
			return jobProto, nil
		}
	default:
		return nil, status.Errorf(codes.Internal, "unknown batch job state: %s", job.State)
	}

	job, err = s.store.SetBatchJobQueuedAction(job.JobID, job.Version, store.BatchJobQueuedActionCancel)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "set batch job queued action: %s", err)
	}
	jobProto, err = job.V1BatchJob()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "convert batch job to proto: %s", err)
	}
	return jobProto, nil
}
