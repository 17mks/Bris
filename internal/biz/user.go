package biz

import (
	"context"
	"followup/model"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	UserDetailPreloadById(ctx context.Context, userId string) (*model.UserDetailInfo, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) UserDetailPreloadById(ctx context.Context, userId string) (*model.UserDetailInfo, error) {
	uc.log.WithContext(ctx).Infof("UserDetailPreloadById")
	return uc.repo.UserDetailPreloadById(ctx, userId)
}
