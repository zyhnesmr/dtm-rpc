// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: busis.proto

package busis

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
	Busis_AddMoney_FullMethodName = "/busis.Busis/AddMoney"
	Busis_DelMoney_FullMethodName = "/busis.Busis/DelMoney"
)

// BusisClient is the client API for Busis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BusisClient interface {
	AddMoney(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*Empty, error)
	DelMoney(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*Empty, error)
}

type busisClient struct {
	cc grpc.ClientConnInterface
}

func NewBusisClient(cc grpc.ClientConnInterface) BusisClient {
	return &busisClient{cc}
}

func (c *busisClient) AddMoney(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Busis_AddMoney_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busisClient) DelMoney(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Busis_DelMoney_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusisServer is the server API for Busis service.
// All implementations must embed UnimplementedBusisServer
// for forward compatibility
type BusisServer interface {
	AddMoney(context.Context, *AddReq) (*Empty, error)
	DelMoney(context.Context, *AddReq) (*Empty, error)
	mustEmbedUnimplementedBusisServer()
}

// UnimplementedBusisServer must be embedded to have forward compatible implementations.
type UnimplementedBusisServer struct {
}

func (UnimplementedBusisServer) AddMoney(context.Context, *AddReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMoney not implemented")
}
func (UnimplementedBusisServer) DelMoney(context.Context, *AddReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelMoney not implemented")
}
func (UnimplementedBusisServer) mustEmbedUnimplementedBusisServer() {}

// UnsafeBusisServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusisServer will
// result in compilation errors.
type UnsafeBusisServer interface {
	mustEmbedUnimplementedBusisServer()
}

func RegisterBusisServer(s grpc.ServiceRegistrar, srv BusisServer) {
	s.RegisterService(&Busis_ServiceDesc, srv)
}

func _Busis_AddMoney_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusisServer).AddMoney(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Busis_AddMoney_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusisServer).AddMoney(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Busis_DelMoney_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusisServer).DelMoney(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Busis_DelMoney_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusisServer).DelMoney(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Busis_ServiceDesc is the grpc.ServiceDesc for Busis service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Busis_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "busis.Busis",
	HandlerType: (*BusisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMoney",
			Handler:    _Busis_AddMoney_Handler,
		},
		{
			MethodName: "DelMoney",
			Handler:    _Busis_DelMoney_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "busis.proto",
}
