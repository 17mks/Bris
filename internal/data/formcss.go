package data

import (
	"context"
	"followup/gencode/chdb"
	"followup/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type FormCssRepo struct {
	data *Data
	log  *log.Helper
}

func NewFormCssRepo(data *Data, logger log.Logger) biz.FormCssRepo {
	return &FormCssRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *FormCssRepo) FormCssCreate(ctx context.Context, tbFormCss *chdb.TbFormCSS) (*chdb.TbFormCSS, error) {
	if err := r.data.gormDB.Model(&chdb.TbFormCSS{}).Create(tbFormCss).Error; err != nil {
		return nil, err
	}
	return tbFormCss, nil
}

func (r *FormCssRepo) UpdateFormCssById(ctx context.Context, tbFormCss *chdb.TbFormCSS) (*chdb.TbFormCSS, error) {
	if err := r.data.gormDB.Save(tbFormCss).Error; err != nil {
		return nil, err
	}
	return tbFormCss, nil
}
