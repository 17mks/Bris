package data

import (
	"context"
	"followup/internal/biz"
	"followup/internal/conf"
	A "followup/internal/pkg/middleware/auth"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type AuthRepo struct {
	data *Data
	jwtc *conf.JWT
	log  *log.Helper
}

type Auth struct {
	gorm.Model
	UserId   string `gorm:"size:500"`
	UserName string `gorm:"size:500"`
	AppId    string `gorm:"size:500"`
	Token    string `gorm:"size:500"`
}

func (r *AuthRepo) generateToken(userID string) string {
	return A.GenerateToken(r.jwtc.Secret, userID)
}

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &AuthRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *AuthRepo) Save(ctx context.Context, auth *biz.Auth) error {
	a := Auth{
		UserId:   auth.Userid,
		UserName: auth.Username,
		AppId:    auth.Appid,
	}
	res := r.data.gormDB.Create(&a)
	return res.Error
}
