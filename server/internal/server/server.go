package server

import (
	"context"
	"fmt"
	"log"
	"net"

	fv1 "github.com/llm-operator/file-manager/api/v1"
	v1 "github.com/llm-operator/job-manager/api/v1"
	"github.com/llm-operator/job-manager/common/pkg/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type fileGetClient interface {
	GetFile(ctx context.Context, in *fv1.GetFileRequest, opts ...grpc.CallOption) (*fv1.File, error)
}

type k8sJobClient interface {
	CancelJob(ctx context.Context, jobID string) error
}

// New creates a server.
func New(store *store.S, fileGetClient fileGetClient, k8sJobClient k8sJobClient) *S {
	return &S{
		store:         store,
		fileGetClient: fileGetClient,
		k8sJobClient:  k8sJobClient,
	}
}

// S is a server.
type S struct {
	v1.UnimplementedFineTuningServiceServer

	srv *grpc.Server

	store         *store.S
	fileGetClient fileGetClient
	k8sJobClient  k8sJobClient
}

// Run starts the gRPC server.
func (s *S) Run(port int) error {
	log.Printf("Starting server on port %d\n", port)

	grpcServer := grpc.NewServer()
	v1.RegisterFineTuningServiceServer(grpcServer, s)
	reflection.Register(grpcServer)

	s.srv = grpcServer

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("listen: %s", err)
	}
	if err := grpcServer.Serve(l); err != nil {
		return fmt.Errorf("serve: %s", err)
	}
	return nil
}

// Stop stops the gRPC server.
func (s *S) Stop() {
	s.srv.Stop()
}
