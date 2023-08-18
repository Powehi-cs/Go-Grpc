// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.0
// source: hello.proto

package helloproto

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

// HelloServerClient is the client API for HelloServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServerClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type helloServerClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServerClient(cc grpc.ClientConnInterface) HelloServerClient {
	return &helloServerClient{cc}
}

func (c *helloServerClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.HelloServer/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServerServer is the server API for HelloServer service.
// All implementations must embed UnimplementedHelloServerServer
// for forward compatibility
type HelloServerServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedHelloServerServer()
}

// UnimplementedHelloServerServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServerServer struct {
}

func (UnimplementedHelloServerServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloServerServer) mustEmbedUnimplementedHelloServerServer() {}

// UnsafeHelloServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServerServer will
// result in compilation errors.
type UnsafeHelloServerServer interface {
	mustEmbedUnimplementedHelloServerServer()
}

func RegisterHelloServerServer(s grpc.ServiceRegistrar, srv HelloServerServer) {
	s.RegisterService(&HelloServer_ServiceDesc, srv)
}

func _HelloServer_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServerServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HelloServer/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServerServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloServer_ServiceDesc is the grpc.ServiceDesc for HelloServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.HelloServer",
	HandlerType: (*HelloServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloServer_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
