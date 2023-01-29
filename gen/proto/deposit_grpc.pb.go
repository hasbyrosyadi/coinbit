// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// DepositServiceClient is the client API for DepositService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepositServiceClient interface {
	Deposit(ctx context.Context, in *PostDeposit, opts ...grpc.CallOption) (*ResponsePostDeposit, error)
	GetDeposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*ResponseGetDeposit, error)
}

type depositServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepositServiceClient(cc grpc.ClientConnInterface) DepositServiceClient {
	return &depositServiceClient{cc}
}

func (c *depositServiceClient) Deposit(ctx context.Context, in *PostDeposit, opts ...grpc.CallOption) (*ResponsePostDeposit, error) {
	out := new(ResponsePostDeposit)
	err := c.cc.Invoke(ctx, "/proto.DepositService/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depositServiceClient) GetDeposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*ResponseGetDeposit, error) {
	out := new(ResponseGetDeposit)
	err := c.cc.Invoke(ctx, "/proto.DepositService/GetDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepositServiceServer is the server API for DepositService service.
// All implementations must embed UnimplementedDepositServiceServer
// for forward compatibility
type DepositServiceServer interface {
	Deposit(context.Context, *PostDeposit) (*ResponsePostDeposit, error)
	GetDeposit(context.Context, *DepositRequest) (*ResponseGetDeposit, error)
	mustEmbedUnimplementedDepositServiceServer()
}

// UnimplementedDepositServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDepositServiceServer struct {
}

func (UnimplementedDepositServiceServer) Deposit(context.Context, *PostDeposit) (*ResponsePostDeposit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedDepositServiceServer) GetDeposit(context.Context, *DepositRequest) (*ResponseGetDeposit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeposit not implemented")
}
func (UnimplementedDepositServiceServer) mustEmbedUnimplementedDepositServiceServer() {}

// UnsafeDepositServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepositServiceServer will
// result in compilation errors.
type UnsafeDepositServiceServer interface {
	mustEmbedUnimplementedDepositServiceServer()
}

func RegisterDepositServiceServer(s grpc.ServiceRegistrar, srv DepositServiceServer) {
	s.RegisterService(&DepositService_ServiceDesc, srv)
}

func _DepositService_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostDeposit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepositServiceServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DepositService/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepositServiceServer).Deposit(ctx, req.(*PostDeposit))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepositService_GetDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepositServiceServer).GetDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DepositService/GetDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepositServiceServer).GetDeposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DepositService_ServiceDesc is the grpc.ServiceDesc for DepositService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepositService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DepositService",
	HandlerType: (*DepositServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deposit",
			Handler:    _DepositService_Deposit_Handler,
		},
		{
			MethodName: "GetDeposit",
			Handler:    _DepositService_GetDeposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deposit.proto",
}
