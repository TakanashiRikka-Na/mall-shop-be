package biz

import "github.com/go-kratos/kratos/v2/log"

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

type UserRepo interface {
	Register(user User) error
	Login(user User) (string, error)
}

type User struct {
	UserName string
	Password string
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func (u UserUseCase) Register(user User) error {
	return u.repo.Register(user)
}

func (u UserUseCase) Login(user User) (string, error) {
	return u.repo.Login(user)
}
