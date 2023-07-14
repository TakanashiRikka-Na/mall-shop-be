package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "mall-shop-be/api/helloworld/v1"
	"mall-shop-be/internal/biz"
)

func (s *MallService) CreateGood(ctx context.Context, req *pb.CreateGoodRequest) (*pb.CreateGoodReply, error) {
	err := s.gc.CreateGood(biz.Good{
		GoodID:      req.ID,
		Name:        req.Name,
		Description: req.Description,
		Kind:        req.Kind,
		Price:       req.Price,
		FromUser:    req.FromUser,
		IsSold:      req.IsSold,
	})
	if err != nil {
		return &pb.CreateGoodReply{}, err
	}
	return &pb.CreateGoodReply{
		GoodID: req.ID,
		Msg:    "success",
		Code:   "0",
	}, nil
}
func (s *MallService) UpdateGood(ctx context.Context, req *pb.UpdateGoodRequest) (*pb.UpdateGoodReply, error) {
	err := s.gc.UpdateGood(biz.Good{
		GoodID:      req.ID,
		Name:        req.Name,
		Description: req.Description,
		Kind:        req.Kind,
		Price:       req.Price,
		FromUser:    req.FromUser,
		IsSold:      req.IsSold,
	})
	if err != nil {
		return &pb.UpdateGoodReply{}, err
	}
	return &pb.UpdateGoodReply{
		GoodID: req.ID,
		Msg:    "success",
		Code:   "0",
	}, nil
}
func (s *MallService) GetGoodsByKind(ctx context.Context, req *pb.GetGoodsByKindRequest) (*pb.GetGoodsReply, error) {
	goods, err := s.gc.GetGoodsByKind(req.Kind, biz.PageMsg{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		return nil, err
	}
	return TransformToPbs(goods), nil
}
func (s *MallService) GetGoodsByUserName(ctx context.Context, req *pb.GetGoodsByUserNameRequest) (*pb.GetGoodsReply, error) {
	goods, err := s.gc.GetGoodsByUserName(req.UserName, biz.PageMsg{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		return nil, err
	}
	return TransformToPbs(goods), nil
}
func (s *MallService) GetOwnGoods(ctx context.Context, req *pb.GetOwnGoodsRequest) (*pb.GetGoodsReply, error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		log.Info("can't read context")
		return nil, s.erc.UnKnownError()
	}
	goods, err := s.gc.GetOwnGoods(username, biz.PageMsg{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		return nil, err
	}
	return TransformToPbs(goods), nil
}
func (s *MallService) GetGoodsByName(ctx context.Context, req *pb.GetGoodsByNameRequest) (*pb.GetGoodsReply, error) {
	goods, err := s.gc.GetGoodsByName(req.Name, biz.PageMsg{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		return nil, err
	}
	return TransformToPbs(goods), nil
}
func (s *MallService) GetGoodByID(ctx context.Context, req *pb.GetGoodByIDRequest) (*pb.GetGoodReply, error) {
	good, err := s.gc.GetGoodByID(req.GetID())
	if err != nil {
		return nil, err
	}
	return TransformToPb(good), nil
}
func (s *MallService) DeleteGoodByID(ctx context.Context, req *pb.DeleteGoodByIDRequest) (*pb.DeleteGoodReply, error) {
	if err := s.gc.DeleteGoodByID(req.ID); err != nil {
		return nil, err
	}
	return &pb.DeleteGoodReply{GoodID: req.ID, Msg: "success", Code: "0"}, nil
}

func TransformToPbs(bizGoods []biz.Good) *pb.GetGoodsReply {
	var goodsPb []*pb.GetGoodsReply_Good
	for _, good := range bizGoods {
		var pbGood = &pb.GetGoodsReply_Good{
			ID:          good.GoodID,
			Name:        good.Name,
			Description: good.Description,
			Price:       good.Price,
			FromUser:    good.FromUser,
			Kind:        good.Kind,
			IsSold:      good.IsSold,
		}
		goodsPb = append(goodsPb, pbGood)
	}
	return &pb.GetGoodsReply{
		Goods: goodsPb,
		Msg:   "success",
		Code:  "0",
	}
}
func TransformToPb(good biz.Good) *pb.GetGoodReply {
	return &pb.GetGoodReply{
		ID:          good.GoodID,
		Name:        good.Name,
		Description: good.Description,
		Price:       good.Price,
		FromUser:    good.FromUser,
		Kind:        good.Kind,
		IsSold:      good.IsSold,
		Msg:         "success",
		Code:        "0",
	}
}
