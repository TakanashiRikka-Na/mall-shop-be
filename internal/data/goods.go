package data

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"mall-shop-be/internal/biz"
	errs "mall-shop-be/internal/errors"
	"mall-shop-be/internal/utils"
	"strconv"
	"time"
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

func (g *goodsRepo) GetGoodsByKind(kind int64, msg biz.PageMsg) ([]biz.Good, error) {
	offset := (msg.Page - 1) * msg.PageSize
	goods := make([]Good, 0)
	err := g.data.DB.Model(&Good{}).Where("kind=?", kind).Offset(int(offset)).Limit(int(msg.PageSize)).Find(&goods).Error
	if err != nil {
		g.log.Error(err)
		return nil, g.unknownErr(err, GoodsError)
	}
	return TransformToBizs(goods), err
}

func (g *goodsRepo) GetGoodsByUserName(username string, msg biz.PageMsg) ([]biz.Good, error) {
	offset := (msg.Page - 1) * msg.PageSize
	goods := make([]Good, 0)
	err := g.data.DB.Model(&Good{}).Where("username=?", username).Offset(int(offset)).Limit(int(msg.PageSize)).Find(&goods).Error
	if err != nil {
		g.log.Error(err)
		return nil, g.unknownErr(err, GoodsError)
	}
	return TransformToBizs(goods), err
}

func (g *goodsRepo) GetOwnGoods(username string, msg biz.PageMsg) ([]biz.Good, error) {
	offset := (msg.Page - 1) * msg.PageSize
	goods := make([]Good, 0)
	err := g.data.DB.Model(&Good{}).Where("from_user=?", username).Offset(int(offset)).Limit(int(msg.PageSize)).Find(&goods).Error
	if err != nil {
		g.log.Error(err)
		return nil, g.unknownErr(err, GoodsError)
	}
	return TransformToBizs(goods), err
}

func (g *goodsRepo) GetGoodsByName(name string, msg biz.PageMsg) ([]biz.Good, error) {
	offset := (msg.Page - 1) * msg.PageSize
	goods := make([]Good, 0)
	err := g.data.DB.Model(&Good{}).Where("name=?", name).Offset(int(offset)).Limit(int(msg.PageSize)).Find(&goods).Error
	if err != nil {
		g.log.Error(err)
		return nil, g.unknownErr(err, GoodsError)
	}
	return TransformToBizs(goods), err
}

func (g *goodsRepo) GetGoodByID(id string) (biz.Good, error) {
	var (
		good    Good
		bizGood biz.Good
	)
	fmt.Println(id)
	//从 redis 中取
	err := ReadCache(id, g.data.Rdb, &bizGood)
	fmt.Println(bizGood)
	if err == nil && !utils.IsStructEmpty(bizGood) {
		fmt.Println(158)
		return bizGood, nil
	}
	if err != nil && err != redis.Nil {
		log.Error(err)
	}
	err = g.data.DB.Model(&Good{}).First(&Good{}, "good_id=?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return biz.Good{}, g.erc.GoodNotExist()
	}

	err = g.data.DB.Model(&Good{}).Where("good_id=?", id).Find(&good).Error
	if err != nil {
		g.log.Error(err)
		return biz.Good{}, g.unknownErr(err, GoodsError)
	}
	bizGood = TransformToBiz(good)
	if err = Cache(id, &bizGood, g.data.Rdb, time.Hour); err != nil {
		g.log.Error(err)
	}
	return bizGood, nil
}

func (g *goodsRepo) DeleteGoodByID(id string) error {
	var isExisted bool
	err := g.data.DB.Model(&Good{}).Select("1").Where("good_id=?", id).Find(&isExisted).Error
	if err != nil {
		g.log.Error(err)
		return g.unknownErr(err, GoodsError)
	}
	if !isExisted {
		return g.erc.GoodNotExist()
	}
	if err = g.data.DB.Model(&Good{}).Where("good_id=?", id).Delete(&Good{}).Error; err != nil {
		g.log.Error(err)
		return g.unknownErr(err, GoodsError)
	}
	if err = DeleteCache(g.data.Rdb, id); err != nil {
		g.log.Error(err)
		return g.unknownErr(err, GoodsError)
	}
	return nil
}

func TransformToBizs(goods []Good) (bizGoods []biz.Good) {
	bizGoods = make([]biz.Good, 0)
	for _, good := range goods {
		bizGood := biz.Good{GoodID: good.GoodID, Name: good.Name, Kind: good.Kind, Description: good.Description, Price: good.Price, FromUser: good.FromUser, IsSold: good.IsSold}
		bizGoods = append(bizGoods, bizGood)
	}
	return
}

func TransformToBiz(good Good) biz.Good {
	return biz.Good{GoodID: good.GoodID, Name: good.Name, Kind: good.Kind, Description: good.Description, Price: good.Price, FromUser: good.FromUser, IsSold: good.IsSold}
}
