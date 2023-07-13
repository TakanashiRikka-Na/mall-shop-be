package data

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"mall-shop-be/internal/biz"
	errs "mall-shop-be/internal/errors"
	"mall-shop-be/internal/utils"
	"strings"
)

type Profile struct {
	ID        int32
	UserName  string
	Gender    string
	Phone     string
	Email     string
	Addr      string
	AddrOrder string
}

type profileRepo struct {
	data *Data
	log  *log.Helper
	erc  errs.ErrorController
}

func (p *profileRepo) unknownErr(err error, msg string) *errors.Error {
	return p.erc.UnKnownError().WithMetadata(map[string]string{msg: err.Error()})
}

func NewProfileRepo(data *Data, logger log.Logger, erc errs.ErrorController) biz.ProfileRepo {
	return &profileRepo{data: data, log: log.NewHelper(logger), erc: erc}
}

func (p *profileRepo) GetProfile(username string) (*biz.Profile, error) {
	var profile biz.Profile
	// 依旧先读 redis
	err := ReadCache(username, p.data.Rdb, &profile)
	if err == nil && !utils.IsStructEmpty(profile) {
		return &profile, nil
	}
	// 如果错误原因只是空值就不打印
	if err.Error() != "redis: nil" {
		p.log.Error("read redis error reason:", err)
	}
	var profileInfo Profile
	//从数据库拿
	if err = p.data.DB.Model(&Profile{}).Where("user_name=?", username).Find(&profileInfo).Error; err != nil {
		return nil, p.unknownErr(err, ProfileError)
	}
	Addr := strings.Split(profileInfo.Addr, ",")
	Order := strings.Split(profileInfo.AddrOrder, ",")
	profile = biz.Profile{
		UserName: profileInfo.UserName,
		Gender:   profileInfo.Gender,
		Phone:    profileInfo.Phone,
		Email:    profileInfo.Email,
		Addr:     Addr,
		Order:    Order,
	}
	err = Cache(username, &profile, p.data.Rdb, 0)
	if err != nil {
		p.log.Error(err)
	}
	return &profile, nil
}

func (p *profileRepo) UpdateProfile(profile biz.Profile) error {
	Addr := strings.Join(profile.Addr, ",")
	Order := strings.Join(profile.Order, ",")
	tx := p.data.DB.Begin()
	err := tx.Model(&Profile{}).Where("user_name=?", profile.UserName).Updates(&Profile{
		UserName:  profile.UserName,
		Gender:    profile.Gender,
		Email:     profile.Email,
		Phone:     profile.Phone,
		Addr:      Addr,
		AddrOrder: Order,
	}).Error
	if err != nil {
		p.log.Error(err)
		tx.Rollback()
		return p.unknownErr(err, ProfileError)
	}
	// 这步不允许先 commit 在 delete 必须等待 redis 成功删除旧数据后才能commit
	if err = DeleteCache(p.data.Rdb, profile.UserName); err != nil {
		p.log.Error(err)
		return p.unknownErr(err, ProfileError)
	}
	tx.Commit()
	return nil
}
func (p *profileRepo) CreateProfile(profile biz.Profile) error {
	Addr := strings.Join(profile.Addr, ",")
	Order := strings.Join(profile.Order, ",")
	err := p.data.DB.Model(&Profile{}).Create(&Profile{
		UserName:  profile.UserName,
		Gender:    profile.Gender,
		Email:     profile.Email,
		Phone:     profile.Phone,
		Addr:      Addr,
		AddrOrder: Order,
	}).Error
	if err != nil {
		p.log.Error(err)
		return p.unknownErr(err, ProfileError)
	}
	return nil
}
