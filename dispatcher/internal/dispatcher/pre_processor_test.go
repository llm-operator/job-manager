package dispatcher

import (
	"context"
	"fmt"
	"testing"
	"time"

	fv1 "github.com/llm-operator/file-manager/api/v1"
	v1 "github.com/llm-operator/job-manager/api/v1"
	"github.com/llm-operator/job-manager/common/pkg/store"
	mv1 "github.com/llm-operator/model-manager/api/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func TestPreProcess(t *testing.T) {
	fc := &fakeFileClient{
		id: "training-file-id",
	}
	mc := &fakeModelClient{
		id: "model-id",
	}
	sc := &fakeS3Client{}

	p := NewPreProcessor(fc, mc, sc)

	jobProto := &v1.Job{
		Model:        "model-id",
		TrainingFile: "training-file-id",
	}
	b, err := proto.Marshal(jobProto)
	assert.NoError(t, err)

	job := &store.Job{
		JobID:   "job-id",
		Message: b,
	}

	got, err := p.Process(context.Background(), job)
	assert.NoError(t, err)
	want := &PreProcessResult{
		BaseModelURL:    "presigned-model-path",
		TrainingFileURL: "presigned-file-path",
		OutputModelID:   "generated-model-id",
		OutputModelURL:  "presigned-generated-model-path",
	}
	assert.Equal(t, want, got)
}

type fakeFileClient struct {
	id string
}

func (f *fakeFileClient) GetFilePath(ctx context.Context, in *fv1.GetFilePathRequest, opts ...grpc.CallOption) (*fv1.GetFilePathResponse, error) {
	if in.Id != f.id {
		return nil, fmt.Errorf("unexpected id: %s", in.Id)
	}

	return &fv1.GetFilePathResponse{
		Path: "file-path",
	}, nil
}

type fakeModelClient struct {
	id string
}

func (f *fakeModelClient) RegisterModel(ctx context.Context, in *mv1.RegisterModelRequest, opts ...grpc.CallOption) (*mv1.RegisterModelResponse, error) {
	return &mv1.RegisterModelResponse{
		Id:   "generated-model-id",
		Path: "generated-model-path",
	}, nil
}

func (f *fakeModelClient) GetModelPath(ctx context.Context, in *mv1.GetModelPathRequest, opts ...grpc.CallOption) (*mv1.GetModelPathResponse, error) {
	if in.Id != f.id {
		return nil, fmt.Errorf("unexpected id: %s", in.Id)
	}

	return &mv1.GetModelPathResponse{
		Path: "model-path",
	}, nil
}

type fakeS3Client struct {
}

func (f *fakeS3Client) GeneratePresignedURL(key string, expire time.Duration) (string, error) {
	return fmt.Sprintf("presigned-%s", key), nil
}