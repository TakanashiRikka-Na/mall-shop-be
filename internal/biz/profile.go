package biz

import "github.com/go-kratos/kratos/v2/log"

func NewProfileUseCase(repo ProfileRepo, logger log.Logger) *ProfileUseCase {
	return &ProfileUseCase{repo: repo, log: log.NewHelper(logger)}
}

type ProfileRepo interface {
	GetProfile(username string) (*Profile, error)
	UpdateProfile(profile Profile) error
	CreateProfile(profile Profile) error
}

type Profile struct {
	UserName string
	Gender   string
	Phone    string
	Email    string
	Addr     []string
	Order    []string
}

type ProfileUseCase struct {
	repo ProfileRepo
	log  *log.Helper
}

func (p *ProfileUseCase) CreateProfile(profile Profile) error {
	return p.repo.CreateProfile(profile)
}

func (p *ProfileUseCase) GetProfile(username string) (*Profile, error) {
	return p.repo.GetProfile(username)
}

func (p *ProfileUseCase) UpdateProfile(profile Profile) error {
	return p.repo.UpdateProfile(profile)
}
