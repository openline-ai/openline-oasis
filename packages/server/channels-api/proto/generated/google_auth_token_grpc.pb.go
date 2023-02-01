// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: google_auth_token.proto

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

// GmailAuthTokenServiceClient is the client API for GmailAuthTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GmailAuthTokenServiceClient interface {
	SetGmailAuth(ctx context.Context, in *GmailCredential, opts ...grpc.CallOption) (*EventEmpty, error)
}

type gmailAuthTokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGmailAuthTokenServiceClient(cc grpc.ClientConnInterface) GmailAuthTokenServiceClient {
	return &gmailAuthTokenServiceClient{cc}
}

func (c *gmailAuthTokenServiceClient) SetGmailAuth(ctx context.Context, in *GmailCredential, opts ...grpc.CallOption) (*EventEmpty, error) {
	out := new(EventEmpty)
	err := c.cc.Invoke(ctx, "/GmailAuthTokenService/setGmailAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GmailAuthTokenServiceServer is the server API for GmailAuthTokenService service.
// All implementations must embed UnimplementedGmailAuthTokenServiceServer
// for forward compatibility
type GmailAuthTokenServiceServer interface {
	SetGmailAuth(context.Context, *GmailCredential) (*EventEmpty, error)
	mustEmbedUnimplementedGmailAuthTokenServiceServer()
}

// UnimplementedGmailAuthTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGmailAuthTokenServiceServer struct {
}

func (UnimplementedGmailAuthTokenServiceServer) SetGmailAuth(context.Context, *GmailCredential) (*EventEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGmailAuth not implemented")
}
func (UnimplementedGmailAuthTokenServiceServer) mustEmbedUnimplementedGmailAuthTokenServiceServer() {}

// UnsafeGmailAuthTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GmailAuthTokenServiceServer will
// result in compilation errors.
type UnsafeGmailAuthTokenServiceServer interface {
	mustEmbedUnimplementedGmailAuthTokenServiceServer()
}

func RegisterGmailAuthTokenServiceServer(s grpc.ServiceRegistrar, srv GmailAuthTokenServiceServer) {
	s.RegisterService(&GmailAuthTokenService_ServiceDesc, srv)
}

func _GmailAuthTokenService_SetGmailAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GmailCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GmailAuthTokenServiceServer).SetGmailAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GmailAuthTokenService/setGmailAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GmailAuthTokenServiceServer).SetGmailAuth(ctx, req.(*GmailCredential))
	}
	return interceptor(ctx, in, info, handler)
}

// GmailAuthTokenService_ServiceDesc is the grpc.ServiceDesc for GmailAuthTokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GmailAuthTokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GmailAuthTokenService",
	HandlerType: (*GmailAuthTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "setGmailAuth",
			Handler:    _GmailAuthTokenService_SetGmailAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google_auth_token.proto",
}
