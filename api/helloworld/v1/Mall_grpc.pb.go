// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: api/helloworld/v1/mall.proto

package v1

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
	Mall_CreateGood_FullMethodName         = "/api.helloworld.v1.Mall/CreateGood"
	Mall_UpdateGood_FullMethodName         = "/api.helloworld.v1.Mall/UpdateGood"
	Mall_GetGoodsByKind_FullMethodName     = "/api.helloworld.v1.Mall/GetGoodsByKind"
	Mall_GetGoodsByUserName_FullMethodName = "/api.helloworld.v1.Mall/GetGoodsByUserName"
	Mall_GetOwnGoods_FullMethodName        = "/api.helloworld.v1.Mall/GetOwnGoods"
	Mall_GetGoodsByName_FullMethodName     = "/api.helloworld.v1.Mall/GetGoodsByName"
	Mall_GetGoodByID_FullMethodName        = "/api.helloworld.v1.Mall/GetGoodByID"
	Mall_DeleteGoodByID_FullMethodName     = "/api.helloworld.v1.Mall/DeleteGoodByID"
	Mall_CreateProfile_FullMethodName      = "/api.helloworld.v1.Mall/CreateProfile"
	Mall_UpdateProfile_FullMethodName      = "/api.helloworld.v1.Mall/UpdateProfile"
	Mall_GetProfile_FullMethodName         = "/api.helloworld.v1.Mall/GetProfile"
	Mall_Register_FullMethodName           = "/api.helloworld.v1.Mall/Register"
	Mall_Login_FullMethodName              = "/api.helloworld.v1.Mall/Login"
)

// MallClient is the client API for Mall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MallClient interface {
	// 商品信息接口--------------------------------------------------------------------------------------------------
	CreateGood(ctx context.Context, in *CreateGoodRequest, opts ...grpc.CallOption) (*CreateGoodReply, error)
	UpdateGood(ctx context.Context, in *UpdateGoodRequest, opts ...grpc.CallOption) (*UpdateGoodReply, error)
	GetGoodsByKind(ctx context.Context, in *GetGoodsByKindRequest, opts ...grpc.CallOption) (*GetGoodsReply, error)
	GetGoodsByUserName(ctx context.Context, in *GetGoodsByUserNameRequest, opts ...grpc.CallOption) (*GetGoodsReply, error)
	GetOwnGoods(ctx context.Context, in *GetOwnGoodsRequest, opts ...grpc.CallOption) (*GetGoodsReply, error)
	GetGoodsByName(ctx context.Context, in *GetGoodsByNameRequest, opts ...grpc.CallOption) (*GetGoodsReply, error)
	GetGoodByID(ctx context.Context, in *GetGoodByIDRequest, opts ...grpc.CallOption) (*GetGoodReply, error)
	DeleteGoodByID(ctx context.Context, in *DeleteGoodByIDRequest, opts ...grpc.CallOption) (*DeleteGoodReply, error)
	// 用户信息接口---------------------------------------------------------------------------------------------------
	CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*CreateProfileReply, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileReply, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileReply, error)
	// 登录注册接口--------------------------------------------------------------------------------------------------
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type mallClient struct {
	cc grpc.ClientConnInterface
}

func NewMallClient(cc grpc.ClientConnInterface) MallClient {
	return &mallClient{cc}
}

