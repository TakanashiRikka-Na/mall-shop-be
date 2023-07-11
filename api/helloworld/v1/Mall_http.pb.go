// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v4.23.2
// source: api/helloworld/v1/Mall.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationProfileCreateProfile = "/api.helloworld.v1.Profile/CreateProfile"
const OperationProfileGetProfile = "/api.helloworld.v1.Profile/GetProfile"
const OperationProfileUpdateProfile = "/api.helloworld.v1.Profile/UpdateProfile"

type ProfileHTTPServer interface {
	CreateProfile(context.Context, *CreateProfileRequest) (*CreateProfileReply, error)
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileReply, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileReply, error)
}

func RegisterProfileHTTPServer(s *http.Server, srv ProfileHTTPServer) {
	r := s.Route("/")
	r.POST("/user/profile", _Profile_CreateProfile0_HTTP_Handler(srv))
	r.PUT("/user/profile", _Profile_UpdateProfile0_HTTP_Handler(srv))
	r.GET("/user/profile", _Profile_GetProfile0_HTTP_Handler(srv))
}

func _Profile_CreateProfile0_HTTP_Handler(srv ProfileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateProfileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProfileCreateProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateProfile(ctx, req.(*CreateProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateProfileReply)
		return ctx.Result(200, reply)
	}
}

func _Profile_UpdateProfile0_HTTP_Handler(srv ProfileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateProfileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProfileUpdateProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateProfile(ctx, req.(*UpdateProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateProfileReply)
		return ctx.Result(200, reply)
	}
}

func _Profile_GetProfile0_HTTP_Handler(srv ProfileHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProfileRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProfileGetProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProfile(ctx, req.(*GetProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProfileReply)
		return ctx.Result(200, reply)
	}
}

type ProfileHTTPClient interface {
	CreateProfile(ctx context.Context, req *CreateProfileRequest, opts ...http.CallOption) (rsp *CreateProfileReply, err error)
	GetProfile(ctx context.Context, req *GetProfileRequest, opts ...http.CallOption) (rsp *GetProfileReply, err error)
	UpdateProfile(ctx context.Context, req *UpdateProfileRequest, opts ...http.CallOption) (rsp *UpdateProfileReply, err error)
}

type ProfileHTTPClientImpl struct {
	cc *http.Client
}

func NewProfileHTTPClient(client *http.Client) ProfileHTTPClient {
	return &ProfileHTTPClientImpl{client}
}

func (c *ProfileHTTPClientImpl) CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...http.CallOption) (*CreateProfileReply, error) {
	var out CreateProfileReply
	pattern := "/user/profile"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProfileCreateProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProfileHTTPClientImpl) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...http.CallOption) (*GetProfileReply, error) {
	var out GetProfileReply
	pattern := "/user/profile"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProfileGetProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProfileHTTPClientImpl) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...http.CallOption) (*UpdateProfileReply, error) {
	var out UpdateProfileReply
	pattern := "/user/profile"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProfileUpdateProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationUserLogin = "/api.helloworld.v1.User/Login"
const OperationUserRegister = "/api.helloworld.v1.User/Register"

type UserHTTPServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/register", _User_Register0_HTTP_Handler(srv))
	r.POST("/login", _User_Login0_HTTP_Handler(srv))
}

func _User_Register0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _User_Login0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
