package biz

import (
	"context"
	v1 "followup/api"
	"followup/gencode/chdb"
	"github.com/go-kratos/kratos/v2/log"
)

// 定义FormRow 的操作接口
type FormRowRepo interface {
	Save(ctx context.Context, req *chdb.TbFormRow) (*chdb.TbFormRow, error)
	DeleteByID(ctx context.Context, id string) (string, error)
	Update(ctx context.Context, tbFormRow *chdb.TbFormRow) (*chdb.TbFormRow, error)
	GetByID(ctx context.Context, id string) (*chdb.TbFormRow, error)
	Filter(ctx context.Context, req *v1.FormRowFilterRequest) ([]chdb.TbFormRow, int64, error)
}

type FormRowUseCase struct {
	repo FormRowRepo
	log  *log.Helper
}

func NewFormRowUseCase(repo FormRowRepo, logger log.Logger) *FormRowUseCase {
	return &FormRowUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *FormRowUseCase) FormRowCreate(ctx context.Context, req *chdb.TbFormRow) (*chdb.TbFormRow, error) {
	uc.log.WithContext(ctx).Infof("FormRowCreate")
	return uc.repo.Save(ctx, req)
}

func (uc *FormRowUseCase) FormRowDelete(ctx context.Context, id string) (string, error) {
	uc.log.WithContext(ctx).Infof("FormRowDelete")
	return uc.repo.DeleteByID(ctx, id)
}

func (uc *FormRowUseCase) FormRowUpdate(ctx context.Context, tbFormRow *chdb.TbFormRow) (*chdb.TbFormRow, error) {
	uc.log.WithContext(ctx).Infof("FormRowUpdate")
	return uc.repo.Update(ctx, tbFormRow)
}

func (uc *FormRowUseCase) FormRowDetail(ctx context.Context, id string) (*chdb.TbFormRow, error) {
	uc.log.WithContext(ctx).Infof("FormRowDetail")
	return uc.repo.GetByID(ctx, id)
}
func (uc *FormRowUseCase) FormRowFilter(ctx context.Context, req *v1.FormRowFilterRequest) ([]chdb.TbFormRow, int64, error) {
	uc.log.WithContext(ctx).Infof("FormRowFilter")
	return uc.repo.Filter(ctx, req)
}