func (c *mallClient) CreateGood(ctx context.Context, in *CreateGoodRequest, opts ...grpc.CallOption) (*CreateGoodReply, error) {
	out := new(CreateGoodReply)
	err := c.cc.Invoke(ctx, Mall_CreateGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) UpdateGood(ctx context.Context, in *UpdateGoodRequest, opts ...grpc.CallOption) (*UpdateGoodReply, error) {
	out := new(UpdateGoodReply)
	err := c.cc.Invoke(ctx, Mall_UpdateGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) GetGoodsByKind(ctx context.Context, in *GetGoodsByKindRequest, opts ...grpc.CallOption) (*GetGoodsReply, error) {
	out := new(GetGoodsReply)
	err := c.cc.Invoke(ctx, Mall_GetGoodsByKind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) GetGoodsByUserName(ctx context.Context, in *GetGoodsByUserNameRequest, opts ...grpc.CallOption) (*GetGoodsReply, error) {
	out := new(GetGoodsReply)
	err := c.cc.Invoke(ctx, Mall_GetGoodsByUserName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) GetOwnGoods(ctx context.Context, in *GetOwnGoodsRequest, opts ...grpc.CallOption) (*GetGoodsReply, error) {
	out := new(GetGoodsReply)
	err := c.cc.Invoke(ctx, Mall_GetOwnGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) GetGoodsByName(ctx context.Context, in *GetGoodsByNameRequest, opts ...grpc.CallOption) (*GetGoodsReply, error) {
	out := new(GetGoodsReply)
	err := c.cc.Invoke(ctx, Mall_GetGoodsByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) GetGoodByID(ctx context.Context, in *GetGoodByIDRequest, opts ...grpc.CallOption) (*GetGoodReply, error) {
	out := new(GetGoodReply)
	err := c.cc.Invoke(ctx, Mall_GetGoodByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) DeleteGoodByID(ctx context.Context, in *DeleteGoodByIDRequest, opts ...grpc.CallOption) (*DeleteGoodReply, error) {
	out := new(DeleteGoodReply)
	err := c.cc.Invoke(ctx, Mall_DeleteGoodByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*CreateProfileReply, error) {
	out := new(CreateProfileReply)
	err := c.cc.Invoke(ctx, Mall_CreateProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileReply, error) {
	out := new(UpdateProfileReply)
	err := c.cc.Invoke(ctx, Mall_UpdateProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileReply, error) {
	out := new(GetProfileReply)
	err := c.cc.Invoke(ctx, Mall_GetProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, Mall_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mallClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, Mall_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MallServer is the server API for Mall service.
// All implementations must embed UnimplementedMallServer
// for forward compatibility
type MallServer interface {
	// 商品信息接口--------------------------------------------------------------------------------------------------
	CreateGood(context.Context, *CreateGoodRequest) (*CreateGoodReply, error)
	UpdateGood(context.Context, *UpdateGoodRequest) (*UpdateGoodReply, error)
	GetGoodsByKind(context.Context, *GetGoodsByKindRequest) (*GetGoodsReply, error)
	GetGoodsByUserName(context.Context, *GetGoodsByUserNameRequest) (*GetGoodsReply, error)
	GetOwnGoods(context.Context, *GetOwnGoodsRequest) (*GetGoodsReply, error)
	GetGoodsByName(context.Context, *GetGoodsByNameRequest) (*GetGoodsReply, error)
	GetGoodByID(context.Context, *GetGoodByIDRequest) (*GetGoodReply, error)
	DeleteGoodByID(context.Context, *DeleteGoodByIDRequest) (*DeleteGoodReply, error)
	// 用户信息接口---------------------------------------------------------------------------------------------------
	CreateProfile(context.Context, *CreateProfileRequest) (*CreateProfileReply, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileReply, error)
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileReply, error)
	// 登录注册接口--------------------------------------------------------------------------------------------------
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	mustEmbedUnimplementedMallServer()
}

// UnimplementedMallServer must be embedded to have forward compatible implementations.
type UnimplementedMallServer struct {
}

func (UnimplementedMallServer) CreateGood(context.Context, *CreateGoodRequest) (*CreateGoodReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGood not implemented")
}
func (UnimplementedMallServer) UpdateGood(context.Context, *UpdateGoodRequest) (*UpdateGoodReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGood not implemented")
}
func (UnimplementedMallServer) GetGoodsByKind(context.Context, *GetGoodsByKindRequest) (*GetGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsByKind not implemented")
}
func (UnimplementedMallServer) GetGoodsByUserName(context.Context, *GetGoodsByUserNameRequest) (*GetGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsByUserName not implemented")
}
func (UnimplementedMallServer) GetOwnGoods(context.Context, *GetOwnGoodsRequest) (*GetGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOwnGoods not implemented")
}
func (UnimplementedMallServer) GetGoodsByName(context.Context, *GetGoodsByNameRequest) (*GetGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsByName not implemented")
}
func (UnimplementedMallServer) GetGoodByID(context.Context, *GetGoodByIDRequest) (*GetGoodReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodByID not implemented")
}
func (UnimplementedMallServer) DeleteGoodByID(context.Context, *DeleteGoodByIDRequest) (*DeleteGoodReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoodByID not implemented")
}
func (UnimplementedMallServer) CreateProfile(context.Context, *CreateProfileRequest) (*CreateProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (UnimplementedMallServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedMallServer) GetProfile(context.Context, *GetProfileRequest) (*GetProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedMallServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedMallServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedMallServer) mustEmbedUnimplementedMallServer() {}

// UnsafeMallServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MallServer will
// result in compilation errors.
type UnsafeMallServer interface {
	mustEmbedUnimplementedMallServer()
}

func RegisterMallServer(s grpc.ServiceRegistrar, srv MallServer) {
	s.RegisterService(&Mall_ServiceDesc, srv)
}

func _Mall_CreateGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).CreateGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mall_CreateGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).CreateGood(ctx, req.(*CreateGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_UpdateGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).UpdateGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mall_UpdateGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).UpdateGood(ctx, req.(*UpdateGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_GetGoodsByKind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsByKindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).GetGoodsByKind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mall_GetGoodsByKind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).GetGoodsByKind(ctx, req.(*GetGoodsByKindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_GetGoodsByUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsByUserNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).GetGoodsByUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mall_GetGoodsByUserName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).GetGoodsByUserName(ctx, req.(*GetGoodsByUserNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_GetOwnGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOwnGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).GetOwnGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mall_GetOwnGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).GetOwnGoods(ctx, req.(*GetOwnGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_GetGoodsByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).GetGoodsByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mall_GetGoodsByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).GetGoodsByName(ctx, req.(*GetGoodsByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_GetGoodByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).GetGoodByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mall_GetGoodByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).GetGoodByID(ctx, req.(*GetGoodByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_DeleteGoodByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGoodByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).DeleteGoodByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mall_DeleteGoodByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).DeleteGoodByID(ctx, req.(*DeleteGoodByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mall_CreateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).CreateProfile(ctx, req.(*CreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mall_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mall_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mall_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mall_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MallServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mall_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MallServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mall_ServiceDesc is the grpc.ServiceDesc for Mall service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mall_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.helloworld.v1.Mall",
	HandlerType: (*MallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGood",
			Handler:    _Mall_CreateGood_Handler,
		},
		{
			MethodName: "UpdateGood",
			Handler:    _Mall_UpdateGood_Handler,
		},
		{
			MethodName: "GetGoodsByKind",
			Handler:    _Mall_GetGoodsByKind_Handler,
		},
		{
			MethodName: "GetGoodsByUserName",
			Handler:    _Mall_GetGoodsByUserName_Handler,
		},
		{
			MethodName: "GetOwnGoods",
			Handler:    _Mall_GetOwnGoods_Handler,
		},
		{
			MethodName: "GetGoodsByName",
			Handler:    _Mall_GetGoodsByName_Handler,
		},
		{
			MethodName: "GetGoodByID",
			Handler:    _Mall_GetGoodByID_Handler,
		},
		{
			MethodName: "DeleteGoodByID",
			Handler:    _Mall_DeleteGoodByID_Handler,
		},
		{
			MethodName: "CreateProfile",
			Handler:    _Mall_CreateProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _Mall_UpdateProfile_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _Mall_GetProfile_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Mall_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Mall_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/helloworld/v1/mall.proto",
}
