package data

import (
	"context"
	"followup/gencode/chdb"
	"followup/internal/biz"
	"followup/model"
	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepo) UserDetailPreloadById(ctx context.Context, id string) (*model.UserDetailInfo, error) {
	userDetailInfo := model.UserDetailInfo{
		TbUser:          chdb.TbUser{},
		TbUserInfo:      chdb.TbUserInfo{},
		TbThirdAccounts: make([]chdb.TbThirdAccount, 0),
	}
	tx := r.data.gormDB.Model(&model.UserDetailInfo{}).Where("id = ?", id)
	tx.Preload("TbUserInfo")
	tx.Preload("TbThirdAccounts")

	if err := tx.First(&userDetailInfo).Error; err != nil {
		return nil, err
	}
	return &userDetailInfo, nil
}
