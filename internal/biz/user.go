package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type (
	User struct {
		Name     string
		Password string
	}

	UserRepo interface {
		GetUser(context.Context, *User) error
	}

	UserUsecase struct {
		repo UserRepo
		log  *log.Helper
	}
)


func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uu *UserUsecase) GetUser(ctx context.Context, us *User) error {
	return uu.repo.GetUser(ctx, us)
}