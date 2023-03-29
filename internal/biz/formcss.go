package biz

import (
	"context"
	"followup/gencode/chdb"
	"github.com/go-kratos/kratos/v2/log"
)

// 定义 FormCss 的操作接口
type FormCssRepo interface {
	FormCssCreate(ctx context.Context, req *chdb.TbFormCSS) (*chdb.TbFormCSS, error)
	UpdateFormCssById(ctx context.Context, tbFormCss *chdb.TbFormCSS) (*chdb.TbFormCSS, error)
}

type FormCssUsecase struct {
	repo FormCssRepo
	log  *log.Helper
}

func NewFormCssUsecase(repo FormCssRepo, logger log.Logger) *FormCssUsecase {
	return &FormCssUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *FormCssUsecase) FormCssCreate(ctx context.Context, req *chdb.TbFormCSS) (*chdb.TbFormCSS, error) {
	uc.log.WithContext(ctx).Infof("FormCssCreate")
	return uc.repo.FormCssCreate(ctx, req)
}

func (uc *FormCssUsecase) UpdateFormCssById(ctx context.Context, tbFormCss *chdb.TbFormCSS) (*chdb.TbFormCSS, error) {
	uc.log.WithContext(ctx).Infof("UpdateFormCssById")
	return uc.repo.FormCssCreate(ctx, tbFormCss)
}
