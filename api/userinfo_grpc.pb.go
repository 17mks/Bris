// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: userinfo.proto

package api

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

// UserInfoClient is the client API for UserInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInfoClient interface {
	// 用户基础信息表数据添加
	UserInfoCreate(ctx context.Context, in *UserInfoCreateRequest, opts ...grpc.CallOption) (*UserInfoCreateResponse, error)
	// 用户信息表数据删除
	UserInfoDelete(ctx context.Context, in *UserInfoDeleteRequest, opts ...grpc.CallOption) (*UserInfoDeleteResponse, error)
	// 用户信息表数据更新
	UserInfoUpdate(ctx context.Context, in *UserInfoUpdateRequest, opts ...grpc.CallOption) (*UserInfoUpdateResponse, error)
	// 用户信息表数据查询
	UserInfoDetail(ctx context.Context, in *UserInfoDetailRequest, opts ...grpc.CallOption) (*UserInfoDetailResponse, error)
	// 用户信息表数据过滤查询
	UserInfoFilter(ctx context.Context, in *UserInfoFilterRequest, opts ...grpc.CallOption) (*UserInfoFilterResponse, error)
}

type userInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoClient(cc grpc.ClientConnInterface) UserInfoClient {
	return &userInfoClient{cc}
}

func (c *userInfoClient) UserInfoCreate(ctx context.Context, in *UserInfoCreateRequest, opts ...grpc.CallOption) (*UserInfoCreateResponse, error) {
	out := new(UserInfoCreateResponse)
	err := c.cc.Invoke(ctx, "/api.UserInfo/UserInfoCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoClient) UserInfoDelete(ctx context.Context, in *UserInfoDeleteRequest, opts ...grpc.CallOption) (*UserInfoDeleteResponse, error) {
	out := new(UserInfoDeleteResponse)
	err := c.cc.Invoke(ctx, "/api.UserInfo/UserInfoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoClient) UserInfoUpdate(ctx context.Context, in *UserInfoUpdateRequest, opts ...grpc.CallOption) (*UserInfoUpdateResponse, error) {
	out := new(UserInfoUpdateResponse)
	err := c.cc.Invoke(ctx, "/api.UserInfo/UserInfoUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoClient) UserInfoDetail(ctx context.Context, in *UserInfoDetailRequest, opts ...grpc.CallOption) (*UserInfoDetailResponse, error) {
	out := new(UserInfoDetailResponse)
	err := c.cc.Invoke(ctx, "/api.UserInfo/UserInfoDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoClient) UserInfoFilter(ctx context.Context, in *UserInfoFilterRequest, opts ...grpc.CallOption) (*UserInfoFilterResponse, error) {
	out := new(UserInfoFilterResponse)
	err := c.cc.Invoke(ctx, "/api.UserInfo/UserInfoFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServer is the server API for UserInfo service.
// All implementations must embed UnimplementedUserInfoServer
// for forward compatibility
type UserInfoServer interface {
	// 用户基础信息表数据添加
	UserInfoCreate(context.Context, *UserInfoCreateRequest) (*UserInfoCreateResponse, error)
	// 用户信息表数据删除
	UserInfoDelete(context.Context, *UserInfoDeleteRequest) (*UserInfoDeleteResponse, error)
	// 用户信息表数据更新
	UserInfoUpdate(context.Context, *UserInfoUpdateRequest) (*UserInfoUpdateResponse, error)
	// 用户信息表数据查询
	UserInfoDetail(context.Context, *UserInfoDetailRequest) (*UserInfoDetailResponse, error)
	// 用户信息表数据过滤查询
	UserInfoFilter(context.Context, *UserInfoFilterRequest) (*UserInfoFilterResponse, error)
	mustEmbedUnimplementedUserInfoServer()
}

// UnimplementedUserInfoServer must be embedded to have forward compatible implementations.
type UnimplementedUserInfoServer struct {
}

func (UnimplementedUserInfoServer) UserInfoCreate(context.Context, *UserInfoCreateRequest) (*UserInfoCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfoCreate not implemented")
}
func (UnimplementedUserInfoServer) UserInfoDelete(context.Context, *UserInfoDeleteRequest) (*UserInfoDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfoDelete not implemented")
}
func (UnimplementedUserInfoServer) UserInfoUpdate(context.Context, *UserInfoUpdateRequest) (*UserInfoUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfoUpdate not implemented")
}
func (UnimplementedUserInfoServer) UserInfoDetail(context.Context, *UserInfoDetailRequest) (*UserInfoDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfoDetail not implemented")
}
func (UnimplementedUserInfoServer) UserInfoFilter(context.Context, *UserInfoFilterRequest) (*UserInfoFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfoFilter not implemented")
}
func (UnimplementedUserInfoServer) mustEmbedUnimplementedUserInfoServer() {}

// UnsafeUserInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInfoServer will
// result in compilation errors.
type UnsafeUserInfoServer interface {
	mustEmbedUnimplementedUserInfoServer()
}

func RegisterUserInfoServer(s grpc.ServiceRegistrar, srv UserInfoServer) {
	s.RegisterService(&UserInfo_ServiceDesc, srv)
}

func _UserInfo_UserInfoCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).UserInfoCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserInfo/UserInfoCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).UserInfoCreate(ctx, req.(*UserInfoCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfo_UserInfoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).UserInfoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserInfo/UserInfoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).UserInfoDelete(ctx, req.(*UserInfoDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfo_UserInfoUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).UserInfoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserInfo/UserInfoUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).UserInfoUpdate(ctx, req.(*UserInfoUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfo_UserInfoDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).UserInfoDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserInfo/UserInfoDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).UserInfoDetail(ctx, req.(*UserInfoDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfo_UserInfoFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServer).UserInfoFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserInfo/UserInfoFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServer).UserInfoFilter(ctx, req.(*UserInfoFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInfo_ServiceDesc is the grpc.ServiceDesc for UserInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserInfo",
	HandlerType: (*UserInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserInfoCreate",
			Handler:    _UserInfo_UserInfoCreate_Handler,
		},
		{
			MethodName: "UserInfoDelete",
			Handler:    _UserInfo_UserInfoDelete_Handler,
		},
		{
			MethodName: "UserInfoUpdate",
			Handler:    _UserInfo_UserInfoUpdate_Handler,
		},
		{
			MethodName: "UserInfoDetail",
			Handler:    _UserInfo_UserInfoDetail_Handler,
		},
		{
			MethodName: "UserInfoFilter",
			Handler:    _UserInfo_UserInfoFilter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userinfo.proto",
}
