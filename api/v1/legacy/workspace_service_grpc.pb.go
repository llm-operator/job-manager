// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package legacy

import (
	context "context"
	v1 "github.com/llm-operator/job-manager/api/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WorkspaceWorkerServiceClient is the client API for WorkspaceWorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceWorkerServiceClient interface {
	ListQueuedInternalNotebooks(ctx context.Context, in *v1.ListQueuedInternalNotebooksRequest, opts ...grpc.CallOption) (*v1.ListQueuedInternalNotebooksResponse, error)
	UpdateNotebookState(ctx context.Context, in *v1.UpdateNotebookStateRequest, opts ...grpc.CallOption) (*v1.UpdateNotebookStateResponse, error)
}

type workspaceWorkerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceWorkerServiceClient(cc grpc.ClientConnInterface) WorkspaceWorkerServiceClient {
	return &workspaceWorkerServiceClient{cc}
}

func (c *workspaceWorkerServiceClient) ListQueuedInternalNotebooks(ctx context.Context, in *v1.ListQueuedInternalNotebooksRequest, opts ...grpc.CallOption) (*v1.ListQueuedInternalNotebooksResponse, error) {
	out := new(v1.ListQueuedInternalNotebooksResponse)
	err := c.cc.Invoke(ctx, "/llmoperator.workspace.server.v1.WorkspaceWorkerService/ListQueuedInternalNotebooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceWorkerServiceClient) UpdateNotebookState(ctx context.Context, in *v1.UpdateNotebookStateRequest, opts ...grpc.CallOption) (*v1.UpdateNotebookStateResponse, error) {
	out := new(v1.UpdateNotebookStateResponse)
	err := c.cc.Invoke(ctx, "/llmoperator.workspace.server.v1.WorkspaceWorkerService/UpdateNotebookState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceWorkerServiceServer is the server API for WorkspaceWorkerService service.
// All implementations must embed UnimplementedWorkspaceWorkerServiceServer
// for forward compatibility
type WorkspaceWorkerServiceServer interface {
	ListQueuedInternalNotebooks(context.Context, *v1.ListQueuedInternalNotebooksRequest) (*v1.ListQueuedInternalNotebooksResponse, error)
	UpdateNotebookState(context.Context, *v1.UpdateNotebookStateRequest) (*v1.UpdateNotebookStateResponse, error)
	mustEmbedUnimplementedWorkspaceWorkerServiceServer()
}

// UnimplementedWorkspaceWorkerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceWorkerServiceServer struct {
}

func (UnimplementedWorkspaceWorkerServiceServer) ListQueuedInternalNotebooks(context.Context, *v1.ListQueuedInternalNotebooksRequest) (*v1.ListQueuedInternalNotebooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQueuedInternalNotebooks not implemented")
}
func (UnimplementedWorkspaceWorkerServiceServer) UpdateNotebookState(context.Context, *v1.UpdateNotebookStateRequest) (*v1.UpdateNotebookStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotebookState not implemented")
}
func (UnimplementedWorkspaceWorkerServiceServer) mustEmbedUnimplementedWorkspaceWorkerServiceServer() {
}

// UnsafeWorkspaceWorkerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceWorkerServiceServer will
// result in compilation errors.
type UnsafeWorkspaceWorkerServiceServer interface {
	mustEmbedUnimplementedWorkspaceWorkerServiceServer()
}

func RegisterWorkspaceWorkerServiceServer(s grpc.ServiceRegistrar, srv WorkspaceWorkerServiceServer) {
	s.RegisterService(&WorkspaceWorkerService_ServiceDesc, srv)
}

func _WorkspaceWorkerService_ListQueuedInternalNotebooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ListQueuedInternalNotebooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceWorkerServiceServer).ListQueuedInternalNotebooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llmoperator.workspace.server.v1.WorkspaceWorkerService/ListQueuedInternalNotebooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceWorkerServiceServer).ListQueuedInternalNotebooks(ctx, req.(*v1.ListQueuedInternalNotebooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceWorkerService_UpdateNotebookState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdateNotebookStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceWorkerServiceServer).UpdateNotebookState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llmoperator.workspace.server.v1.WorkspaceWorkerService/UpdateNotebookState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceWorkerServiceServer).UpdateNotebookState(ctx, req.(*v1.UpdateNotebookStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkspaceWorkerService_ServiceDesc is the grpc.ServiceDesc for WorkspaceWorkerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkspaceWorkerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "llmoperator.workspace.server.v1.WorkspaceWorkerService",
	HandlerType: (*WorkspaceWorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListQueuedInternalNotebooks",
			Handler:    _WorkspaceWorkerService_ListQueuedInternalNotebooks_Handler,
		},
		{
			MethodName: "UpdateNotebookState",
			Handler:    _WorkspaceWorkerService_UpdateNotebookState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/legacy/workspace_service.proto",
}
