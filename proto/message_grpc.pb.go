// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package helloworld

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

// DiscordMessageClient is the client API for DiscordMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscordMessageClient interface {
	SendChanMessage(ctx context.Context, in *MessageChannel, opts ...grpc.CallOption) (*ServerResponse, error)
}

type discordMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscordMessageClient(cc grpc.ClientConnInterface) DiscordMessageClient {
	return &discordMessageClient{cc}
}

func (c *discordMessageClient) SendChanMessage(ctx context.Context, in *MessageChannel, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/proto.DiscordMessage/SendChanMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscordMessageServer is the server API for DiscordMessage service.
// All implementations must embed UnimplementedDiscordMessageServer
// for forward compatibility
type DiscordMessageServer interface {
	SendChanMessage(context.Context, *MessageChannel) (*ServerResponse, error)
	mustEmbedUnimplementedDiscordMessageServer()
}

// UnimplementedDiscordMessageServer must be embedded to have forward compatible implementations.
type UnimplementedDiscordMessageServer struct {
}

func (UnimplementedDiscordMessageServer) SendChanMessage(context.Context, *MessageChannel) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChanMessage not implemented")
}
func (UnimplementedDiscordMessageServer) mustEmbedUnimplementedDiscordMessageServer() {}

// UnsafeDiscordMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscordMessageServer will
// result in compilation errors.
type UnsafeDiscordMessageServer interface {
	mustEmbedUnimplementedDiscordMessageServer()
}

func RegisterDiscordMessageServer(s grpc.ServiceRegistrar, srv DiscordMessageServer) {
	s.RegisterService(&DiscordMessage_ServiceDesc, srv)
}

func _DiscordMessage_SendChanMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageChannel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscordMessageServer).SendChanMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DiscordMessage/SendChanMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscordMessageServer).SendChanMessage(ctx, req.(*MessageChannel))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscordMessage_ServiceDesc is the grpc.ServiceDesc for DiscordMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscordMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DiscordMessage",
	HandlerType: (*DiscordMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendChanMessage",
			Handler:    _DiscordMessage_SendChanMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/message.proto",
}
