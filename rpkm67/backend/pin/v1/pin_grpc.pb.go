// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: rpkm67/backend/pin/v1/pin.proto

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
	PinService_FindAll_FullMethodName  = "/rpkm67.backend.pin.v1.PinService/FindAll"
	PinService_ResetPin_FullMethodName = "/rpkm67.backend.pin.v1.PinService/ResetPin"
)

// PinServiceClient is the client API for PinService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PinServiceClient interface {
	FindAll(ctx context.Context, in *FindAllPinRequest, opts ...grpc.CallOption) (*FindAllPinResponse, error)
	ResetPin(ctx context.Context, in *ResetPinRequest, opts ...grpc.CallOption) (*ResetPinResponse, error)
}

type pinServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPinServiceClient(cc grpc.ClientConnInterface) PinServiceClient {
	return &pinServiceClient{cc}
}

func (c *pinServiceClient) FindAll(ctx context.Context, in *FindAllPinRequest, opts ...grpc.CallOption) (*FindAllPinResponse, error) {
	out := new(FindAllPinResponse)
	err := c.cc.Invoke(ctx, PinService_FindAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pinServiceClient) ResetPin(ctx context.Context, in *ResetPinRequest, opts ...grpc.CallOption) (*ResetPinResponse, error) {
	out := new(ResetPinResponse)
	err := c.cc.Invoke(ctx, PinService_ResetPin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PinServiceServer is the server API for PinService service.
// All implementations must embed UnimplementedPinServiceServer
// for forward compatibility
type PinServiceServer interface {
	FindAll(context.Context, *FindAllPinRequest) (*FindAllPinResponse, error)
	ResetPin(context.Context, *ResetPinRequest) (*ResetPinResponse, error)
	mustEmbedUnimplementedPinServiceServer()
}

// UnimplementedPinServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPinServiceServer struct {
}

func (UnimplementedPinServiceServer) FindAll(context.Context, *FindAllPinRequest) (*FindAllPinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedPinServiceServer) ResetPin(context.Context, *ResetPinRequest) (*ResetPinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPin not implemented")
}
func (UnimplementedPinServiceServer) mustEmbedUnimplementedPinServiceServer() {}

// UnsafePinServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PinServiceServer will
// result in compilation errors.
type UnsafePinServiceServer interface {
	mustEmbedUnimplementedPinServiceServer()
}

func RegisterPinServiceServer(s grpc.ServiceRegistrar, srv PinServiceServer) {
	s.RegisterService(&PinService_ServiceDesc, srv)
}

func _PinService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllPinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PinServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PinService_FindAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PinServiceServer).FindAll(ctx, req.(*FindAllPinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PinService_ResetPin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PinServiceServer).ResetPin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PinService_ResetPin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PinServiceServer).ResetPin(ctx, req.(*ResetPinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PinService_ServiceDesc is the grpc.ServiceDesc for PinService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PinService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpkm67.backend.pin.v1.PinService",
	HandlerType: (*PinServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAll",
			Handler:    _PinService_FindAll_Handler,
		},
		{
			MethodName: "ResetPin",
			Handler:    _PinService_ResetPin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpkm67/backend/pin/v1/pin.proto",
}