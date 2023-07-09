package biz

import "github.com/go-kratos/kratos/v2/log"

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

type UserRepo interface {
	Register() error
	Login() (string, error)
}

type User struct {
	UserName string
	Password string
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func (u UserUseCase) Register() error {
	//TODO implement me
	panic("implement me")
}

func (u UserUseCase) Login() (string, error) {
	//TODO implement me
	panic("implement me")
}
