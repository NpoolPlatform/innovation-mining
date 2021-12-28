// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package npool

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

// InnovationMinningClient is the client API for InnovationMinning service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InnovationMinningClient interface {
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type innovationMinningClient struct {
	cc grpc.ClientConnInterface
}

func NewInnovationMinningClient(cc grpc.ClientConnInterface) InnovationMinningClient {
	return &innovationMinningClient{cc}
}

func (c *innovationMinningClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/innovation.minning.v1.InnovationMinning/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InnovationMinningServer is the server API for InnovationMinning service.
// All implementations must embed UnimplementedInnovationMinningServer
// for forward compatibility
type InnovationMinningServer interface {
	// Method Version
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	mustEmbedUnimplementedInnovationMinningServer()
}

// UnimplementedInnovationMinningServer must be embedded to have forward compatible implementations.
type UnimplementedInnovationMinningServer struct {
}

func (UnimplementedInnovationMinningServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedInnovationMinningServer) mustEmbedUnimplementedInnovationMinningServer() {}

// UnsafeInnovationMinningServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InnovationMinningServer will
// result in compilation errors.
type UnsafeInnovationMinningServer interface {
	mustEmbedUnimplementedInnovationMinningServer()
}

func RegisterInnovationMinningServer(s grpc.ServiceRegistrar, srv InnovationMinningServer) {
	s.RegisterService(&InnovationMinning_ServiceDesc, srv)
}

func _InnovationMinning_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnovationMinningServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/innovation.minning.v1.InnovationMinning/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnovationMinningServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// InnovationMinning_ServiceDesc is the grpc.ServiceDesc for InnovationMinning service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InnovationMinning_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "innovation.minning.v1.InnovationMinning",
	HandlerType: (*InnovationMinningServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _InnovationMinning_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/innovation-minning.proto",
}
