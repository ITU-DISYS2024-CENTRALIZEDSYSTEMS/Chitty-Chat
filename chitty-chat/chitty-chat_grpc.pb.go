// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: chitty-chat.proto

package chitty_chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChittyChat_PublishMessage_FullMethodName    = "/ChittyChat/PublishMessage"
	ChittyChat_JoinConversation_FullMethodName  = "/ChittyChat/JoinConversation"
	ChittyChat_LeaveConversation_FullMethodName = "/ChittyChat/LeaveConversation"
)

// ChittyChatClient is the client API for ChittyChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittyChatClient interface {
	PublishMessage(ctx context.Context, in *TextMessage, opts ...grpc.CallOption) (*TextMessage, error)
	JoinConversation(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	LeaveConversation(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type chittyChatClient struct {
	cc grpc.ClientConnInterface
}

func NewChittyChatClient(cc grpc.ClientConnInterface) ChittyChatClient {
	return &chittyChatClient{cc}
}

func (c *chittyChatClient) PublishMessage(ctx context.Context, in *TextMessage, opts ...grpc.CallOption) (*TextMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TextMessage)
	err := c.cc.Invoke(ctx, ChittyChat_PublishMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatClient) JoinConversation(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Message)
	err := c.cc.Invoke(ctx, ChittyChat_JoinConversation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatClient) LeaveConversation(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Message)
	err := c.cc.Invoke(ctx, ChittyChat_LeaveConversation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChittyChatServer is the server API for ChittyChat service.
// All implementations must embed UnimplementedChittyChatServer
// for forward compatibility.
type ChittyChatServer interface {
	PublishMessage(context.Context, *TextMessage) (*TextMessage, error)
	JoinConversation(context.Context, *Message) (*Message, error)
	LeaveConversation(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedChittyChatServer()
}

// UnimplementedChittyChatServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChittyChatServer struct{}

func (UnimplementedChittyChatServer) PublishMessage(context.Context, *TextMessage) (*TextMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishMessage not implemented")
}
func (UnimplementedChittyChatServer) JoinConversation(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinConversation not implemented")
}
func (UnimplementedChittyChatServer) LeaveConversation(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveConversation not implemented")
}
func (UnimplementedChittyChatServer) mustEmbedUnimplementedChittyChatServer() {}
func (UnimplementedChittyChatServer) testEmbeddedByValue()                    {}

// UnsafeChittyChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittyChatServer will
// result in compilation errors.
type UnsafeChittyChatServer interface {
	mustEmbedUnimplementedChittyChatServer()
}

func RegisterChittyChatServer(s grpc.ServiceRegistrar, srv ChittyChatServer) {
	// If the following call pancis, it indicates UnimplementedChittyChatServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChittyChat_ServiceDesc, srv)
}

func _ChittyChat_PublishMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServer).PublishMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChittyChat_PublishMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServer).PublishMessage(ctx, req.(*TextMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChat_JoinConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServer).JoinConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChittyChat_JoinConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServer).JoinConversation(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChat_LeaveConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServer).LeaveConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChittyChat_LeaveConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServer).LeaveConversation(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// ChittyChat_ServiceDesc is the grpc.ServiceDesc for ChittyChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChittyChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChittyChat",
	HandlerType: (*ChittyChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishMessage",
			Handler:    _ChittyChat_PublishMessage_Handler,
		},
		{
			MethodName: "JoinConversation",
			Handler:    _ChittyChat_JoinConversation_Handler,
		},
		{
			MethodName: "LeaveConversation",
			Handler:    _ChittyChat_LeaveConversation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chitty-chat.proto",
}
