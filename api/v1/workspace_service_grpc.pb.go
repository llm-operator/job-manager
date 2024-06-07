// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WorkspaceServiceClient is the client API for WorkspaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceServiceClient interface {
	CreateNotebook(ctx context.Context, in *CreateNotebookRequest, opts ...grpc.CallOption) (*Notebook, error)
	ListNotebooks(ctx context.Context, in *ListNotebooksRequest, opts ...grpc.CallOption) (*ListNotebooksResponse, error)
	GetNotebook(ctx context.Context, in *GetNotebookRequest, opts ...grpc.CallOption) (*Notebook, error)
	DeleteNotebook(ctx context.Context, in *DeleteNotebookRequest, opts ...grpc.CallOption) (*DeleteNotebookResponse, error)
	StopNotebook(ctx context.Context, in *StopNotebookRequest, opts ...grpc.CallOption) (*Notebook, error)
	StartNotebook(ctx context.Context, in *StartNotebookRequest, opts ...grpc.CallOption) (*Notebook, error)
}

type workspaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceServiceClient(cc grpc.ClientConnInterface) WorkspaceServiceClient {
	return &workspaceServiceClient{cc}
}

func (c *workspaceServiceClient) CreateNotebook(ctx context.Context, in *CreateNotebookRequest, opts ...grpc.CallOption) (*Notebook, error) {
	out := new(Notebook)
	err := c.cc.Invoke(ctx, "/llmoperator.workspace.server.v1.WorkspaceService/CreateNotebook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) ListNotebooks(ctx context.Context, in *ListNotebooksRequest, opts ...grpc.CallOption) (*ListNotebooksResponse, error) {
	out := new(ListNotebooksResponse)
	err := c.cc.Invoke(ctx, "/llmoperator.workspace.server.v1.WorkspaceService/ListNotebooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) GetNotebook(ctx context.Context, in *GetNotebookRequest, opts ...grpc.CallOption) (*Notebook, error) {
	out := new(Notebook)
	err := c.cc.Invoke(ctx, "/llmoperator.workspace.server.v1.WorkspaceService/GetNotebook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) DeleteNotebook(ctx context.Context, in *DeleteNotebookRequest, opts ...grpc.CallOption) (*DeleteNotebookResponse, error) {
	out := new(DeleteNotebookResponse)
	err := c.cc.Invoke(ctx, "/llmoperator.workspace.server.v1.WorkspaceService/DeleteNotebook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) StopNotebook(ctx context.Context, in *StopNotebookRequest, opts ...grpc.CallOption) (*Notebook, error) {
	out := new(Notebook)
	err := c.cc.Invoke(ctx, "/llmoperator.workspace.server.v1.WorkspaceService/StopNotebook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) StartNotebook(ctx context.Context, in *StartNotebookRequest, opts ...grpc.CallOption) (*Notebook, error) {
	out := new(Notebook)
	err := c.cc.Invoke(ctx, "/llmoperator.workspace.server.v1.WorkspaceService/StartNotebook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceServiceServer is the server API for WorkspaceService service.
// All implementations must embed UnimplementedWorkspaceServiceServer
// for forward compatibility
type WorkspaceServiceServer interface {
	CreateNotebook(context.Context, *CreateNotebookRequest) (*Notebook, error)
	ListNotebooks(context.Context, *ListNotebooksRequest) (*ListNotebooksResponse, error)
	GetNotebook(context.Context, *GetNotebookRequest) (*Notebook, error)
	DeleteNotebook(context.Context, *DeleteNotebookRequest) (*DeleteNotebookResponse, error)
	StopNotebook(context.Context, *StopNotebookRequest) (*Notebook, error)
	StartNotebook(context.Context, *StartNotebookRequest) (*Notebook, error)
	mustEmbedUnimplementedWorkspaceServiceServer()
}

// UnimplementedWorkspaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceServiceServer struct {
}

func (UnimplementedWorkspaceServiceServer) CreateNotebook(context.Context, *CreateNotebookRequest) (*Notebook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotebook not implemented")
}
func (UnimplementedWorkspaceServiceServer) ListNotebooks(context.Context, *ListNotebooksRequest) (*ListNotebooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotebooks not implemented")
}
func (UnimplementedWorkspaceServiceServer) GetNotebook(context.Context, *GetNotebookRequest) (*Notebook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotebook not implemented")
}
func (UnimplementedWorkspaceServiceServer) DeleteNotebook(context.Context, *DeleteNotebookRequest) (*DeleteNotebookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotebook not implemented")
}
func (UnimplementedWorkspaceServiceServer) StopNotebook(context.Context, *StopNotebookRequest) (*Notebook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopNotebook not implemented")
}
func (UnimplementedWorkspaceServiceServer) StartNotebook(context.Context, *StartNotebookRequest) (*Notebook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartNotebook not implemented")
}
func (UnimplementedWorkspaceServiceServer) mustEmbedUnimplementedWorkspaceServiceServer() {}

// UnsafeWorkspaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceServiceServer will
// result in compilation errors.
type UnsafeWorkspaceServiceServer interface {
	mustEmbedUnimplementedWorkspaceServiceServer()
}

func RegisterWorkspaceServiceServer(s grpc.ServiceRegistrar, srv WorkspaceServiceServer) {
	s.RegisterService(&WorkspaceService_ServiceDesc, srv)
}

func _WorkspaceService_CreateNotebook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotebookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).CreateNotebook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llmoperator.workspace.server.v1.WorkspaceService/CreateNotebook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).CreateNotebook(ctx, req.(*CreateNotebookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_ListNotebooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotebooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).ListNotebooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llmoperator.workspace.server.v1.WorkspaceService/ListNotebooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).ListNotebooks(ctx, req.(*ListNotebooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_GetNotebook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotebookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetNotebook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llmoperator.workspace.server.v1.WorkspaceService/GetNotebook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetNotebook(ctx, req.(*GetNotebookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_DeleteNotebook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotebookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).DeleteNotebook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llmoperator.workspace.server.v1.WorkspaceService/DeleteNotebook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).DeleteNotebook(ctx, req.(*DeleteNotebookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_StopNotebook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopNotebookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).StopNotebook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llmoperator.workspace.server.v1.WorkspaceService/StopNotebook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).StopNotebook(ctx, req.(*StopNotebookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_StartNotebook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartNotebookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).StartNotebook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llmoperator.workspace.server.v1.WorkspaceService/StartNotebook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).StartNotebook(ctx, req.(*StartNotebookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkspaceService_ServiceDesc is the grpc.ServiceDesc for WorkspaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkspaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "llmoperator.workspace.server.v1.WorkspaceService",
	HandlerType: (*WorkspaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNotebook",
			Handler:    _WorkspaceService_CreateNotebook_Handler,
		},
		{
			MethodName: "ListNotebooks",
			Handler:    _WorkspaceService_ListNotebooks_Handler,
		},
		{
			MethodName: "GetNotebook",
			Handler:    _WorkspaceService_GetNotebook_Handler,
		},
		{
			MethodName: "DeleteNotebook",
			Handler:    _WorkspaceService_DeleteNotebook_Handler,
		},
		{
			MethodName: "StopNotebook",
			Handler:    _WorkspaceService_StopNotebook_Handler,
		},
		{
			MethodName: "StartNotebook",
			Handler:    _WorkspaceService_StartNotebook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/workspace_service.proto",
}

// WorkspaceWorkerServiceClient is the client API for WorkspaceWorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceWorkerServiceClient interface {
	ListQueuedInternalNotebooks(ctx context.Context, in *ListQueuedInternalNotebooksRequest, opts ...grpc.CallOption) (*ListQueuedInternalNotebooksResponse, error)
	UpdateNotebookState(ctx context.Context, in *UpdateNotebookStateRequest, opts ...grpc.CallOption) (*UpdateNotebookStateResponse, error)
}

type workspaceWorkerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceWorkerServiceClient(cc grpc.ClientConnInterface) WorkspaceWorkerServiceClient {
	return &workspaceWorkerServiceClient{cc}
}

func (c *workspaceWorkerServiceClient) ListQueuedInternalNotebooks(ctx context.Context, in *ListQueuedInternalNotebooksRequest, opts ...grpc.CallOption) (*ListQueuedInternalNotebooksResponse, error) {
	out := new(ListQueuedInternalNotebooksResponse)
	err := c.cc.Invoke(ctx, "/llmoperator.workspace.server.v1.WorkspaceWorkerService/ListQueuedInternalNotebooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceWorkerServiceClient) UpdateNotebookState(ctx context.Context, in *UpdateNotebookStateRequest, opts ...grpc.CallOption) (*UpdateNotebookStateResponse, error) {
	out := new(UpdateNotebookStateResponse)
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
	ListQueuedInternalNotebooks(context.Context, *ListQueuedInternalNotebooksRequest) (*ListQueuedInternalNotebooksResponse, error)
	UpdateNotebookState(context.Context, *UpdateNotebookStateRequest) (*UpdateNotebookStateResponse, error)
	mustEmbedUnimplementedWorkspaceWorkerServiceServer()
}

// UnimplementedWorkspaceWorkerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceWorkerServiceServer struct {
}

func (UnimplementedWorkspaceWorkerServiceServer) ListQueuedInternalNotebooks(context.Context, *ListQueuedInternalNotebooksRequest) (*ListQueuedInternalNotebooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQueuedInternalNotebooks not implemented")
}
func (UnimplementedWorkspaceWorkerServiceServer) UpdateNotebookState(context.Context, *UpdateNotebookStateRequest) (*UpdateNotebookStateResponse, error) {
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
	in := new(ListQueuedInternalNotebooksRequest)
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
		return srv.(WorkspaceWorkerServiceServer).ListQueuedInternalNotebooks(ctx, req.(*ListQueuedInternalNotebooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceWorkerService_UpdateNotebookState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotebookStateRequest)
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
		return srv.(WorkspaceWorkerServiceServer).UpdateNotebookState(ctx, req.(*UpdateNotebookStateRequest))
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
	Metadata: "api/v1/workspace_service.proto",
}
