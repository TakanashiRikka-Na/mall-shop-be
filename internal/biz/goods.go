package biz

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
)

type GoodsRepo interface {
	CreateGood(good Good) error
	UpdateGood(good Good) error
	GetGoodsByKind(kind int64, msg PageMsg) ([]Good, error)
	GetGoodsByUserName(username string, msg PageMsg) ([]Good, error)
	GetOwnGoods(username string, msg PageMsg) ([]Good, error)
	GetGoodsByName(name string, msg PageMsg) ([]Good, error)
	GetGoodByID(id string) (Good, error)
	DeleteGoodByID(id string) error
}

type PageMsg struct {
	Page     int64
	PageSize int64
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

func (g Good) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(g)
	return
}

type GoodsUseCase struct {
	repo GoodsRepo
	log  *log.Helper
}

func (g GoodsUseCase) CreateGood(good Good) error {
	return g.repo.CreateGood(good)
}

func (g GoodsUseCase) UpdateGood(good Good) error {
	return g.repo.UpdateGood(good)
}

func (g GoodsUseCase) GetGoodsByKind(kind int64, msg PageMsg) ([]Good, error) {
	return g.repo.GetGoodsByKind(kind, msg)
}

func (g GoodsUseCase) GetGoodsByUserName(username string, msg PageMsg) ([]Good, error) {
	return g.repo.GetGoodsByUserName(username, msg)
}

func (g GoodsUseCase) GetOwnGoods(username string, msg PageMsg) ([]Good, error) {
	return g.repo.GetOwnGoods(username, msg)
}

func (g GoodsUseCase) GetGoodsByName(name string, msg PageMsg) (goods []Good, err error) {
	return g.repo.GetGoodsByName(name, msg)
}

func (g GoodsUseCase) GetGoodByID(id string) (Good, error) {
	return g.repo.GetGoodByID(id)
}

func (g GoodsUseCase) DeleteGoodByID(id string) error {
	return g.repo.DeleteGoodByID(id)
}

func NewGoodsUseCase(repo GoodsRepo, logger log.Logger) *GoodsUseCase {
	return &GoodsUseCase{repo: repo, log: log.NewHelper(logger)}
}
