package biz

import (
	"context"
	v1 "followup/api"
	"followup/gencode/chdb"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// TbDisFunction 病种表
type TbDisFunction struct {
	ID          string
	Name        string
	Description string
	Py          string
	CreateTime  time.Time
	UpdateTime  time.Time
	DeleteAt    time.Time
}

// 定义DisFunction 的操作接口
type DisFunctionRepo interface {
	Save(ctx context.Context, disFunction *chdb.TbDisFunc) error
	Delete(context.Context, string) (string, error)
	Update(context.Context, *chdb.TbDisFunc) (*chdb.TbDisFunc, error)
	GetByID(context.Context, string) (*chdb.TbDisFunc, error)
	Filter(ctx context.Context, req *v1.DisFunctionFilterRequest) ([]chdb.TbDisFunc, int64, error)

	HasDisFunc(ctx context.Context, disFuncName string) (*chdb.TbDisFunc, bool, error)
	DisFuncPseudoDelById(ctx context.Context, id int64) (int64, error) //伪删除
	QueryDisFunByIds(ctx context.Context, ids []string) ([]chdb.TbDisFunc, error)
}

type DisFunctionUseCase struct {
	repo DisFunctionRepo
	log  *log.Helper
}

func NewDisFunctionUseCase(repo DisFunctionRepo, logger log.Logger) *DisFunctionUseCase {
	return &DisFunctionUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *DisFunctionUseCase) CreateDisFunction(ctx context.Context, DisFunction *chdb.TbDisFunc) (*chdb.TbDisFunc, error) {
	df := &chdb.TbDisFunc{
		ID:          DisFunction.ID,
		Name:        DisFunction.Name,
		Description: DisFunction.Description,
		Py:          DisFunction.Py,
		CreateTime:  DisFunction.CreateTime,
		UpdateTime:  DisFunction.UpdateTime,
		DeleteAt:    DisFunction.DeleteAt,
	}
	if err := uc.repo.Save(ctx, df); err != nil {
		return nil, err
	}
	return &chdb.TbDisFunc{
		ID: DisFunction.ID,
	}, nil
}

func (uc *DisFunctionUseCase) DeleteDisFunction(ctx context.Context, id string) (string, error) {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

func (uc *DisFunctionUseCase) UpdateDisFunction(ctx context.Context, DisFunction *chdb.TbDisFunc) (*chdb.TbDisFunc, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", DisFunction.ID)
	return uc.repo.Update(ctx, DisFunction)
}

func (uc *DisFunctionUseCase) DetailDisFunction(ctx context.Context, id string) (*chdb.TbDisFunc, error) {
	uc.log.WithContext(ctx).Infof("Get: %v", id)
	return uc.repo.GetByID(ctx, id)
}

func (uc *DisFunctionUseCase) FilterDisFunction(ctx context.Context, req *v1.DisFunctionFilterRequest) ([]chdb.TbDisFunc, int64, error) {
	uc.log.WithContext(ctx).Infof("filter")
	return uc.repo.Filter(ctx, req)
}

func (uc *DisFunctionUseCase) DisFuncPseudoDelById(ctx context.Context, id int64) (int64, error) {
	uc.log.WithContext(ctx).Infof("DisFuncPseudoDelById")
	return uc.repo.DisFuncPseudoDelById(ctx, id)
}

func (uc *DisFunctionUseCase) HasDisFunc(ctx context.Context, disFuncName string) (*chdb.TbDisFunc, bool, error) {
	uc.log.WithContext(ctx).Infof("HasDisFunc")
	return uc.repo.HasDisFunc(ctx, disFuncName)
}

func (uc *DisFunctionUseCase) QueryDisFunByIds(ctx context.Context, ids []string) ([]chdb.TbDisFunc, error) {
	uc.log.WithContext(ctx).Infof("QueryDisFunByIds")
	return uc.repo.QueryDisFunByIds(ctx, ids)
}
