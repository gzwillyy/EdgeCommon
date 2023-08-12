// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_user_access_key.proto

package pb

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
	UserAccessKeyService_CreateUserAccessKey_FullMethodName           = "/pb.UserAccessKeyService/createUserAccessKey"
	UserAccessKeyService_FindAllEnabledUserAccessKeys_FullMethodName  = "/pb.UserAccessKeyService/findAllEnabledUserAccessKeys"
	UserAccessKeyService_DeleteUserAccessKey_FullMethodName           = "/pb.UserAccessKeyService/deleteUserAccessKey"
	UserAccessKeyService_UpdateUserAccessKeyIsOn_FullMethodName       = "/pb.UserAccessKeyService/updateUserAccessKeyIsOn"
	UserAccessKeyService_CountAllEnabledUserAccessKeys_FullMethodName = "/pb.UserAccessKeyService/countAllEnabledUserAccessKeys"
)

// UserAccessKeyServiceClient is the client API for UserAccessKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAccessKeyServiceClient interface {
	// 创建AccessKey
	CreateUserAccessKey(ctx context.Context, in *CreateUserAccessKeyRequest, opts ...grpc.CallOption) (*CreateUserAccessKeyResponse, error)
	// 查找所有的AccessKey
	FindAllEnabledUserAccessKeys(ctx context.Context, in *FindAllEnabledUserAccessKeysRequest, opts ...grpc.CallOption) (*FindAllEnabledUserAccessKeysResponse, error)
	// 删除AccessKey
	DeleteUserAccessKey(ctx context.Context, in *DeleteUserAccessKeyRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 设置是否启用AccessKey
	UpdateUserAccessKeyIsOn(ctx context.Context, in *UpdateUserAccessKeyIsOnRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算AccessKey数量
	CountAllEnabledUserAccessKeys(ctx context.Context, in *CountAllEnabledUserAccessKeysRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
}

type userAccessKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAccessKeyServiceClient(cc grpc.ClientConnInterface) UserAccessKeyServiceClient {
	return &userAccessKeyServiceClient{cc}
}

func (c *userAccessKeyServiceClient) CreateUserAccessKey(ctx context.Context, in *CreateUserAccessKeyRequest, opts ...grpc.CallOption) (*CreateUserAccessKeyResponse, error) {
	out := new(CreateUserAccessKeyResponse)
	err := c.cc.Invoke(ctx, UserAccessKeyService_CreateUserAccessKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccessKeyServiceClient) FindAllEnabledUserAccessKeys(ctx context.Context, in *FindAllEnabledUserAccessKeysRequest, opts ...grpc.CallOption) (*FindAllEnabledUserAccessKeysResponse, error) {
	out := new(FindAllEnabledUserAccessKeysResponse)
	err := c.cc.Invoke(ctx, UserAccessKeyService_FindAllEnabledUserAccessKeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccessKeyServiceClient) DeleteUserAccessKey(ctx context.Context, in *DeleteUserAccessKeyRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, UserAccessKeyService_DeleteUserAccessKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccessKeyServiceClient) UpdateUserAccessKeyIsOn(ctx context.Context, in *UpdateUserAccessKeyIsOnRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, UserAccessKeyService_UpdateUserAccessKeyIsOn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccessKeyServiceClient) CountAllEnabledUserAccessKeys(ctx context.Context, in *CountAllEnabledUserAccessKeysRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, UserAccessKeyService_CountAllEnabledUserAccessKeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAccessKeyServiceServer is the server API for UserAccessKeyService service.
// All implementations should embed UnimplementedUserAccessKeyServiceServer
// for forward compatibility
type UserAccessKeyServiceServer interface {
	// 创建AccessKey
	CreateUserAccessKey(context.Context, *CreateUserAccessKeyRequest) (*CreateUserAccessKeyResponse, error)
	// 查找所有的AccessKey
	FindAllEnabledUserAccessKeys(context.Context, *FindAllEnabledUserAccessKeysRequest) (*FindAllEnabledUserAccessKeysResponse, error)
	// 删除AccessKey
	DeleteUserAccessKey(context.Context, *DeleteUserAccessKeyRequest) (*RPCSuccess, error)
	// 设置是否启用AccessKey
	UpdateUserAccessKeyIsOn(context.Context, *UpdateUserAccessKeyIsOnRequest) (*RPCSuccess, error)
	// 计算AccessKey数量
	CountAllEnabledUserAccessKeys(context.Context, *CountAllEnabledUserAccessKeysRequest) (*RPCCountResponse, error)
}

// UnimplementedUserAccessKeyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserAccessKeyServiceServer struct {
}

func (UnimplementedUserAccessKeyServiceServer) CreateUserAccessKey(context.Context, *CreateUserAccessKeyRequest) (*CreateUserAccessKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserAccessKey not implemented")
}
func (UnimplementedUserAccessKeyServiceServer) FindAllEnabledUserAccessKeys(context.Context, *FindAllEnabledUserAccessKeysRequest) (*FindAllEnabledUserAccessKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllEnabledUserAccessKeys not implemented")
}
func (UnimplementedUserAccessKeyServiceServer) DeleteUserAccessKey(context.Context, *DeleteUserAccessKeyRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserAccessKey not implemented")
}
func (UnimplementedUserAccessKeyServiceServer) UpdateUserAccessKeyIsOn(context.Context, *UpdateUserAccessKeyIsOnRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAccessKeyIsOn not implemented")
}
func (UnimplementedUserAccessKeyServiceServer) CountAllEnabledUserAccessKeys(context.Context, *CountAllEnabledUserAccessKeysRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountAllEnabledUserAccessKeys not implemented")
}

// UnsafeUserAccessKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAccessKeyServiceServer will
// result in compilation errors.
type UnsafeUserAccessKeyServiceServer interface {
	mustEmbedUnimplementedUserAccessKeyServiceServer()
}

func RegisterUserAccessKeyServiceServer(s grpc.ServiceRegistrar, srv UserAccessKeyServiceServer) {
	s.RegisterService(&UserAccessKeyService_ServiceDesc, srv)
}

func _UserAccessKeyService_CreateUserAccessKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserAccessKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessKeyServiceServer).CreateUserAccessKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAccessKeyService_CreateUserAccessKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessKeyServiceServer).CreateUserAccessKey(ctx, req.(*CreateUserAccessKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccessKeyService_FindAllEnabledUserAccessKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllEnabledUserAccessKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessKeyServiceServer).FindAllEnabledUserAccessKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAccessKeyService_FindAllEnabledUserAccessKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessKeyServiceServer).FindAllEnabledUserAccessKeys(ctx, req.(*FindAllEnabledUserAccessKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccessKeyService_DeleteUserAccessKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserAccessKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessKeyServiceServer).DeleteUserAccessKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAccessKeyService_DeleteUserAccessKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessKeyServiceServer).DeleteUserAccessKey(ctx, req.(*DeleteUserAccessKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccessKeyService_UpdateUserAccessKeyIsOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAccessKeyIsOnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessKeyServiceServer).UpdateUserAccessKeyIsOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAccessKeyService_UpdateUserAccessKeyIsOn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessKeyServiceServer).UpdateUserAccessKeyIsOn(ctx, req.(*UpdateUserAccessKeyIsOnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccessKeyService_CountAllEnabledUserAccessKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountAllEnabledUserAccessKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccessKeyServiceServer).CountAllEnabledUserAccessKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserAccessKeyService_CountAllEnabledUserAccessKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccessKeyServiceServer).CountAllEnabledUserAccessKeys(ctx, req.(*CountAllEnabledUserAccessKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAccessKeyService_ServiceDesc is the grpc.ServiceDesc for UserAccessKeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAccessKeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserAccessKeyService",
	HandlerType: (*UserAccessKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createUserAccessKey",
			Handler:    _UserAccessKeyService_CreateUserAccessKey_Handler,
		},
		{
			MethodName: "findAllEnabledUserAccessKeys",
			Handler:    _UserAccessKeyService_FindAllEnabledUserAccessKeys_Handler,
		},
		{
			MethodName: "deleteUserAccessKey",
			Handler:    _UserAccessKeyService_DeleteUserAccessKey_Handler,
		},
		{
			MethodName: "updateUserAccessKeyIsOn",
			Handler:    _UserAccessKeyService_UpdateUserAccessKeyIsOn_Handler,
		},
		{
			MethodName: "countAllEnabledUserAccessKeys",
			Handler:    _UserAccessKeyService_CountAllEnabledUserAccessKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_user_access_key.proto",
}