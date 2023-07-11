package service

import (
	"context"
	"mall-shop-be/internal/biz"

	pb "mall-shop-be/api/helloworld/v1"
)

func (s *MallService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	if err := s.uc.Register(biz.User{UserName: req.UserName, Password: req.Password}); err != nil {
		return nil, err
	}
	return &pb.RegisterReply{UserName: req.GetUserName(), Code: 0, Msg: "register success"}, nil
}
func (s *MallService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	token, err := s.uc.Login(biz.User{UserName: req.UserName, Password: req.Password})
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{Msg: "login success", Code: 0, Token: token, UserName: req.GetUserName()}, nil
}
