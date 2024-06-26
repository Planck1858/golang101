// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GrpcService_CreateTodo_FullMethodName                 = "/grpc_server.GrpcService/CreateTodo"
	GrpcService_CreateTodoV2_FullMethodName               = "/grpc_server.GrpcService/CreateTodoV2"
	GrpcService_GetAllTodo_FullMethodName                 = "/grpc_server.GrpcService/GetAllTodo"
	GrpcService_GetTodoByID_FullMethodName                = "/grpc_server.GrpcService/GetTodoByID"
	GrpcService_GetTodoByIDWithQueryParams_FullMethodName = "/grpc_server.GrpcService/GetTodoByIDWithQueryParams"
)

// GrpcServiceClient is the client API for GrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpcServiceClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*Todo, error)
	CreateTodoV2(ctx context.Context, in *CreateTodoRequestV2, opts ...grpc.CallOption) (*Todo, error)
	GetAllTodo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TodoList, error)
	GetTodoByID(ctx context.Context, in *GetTodoByIDRequest, opts ...grpc.CallOption) (*Todo, error)
	GetTodoByIDWithQueryParams(ctx context.Context, in *GetTodoByIDRequestWithQueryParams, opts ...grpc.CallOption) (*TodoList, error)
}

type grpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcServiceClient(cc grpc.ClientConnInterface) GrpcServiceClient {
	return &grpcServiceClient{cc}
}

func (c *grpcServiceClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, GrpcService_CreateTodo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcServiceClient) CreateTodoV2(ctx context.Context, in *CreateTodoRequestV2, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, GrpcService_CreateTodoV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcServiceClient) GetAllTodo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TodoList, error) {
	out := new(TodoList)
	err := c.cc.Invoke(ctx, GrpcService_GetAllTodo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcServiceClient) GetTodoByID(ctx context.Context, in *GetTodoByIDRequest, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, GrpcService_GetTodoByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcServiceClient) GetTodoByIDWithQueryParams(ctx context.Context, in *GetTodoByIDRequestWithQueryParams, opts ...grpc.CallOption) (*TodoList, error) {
	out := new(TodoList)
	err := c.cc.Invoke(ctx, GrpcService_GetTodoByIDWithQueryParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcServiceServer is the server API for GrpcService service.
// All implementations must embed UnimplementedGrpcServiceServer
// for forward compatibility
type GrpcServiceServer interface {
	CreateTodo(context.Context, *CreateTodoRequest) (*Todo, error)
	CreateTodoV2(context.Context, *CreateTodoRequestV2) (*Todo, error)
	GetAllTodo(context.Context, *emptypb.Empty) (*TodoList, error)
	GetTodoByID(context.Context, *GetTodoByIDRequest) (*Todo, error)
	GetTodoByIDWithQueryParams(context.Context, *GetTodoByIDRequestWithQueryParams) (*TodoList, error)
	mustEmbedUnimplementedGrpcServiceServer()
}

// UnimplementedGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGrpcServiceServer struct {
}

func (UnimplementedGrpcServiceServer) CreateTodo(context.Context, *CreateTodoRequest) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedGrpcServiceServer) CreateTodoV2(context.Context, *CreateTodoRequestV2) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodoV2 not implemented")
}
func (UnimplementedGrpcServiceServer) GetAllTodo(context.Context, *emptypb.Empty) (*TodoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTodo not implemented")
}
func (UnimplementedGrpcServiceServer) GetTodoByID(context.Context, *GetTodoByIDRequest) (*Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoByID not implemented")
}
func (UnimplementedGrpcServiceServer) GetTodoByIDWithQueryParams(context.Context, *GetTodoByIDRequestWithQueryParams) (*TodoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoByIDWithQueryParams not implemented")
}
func (UnimplementedGrpcServiceServer) mustEmbedUnimplementedGrpcServiceServer() {}

// UnsafeGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcServiceServer will
// result in compilation errors.
type UnsafeGrpcServiceServer interface {
	mustEmbedUnimplementedGrpcServiceServer()
}

func RegisterGrpcServiceServer(s grpc.ServiceRegistrar, srv GrpcServiceServer) {
	s.RegisterService(&GrpcService_ServiceDesc, srv)
}

func _GrpcService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GrpcService_CreateTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServiceServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcService_CreateTodoV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequestV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServiceServer).CreateTodoV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GrpcService_CreateTodoV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServiceServer).CreateTodoV2(ctx, req.(*CreateTodoRequestV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcService_GetAllTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServiceServer).GetAllTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GrpcService_GetAllTodo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServiceServer).GetAllTodo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcService_GetTodoByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServiceServer).GetTodoByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GrpcService_GetTodoByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServiceServer).GetTodoByID(ctx, req.(*GetTodoByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcService_GetTodoByIDWithQueryParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoByIDRequestWithQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServiceServer).GetTodoByIDWithQueryParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GrpcService_GetTodoByIDWithQueryParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServiceServer).GetTodoByIDWithQueryParams(ctx, req.(*GetTodoByIDRequestWithQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

// GrpcService_ServiceDesc is the grpc.ServiceDesc for GrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_server.GrpcService",
	HandlerType: (*GrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _GrpcService_CreateTodo_Handler,
		},
		{
			MethodName: "CreateTodoV2",
			Handler:    _GrpcService_CreateTodoV2_Handler,
		},
		{
			MethodName: "GetAllTodo",
			Handler:    _GrpcService_GetAllTodo_Handler,
		},
		{
			MethodName: "GetTodoByID",
			Handler:    _GrpcService_GetTodoByID_Handler,
		},
		{
			MethodName: "GetTodoByIDWithQueryParams",
			Handler:    _GrpcService_GetTodoByIDWithQueryParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
