package data

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"mall-shop-be/internal/biz"
	errs "mall-shop-be/internal/errors"
	"strconv"
)

type goodsRepo struct {
	data *Data
	log  *log.Helper
	erc  errs.ErrorController
}

type Good struct {
	ID          int32
	GoodID      string
	Name        string
	Kind        int64
	Description string
	Price       string
	FromUser    string
	IsSold      bool
}

func (g *goodsRepo) unknownErr(err error, msg string) *errors.Error {
	return g.erc.UnKnownError().WithMetadata(map[string]string{msg: err.Error()})
}

func NewGoodsRepo(data *Data, logger log.Logger, controller errs.ErrorController) biz.GoodsRepo {
	return &goodsRepo{log: log.NewHelper(logger), data: data, erc: controller}
}
func (g *goodsRepo) CreateGood(good biz.Good) error {
	//检查物品是否存在
	var (
		isExisted bool
		err       error
	)
	if err = g.data.DB.Model(&Good{}).Select("1").Where("good_id=?", good.GoodID).Find(&isExisted).Error; err != nil {
		g.log.Error(err)
		return g.unknownErr(err, GoodsError)
	}
	if isExisted {
		return g.erc.GoodExist()
	}
	tx := g.data.DB.Begin()
	err = tx.Model(&Good{}).Create(&Good{
		GoodID:      good.GoodID,
		Name:        good.Name,
		Kind:        good.Kind,
		Description: good.Description,
		Price:       good.Price,
		FromUser:    good.FromUser,
		IsSold:      good.IsSold,
	}).Error
	if err != nil {
		g.log.Error(err)
		tx.Rollback()
		return g.unknownErr(err, GoodsError)
	}
	tx.Commit()
	return nil
}

func (g *goodsRepo) UpdateGood(good biz.Good) error {
	//检查物品是否存在
	var (
		isExisted bool
		err       error
	)
	if err = g.data.DB.Model(&Good{}).Select("1").Where("good_id=?", good.GoodID).Find(&isExisted).Error; err != nil {
		g.log.Error(err)
		return g.unknownErr(err, GoodsError)
	}
	if isExisted {
		return g.erc.GoodExist()
	}
	tx := g.data.DB.Begin()
	err = tx.Model(&Good{}).Where("good_id=?", good.GoodID).Updates(&Good{
		GoodID:      good.GoodID,
		Name:        good.Name,
		Kind:        good.Kind,
		Description: good.Description,
		Price:       good.Price,
		FromUser:    good.FromUser,
		IsSold:      good.IsSold,
	}).Error
	if err != nil {
		g.log.Error(err)
		tx.Rollback()
		return g.unknownErr(err, GoodsError)
	}
	//删除所有该商品的缓存
	err = DeleteCache(g.data.Rdb, good.GoodID, strconv.FormatInt(good.Kind, 10), good.Name, good.FromUser)
	tx.Commit()
	return nil
}

func (g *goodsRepo) GetGoodsByKind(kind string) ([]biz.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (g *goodsRepo) GetGoodsByUserName(username string) ([]biz.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (g *goodsRepo) GetOwnGoods() ([]biz.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (g *goodsRepo) GetGoodsByName(goods []biz.Good, err error) {
	//TODO implement me
	panic("implement me")
}

func (g *goodsRepo) GetGoodByID(id string) (biz.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (g *goodsRepo) DeleteGoodByID(id string) error {
	//TODO implement me
	panic("implement me")
}
