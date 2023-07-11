package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "mall-shop-be/api/helloworld/v1"
	"mall-shop-be/internal/biz"
)

func (s *MallService) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileReply, error) {
	err := s.pc.UpdateProfile(biz.Profile{
		UserName: req.UserName,
		Gender:   req.Gender,
		Email:    req.Email,
		Phone:    req.Phone,
		Addr:     req.Addr,
		Order:    req.Order,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProfileReply{
		Msg:      "success",
		Code:     0,
		UserName: req.GetUserName(),
	}, nil
}
func (s *MallService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileReply, error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		log.Info("can't read context")
		return nil, s.erc.UnKnownError()
	}
	profile, err := s.pc.GetProfile(username)
	if err != nil {
		return nil, err
	}
	return &pb.GetProfileReply{
		Order:    profile.Order,
		Email:    profile.Email,
		UserName: profile.UserName,
		Phone:    profile.Phone,
		Gender:   profile.Gender,
		Addr:     profile.Addr,
	}, nil
}
func (s *MallService) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileReply, error) {
	err := s.pc.CreateProfile(biz.Profile{
		UserName: req.UserName,
		Gender:   req.Gender,
		Email:    req.Email,
		Phone:    req.Phone,
		Addr:     req.Addr,
		Order:    req.Order,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateProfileReply{
		Msg:      "success",
		Code:     0,
		UserName: req.GetUserName(),
	}, nil
}
