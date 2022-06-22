// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: proto/proto.proto

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

// RemotePlayerClient is the client API for RemotePlayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemotePlayerClient interface {
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*Empty, error)
	Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*InfoResponse, error)
}

type remotePlayerClient struct {
	cc grpc.ClientConnInterface
}

func NewRemotePlayerClient(cc grpc.ClientConnInterface) RemotePlayerClient {
	return &remotePlayerClient{cc}
}

func (c *remotePlayerClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.RemotePlayer/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remotePlayerClient) Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/proto.RemotePlayer/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemotePlayerServer is the server API for RemotePlayer service.
// All implementations must embed UnimplementedRemotePlayerServer
// for forward compatibility
type RemotePlayerServer interface {
	Run(context.Context, *RunRequest) (*Empty, error)
	Info(context.Context, *Empty) (*InfoResponse, error)
	mustEmbedUnimplementedRemotePlayerServer()
}

// UnimplementedRemotePlayerServer must be embedded to have forward compatible implementations.
type UnimplementedRemotePlayerServer struct {
}

func (UnimplementedRemotePlayerServer) Run(context.Context, *RunRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedRemotePlayerServer) Info(context.Context, *Empty) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedRemotePlayerServer) mustEmbedUnimplementedRemotePlayerServer() {}

// UnsafeRemotePlayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemotePlayerServer will
// result in compilation errors.
type UnsafeRemotePlayerServer interface {
	mustEmbedUnimplementedRemotePlayerServer()
}

func RegisterRemotePlayerServer(s grpc.ServiceRegistrar, srv RemotePlayerServer) {
	s.RegisterService(&RemotePlayer_ServiceDesc, srv)
}

func _RemotePlayer_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemotePlayerServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RemotePlayer/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemotePlayerServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemotePlayer_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemotePlayerServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RemotePlayer/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemotePlayerServer).Info(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RemotePlayer_ServiceDesc is the grpc.ServiceDesc for RemotePlayer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemotePlayer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RemotePlayer",
	HandlerType: (*RemotePlayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _RemotePlayer_Run_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _RemotePlayer_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proto.proto",
}