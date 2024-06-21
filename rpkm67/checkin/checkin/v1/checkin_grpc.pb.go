// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: rpkm67/checkin/checkin/v1/checkin.proto

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
	CheckInService_Create_FullMethodName          = "/rpkm67.checkin.checkin.v1.CheckInService/Create"
	CheckInService_FindByStudentId_FullMethodName = "/rpkm67.checkin.checkin.v1.CheckInService/FindByStudentId"
)

// CheckInServiceClient is the client API for CheckInService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckInServiceClient interface {
	Create(ctx context.Context, in *CreateCheckInRequest, opts ...grpc.CallOption) (*CreateCheckInResponse, error)
	FindByStudentId(ctx context.Context, in *FindByStudentIdCheckInRequest, opts ...grpc.CallOption) (*FindByStudentIdCheckInResponse, error)
}

type checkInServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckInServiceClient(cc grpc.ClientConnInterface) CheckInServiceClient {
	return &checkInServiceClient{cc}
}

func (c *checkInServiceClient) Create(ctx context.Context, in *CreateCheckInRequest, opts ...grpc.CallOption) (*CreateCheckInResponse, error) {
	out := new(CreateCheckInResponse)
	err := c.cc.Invoke(ctx, CheckInService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkInServiceClient) FindByStudentId(ctx context.Context, in *FindByStudentIdCheckInRequest, opts ...grpc.CallOption) (*FindByStudentIdCheckInResponse, error) {
	out := new(FindByStudentIdCheckInResponse)
	err := c.cc.Invoke(ctx, CheckInService_FindByStudentId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckInServiceServer is the server API for CheckInService service.
// All implementations must embed UnimplementedCheckInServiceServer
// for forward compatibility
type CheckInServiceServer interface {
	Create(context.Context, *CreateCheckInRequest) (*CreateCheckInResponse, error)
	FindByStudentId(context.Context, *FindByStudentIdCheckInRequest) (*FindByStudentIdCheckInResponse, error)
	mustEmbedUnimplementedCheckInServiceServer()
}

// UnimplementedCheckInServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCheckInServiceServer struct {
}

func (UnimplementedCheckInServiceServer) Create(context.Context, *CreateCheckInRequest) (*CreateCheckInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCheckInServiceServer) FindByStudentId(context.Context, *FindByStudentIdCheckInRequest) (*FindByStudentIdCheckInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByStudentId not implemented")
}
func (UnimplementedCheckInServiceServer) mustEmbedUnimplementedCheckInServiceServer() {}

// UnsafeCheckInServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckInServiceServer will
// result in compilation errors.
type UnsafeCheckInServiceServer interface {
	mustEmbedUnimplementedCheckInServiceServer()
}

func RegisterCheckInServiceServer(s grpc.ServiceRegistrar, srv CheckInServiceServer) {
	s.RegisterService(&CheckInService_ServiceDesc, srv)
}

func _CheckInService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCheckInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckInServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckInService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckInServiceServer).Create(ctx, req.(*CreateCheckInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckInService_FindByStudentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByStudentIdCheckInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckInServiceServer).FindByStudentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckInService_FindByStudentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckInServiceServer).FindByStudentId(ctx, req.(*FindByStudentIdCheckInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckInService_ServiceDesc is the grpc.ServiceDesc for CheckInService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckInService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpkm67.checkin.checkin.v1.CheckInService",
	HandlerType: (*CheckInServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CheckInService_Create_Handler,
		},
		{
			MethodName: "FindByStudentId",
			Handler:    _CheckInService_FindByStudentId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpkm67/checkin/checkin/v1/checkin.proto",
}
