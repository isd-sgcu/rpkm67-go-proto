// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: rpkm67/backend/stamp/v1/stamp.proto

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

const (
	StampService_FindByUserId_FullMethodName  = "/rpkm67.backend.stamp.v1.StampService/FindByUserId"
	StampService_StampByUserId_FullMethodName = "/rpkm67.backend.stamp.v1.StampService/StampByUserId"
)

// StampServiceClient is the client API for StampService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StampServiceClient interface {
	FindByUserId(ctx context.Context, in *FindByUserIdStampRequest, opts ...grpc.CallOption) (*FindByUserIdStampResponse, error)
	StampByUserId(ctx context.Context, in *StampByUserIdRequest, opts ...grpc.CallOption) (*StampByUserIdResponse, error)
}

type stampServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStampServiceClient(cc grpc.ClientConnInterface) StampServiceClient {
	return &stampServiceClient{cc}
}

func (c *stampServiceClient) FindByUserId(ctx context.Context, in *FindByUserIdStampRequest, opts ...grpc.CallOption) (*FindByUserIdStampResponse, error) {
	out := new(FindByUserIdStampResponse)
	err := c.cc.Invoke(ctx, StampService_FindByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stampServiceClient) StampByUserId(ctx context.Context, in *StampByUserIdRequest, opts ...grpc.CallOption) (*StampByUserIdResponse, error) {
	out := new(StampByUserIdResponse)
	err := c.cc.Invoke(ctx, StampService_StampByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StampServiceServer is the server API for StampService service.
// All implementations must embed UnimplementedStampServiceServer
// for forward compatibility
type StampServiceServer interface {
	FindByUserId(context.Context, *FindByUserIdStampRequest) (*FindByUserIdStampResponse, error)
	StampByUserId(context.Context, *StampByUserIdRequest) (*StampByUserIdResponse, error)
	mustEmbedUnimplementedStampServiceServer()
}

// UnimplementedStampServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStampServiceServer struct {
}

func (UnimplementedStampServiceServer) FindByUserId(context.Context, *FindByUserIdStampRequest) (*FindByUserIdStampResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByUserId not implemented")
}
func (UnimplementedStampServiceServer) StampByUserId(context.Context, *StampByUserIdRequest) (*StampByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StampByUserId not implemented")
}
func (UnimplementedStampServiceServer) mustEmbedUnimplementedStampServiceServer() {}

// UnsafeStampServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StampServiceServer will
// result in compilation errors.
type UnsafeStampServiceServer interface {
	mustEmbedUnimplementedStampServiceServer()
}

func RegisterStampServiceServer(s grpc.ServiceRegistrar, srv StampServiceServer) {
	s.RegisterService(&StampService_ServiceDesc, srv)
}

func _StampService_FindByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByUserIdStampRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StampServiceServer).FindByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StampService_FindByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StampServiceServer).FindByUserId(ctx, req.(*FindByUserIdStampRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StampService_StampByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StampByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StampServiceServer).StampByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StampService_StampByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StampServiceServer).StampByUserId(ctx, req.(*StampByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StampService_ServiceDesc is the grpc.ServiceDesc for StampService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StampService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpkm67.backend.stamp.v1.StampService",
	HandlerType: (*StampServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByUserId",
			Handler:    _StampService_FindByUserId_Handler,
		},
		{
			MethodName: "StampByUserId",
			Handler:    _StampService_StampByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpkm67/backend/stamp/v1/stamp.proto",
}
