package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"goProject/internal/biz"
	xerrors "github.com/pkg/errors"
	"strings"

	pb "goProject/api/user/v1"
	v1 "goProject/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer

	uu  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uu *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uu: uu, log:log.NewHelper(logger)}
}

func (s *UserService) Signin(ctx context.Context, req *pb.SigninRequest) (*pb.SigninReply, error) {
	us := &biz.User{
		Name: strings.TrimSpace(req.Body.GetUsername()),
		Password: strings.TrimSpace(req.Body.GetPassword()),
	}

	if us.Name == "" {
		return nil, errors.New("request parameter error")
	}

	if us.Password == "" {
		return nil, errors.New("request parameter error")
	}

	if err := s.uu.GetUser(ctx, us); err != nil {
		s.log.Errorf("service 'Signin' is error", xerrors.WithStack(err))
		return nil, err
	}

	return &v1.SigninReply{Token: "xxxxx"}, nil
}
