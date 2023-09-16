// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: protos/test.proto

package Lab1

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/protos.ChatService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	SayHello(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ChatService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ChatService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/test.proto",
}

// NumberServiceClient is the client API for NumberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NumberServiceClient interface {
	SendKeys(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*NumberResponse, error)
}

type numberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNumberServiceClient(cc grpc.ClientConnInterface) NumberServiceClient {
	return &numberServiceClient{cc}
}

func (c *numberServiceClient) SendKeys(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*NumberResponse, error) {
	out := new(NumberResponse)
	err := c.cc.Invoke(ctx, "/protos.NumberService/SendKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NumberServiceServer is the server API for NumberService service.
// All implementations must embed UnimplementedNumberServiceServer
// for forward compatibility
type NumberServiceServer interface {
	SendKeys(context.Context, *NumberRequest) (*NumberResponse, error)
	mustEmbedUnimplementedNumberServiceServer()
}

// UnimplementedNumberServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNumberServiceServer struct {
}

func (UnimplementedNumberServiceServer) SendKeys(context.Context, *NumberRequest) (*NumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendKeys not implemented")
}
func (UnimplementedNumberServiceServer) mustEmbedUnimplementedNumberServiceServer() {}

// UnsafeNumberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NumberServiceServer will
// result in compilation errors.
type UnsafeNumberServiceServer interface {
	mustEmbedUnimplementedNumberServiceServer()
}

func RegisterNumberServiceServer(s grpc.ServiceRegistrar, srv NumberServiceServer) {
	s.RegisterService(&NumberService_ServiceDesc, srv)
}

func _NumberService_SendKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NumberServiceServer).SendKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.NumberService/SendKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NumberServiceServer).SendKeys(ctx, req.(*NumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NumberService_ServiceDesc is the grpc.ServiceDesc for NumberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NumberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.NumberService",
	HandlerType: (*NumberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendKeys",
			Handler:    _NumberService_SendKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/test.proto",
}
