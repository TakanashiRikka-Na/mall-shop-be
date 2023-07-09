package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mall-shop-be/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{data: data, log: log.NewHelper(logger)}
}

func (u *userRepo) Register() error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepo) Login() (string, error) {
	//TODO implement me
	panic("implement me")
}
