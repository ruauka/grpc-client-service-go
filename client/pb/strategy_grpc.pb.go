// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: strategy.proto

package pb

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

// StrategyClient is the client API for Strategy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StrategyClient interface {
	//  gRPC service health check
	HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// main rpc
	Execute(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type strategyClient struct {
	cc grpc.ClientConnInterface
}

func NewStrategyClient(cc grpc.ClientConnInterface) StrategyClient {
	return &strategyClient{cc}
}

func (c *strategyClient) HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/strategy.Strategy/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *strategyClient) Execute(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/strategy.Strategy/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StrategyServer is the server API for Strategy service.
// All implementations must embed UnimplementedStrategyServer
// for forward compatibility
type StrategyServer interface {
	//  gRPC service health check
	HealthCheck(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// main rpc
	Execute(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedStrategyServer()
}

// UnimplementedStrategyServer must be embedded to have forward compatible implementations.
type UnimplementedStrategyServer struct {
}

func (UnimplementedStrategyServer) HealthCheck(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedStrategyServer) Execute(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedStrategyServer) mustEmbedUnimplementedStrategyServer() {}

// UnsafeStrategyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StrategyServer will
// result in compilation errors.
type UnsafeStrategyServer interface {
	mustEmbedUnimplementedStrategyServer()
}

func RegisterStrategyServer(s grpc.ServiceRegistrar, srv StrategyServer) {
	s.RegisterService(&Strategy_ServiceDesc, srv)
}

func _Strategy_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strategy.Strategy/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyServer).HealthCheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Strategy_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StrategyServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strategy.Strategy/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StrategyServer).Execute(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Strategy_ServiceDesc is the grpc.ServiceDesc for Strategy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Strategy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strategy.Strategy",
	HandlerType: (*StrategyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Strategy_HealthCheck_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _Strategy_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strategy.proto",
}
