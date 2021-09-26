package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"goProject/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) GetUser(ctx context.Context, us *biz.User) error {
	return nil
}
