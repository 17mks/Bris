package biz

import (
	"context"
	"followup/internal/conf"
	"followup/internal/pkg/middleware/auth"
	"github.com/go-kratos/kratos/v2/log"
)

// 定义 Auth 的操作接口
type AuthRepo interface {
	Save(ctx context.Context, auth *Auth) error
}

type Auth struct {
	Username string
	Userid   string
	Appid    string
}

type AuthReply struct {
	Token string
}

type AuthUsecase struct {
	repo AuthRepo
	jwtc *conf.Data
	log  *log.Helper
}

func NewAuthUsecase(repo AuthRepo, jwtc *conf.Data, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{repo: repo, jwtc: jwtc, log: log.NewHelper(logger)}
}

func (uc *AuthUsecase) generateToken(userID string) string {
	return auth.GenerateToken(uc.jwtc.Jwt.Secret, userID)
}

// Save

func (uc *AuthUsecase) CreateToken(ctx context.Context, usename, userid, appid string) (*AuthReply, error) {
	a := &Auth{
		Userid:   userid,
		Username: usename,
		Appid:    appid,
	}
	if err := uc.repo.Save(ctx, a); err != nil {
		return nil, err
	}
	return &AuthReply{
		Token: uc.generateToken(a.Userid),
	}, nil

}
