package biz

import "github.com/go-kratos/kratos/v2/log"

type GoodsRepo interface {
	CreateGood(good Good) error
	UpdateGood(good Good) error
	GetGoodsByKind(kind string) ([]Good, error)
	GetGoodsByUserName(username string) ([]Good, error)
	GetOwnGoods() ([]Good, error)
	GetGoodsByName([]Good, error)
	GetGoodByID(id string) (Good, error)
	DeleteGoodByID(id string) error
}

type Good struct {
	GoodID      string
	Name        string
	Kind        int64
	Description string
	Price       string
	FromUser    string
	IsSold      bool
}

type GoodsUseCase struct {
	repo GoodsRepo
	log  *log.Helper
}

func NewGoodsUseCase(repo GoodsRepo, logger log.Logger) *GoodsUseCase {
	return &GoodsUseCase{repo: repo, log: log.NewHelper(logger)}
}
