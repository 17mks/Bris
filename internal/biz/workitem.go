package biz

import (
	"context"
	"followup/gencode/chdb"
	"followup/model"
	"github.com/go-kratos/kratos/v2/log"
)

// 定义 WorkItem 的操作接口
type WorkItemRepo interface {
	// QueryPlanWorkItemByIds 根据工作项编码查询方案下工作项信息
	QueryPlanWorkItemByIds(ctx context.Context, ids []string) ([]model.WorkItemDetailBean, error)
	// WorkItemCreateWithRelates 新建工作项-附带工作项关联
	WorkItemCreateWithRelates(ctx context.Context, tbWorkItems []chdb.TbWorkItem, tbRelates []chdb.TbRelate) error

	WorkItemPreloadByRelateId(ctx context.Context, relateId string) (*model.WorkItemPreloadInfo, error)

	WorkItemDetailPreloadById(ctx context.Context, id string) (*model.WorkItemDetailBean, error)
}

type WorkItemUsecase struct {
	repo WorkItemRepo
	log  *log.Helper
}

func NewWorkItemUsecase(repo WorkItemRepo, logger log.Logger) *WorkItemUsecase {
	return &WorkItemUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *WorkItemUsecase) QueryPlanWorkItemByIds(ctx context.Context, ids []string) ([]model.WorkItemDetailBean, error) {
	uc.log.WithContext(ctx).Infof("QueryPlanWorkItemByIds")
	return uc.repo.QueryPlanWorkItemByIds(ctx, ids)
}

func (uc *WorkItemUsecase) WorkItemCreateWithRelates(ctx context.Context, tbWorkItems []chdb.TbWorkItem, tbRelates []chdb.TbRelate) error {
	uc.log.WithContext(ctx).Infof("WorkItemCreateWithRelates")
	return uc.repo.WorkItemCreateWithRelates(ctx, tbWorkItems, tbRelates)
}

func (uc *WorkItemUsecase) WorkItemPreloadByRelateId(ctx context.Context, relateId string) (*model.WorkItemPreloadInfo, error) {
	uc.log.WithContext(ctx).Infof("WorkItemPreloadByRelateId")
	return uc.repo.WorkItemPreloadByRelateId(ctx, relateId)
}

func (uc *WorkItemUsecase) WorkItemDetailPreloadById(ctx context.Context, id string) (*model.WorkItemDetailBean, error) {
	uc.log.WithContext(ctx).Infof("WorkItemDetailPreloadById")
	return uc.repo.WorkItemDetailPreloadById(ctx, id)
}
