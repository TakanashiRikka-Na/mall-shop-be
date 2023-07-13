package service

import (
	"github.com/google/wire"
	pb "mall-shop-be/api/helloworld/v1"
	"mall-shop-be/internal/biz"
	errs "mall-shop-be/internal/errors"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMallService)

type MallService struct {
	pb.UnimplementedMallServer
	erc errs.ErrorController
	uc  *biz.UserUseCase
	pc  *biz.ProfileUseCase
	gc  *biz.GoodsUseCase
}

func NewMallService(uc *biz.UserUseCase, pc *biz.ProfileUseCase, gc *biz.GoodsUseCase, erc errs.ErrorController) *MallService {
	return &MallService{uc: uc, pc: pc, gc: gc, erc: erc}
}
