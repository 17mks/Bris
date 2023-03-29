package biz

import (
	"context"
	v1 "followup/api"
	"followup/gencode/chdb"
	"followup/model"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// a PlanUsecase model.
type TbPlan struct {
	ID               string
	Name             string
	Type             string
	Status           string
	BelongType       string
	BelongTo         string
	ApplyDisease     string
	ApplyDysfunction string
	ApplyAges        string
	Event            string
	CreatorID        string
	CreatorName      string
	CreateTime       *time.Time
	UpdateTime       *time.Time
}

type PlanRepo interface {
	Save(ctx context.Context, req *chdb.TbPlan) (*chdb.TbPlan, error)
	Delete(ctx context.Context, id string) (string, error)
	Update(ctx context.Context, plan *chdb.TbPlan) (*chdb.TbPlan, error)
	GetByID(ctx context.Context, id string) (*chdb.TbPlan, error)
	Filter(ctx context.Context, req *v1.PlanFilterQueryRequest) ([]chdb.TbPlan, int64, error)

	PlanDetailPreloadById(ctx context.Context, id string) (*model.PlanDetailPreloadInfo, error)
	PlanPreload(ctx context.Context, id string) (*model.PlanPreload, error)
	// PlanCreateWithResource 带资源创建方案
	PlanCreateWithResource(ctx context.Context, tbPlan *chdb.TbPlan, relates []chdb.TbPlanRelate, items []chdb.TbWorkItem, tbRelates []chdb.TbRelate) (*chdb.TbPlan, error)

	PlanUpdateWithResource(ctx context.Context, planResource *model.ToUpdatePlanResource) (*chdb.TbPlan, error) //  更新新方案和资源详情
	SaPlanDetailPreloadById(ctx context.Context, id string) (*v1.SaPlanDetailPreload, error)                    // 预加载症状评估方案详情
	QueryPlanByProjectId(ctx context.Context, projectId string) (*chdb.TbPlan, error)                           // 预加载症状评估方案详情
}

type PlanUsecase struct {
	repo PlanRepo
	log  *log.Helper
}

func NewPlanUsecase(repo PlanRepo, logger log.Logger) *PlanUsecase {
	return &PlanUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PlanUsecase) Save(ctx context.Context, req *chdb.TbPlan) (*chdb.TbPlan, error) {
	uc.log.WithContext(ctx).Infof("Create")
	return uc.repo.Save(ctx, req)
}

//func (uc *planUsecase) CreatePlan(ctx context.Context, plan *v1.Plans) (*v1.Plans, error) {
//	p := &v1.Plans{
//		Id:               plan.Id,
//		Name:             plan.Name,
//		Type:             plan.Type,
//		Status:           plan.Status,
//		BelongType:       plan.BelongType,
//		BelongTo:         plan.BelongTo,
//		ApplyDisease:     plan.ApplyDisease,
//		ApplyDysfunction: plan.ApplyDysfunction,
//		ApplyAges:        plan.ApplyAges,
//		Event:            plan.Event,
//		CreatorId:        plan.CreatorId,
//		CreatorName:      plan.CreatorName,
//		CreateTime:       plan.CreateTime,
//		UpdateTime:       plan.UpdateTime,
//	}
//	if err := uc.repo.Save(ctx, p); err != nil {
//		return nil, err
//	}
//	return &v1.Plans{
//		Id: plan.Id,
//	}, nil
//}

func (uc *PlanUsecase) PlanDelete(ctx context.Context, id string) (string, error) {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

func (uc *PlanUsecase) UpdatePlan(ctx context.Context, plan *chdb.TbPlan) (*chdb.TbPlan, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", plan.ID)
	return uc.repo.Update(ctx, plan)
}

func (uc *PlanUsecase) DetailPlan(ctx context.Context, id string) (*chdb.TbPlan, error) {
	uc.log.WithContext(ctx).Infof("Get: %d", id)
	return uc.repo.GetByID(ctx, id)
}

func (uc *PlanUsecase) FilterPlan(ctx context.Context, apiReq *v1.PlanFilterQueryRequest) ([]chdb.TbPlan, int64, error) {
	uc.log.WithContext(ctx).Infof("filter")
	return uc.repo.Filter(ctx, apiReq)
}

func (uc *PlanUsecase) PlanDetailPreloadById(ctx context.Context, id string) (*model.PlanDetailPreloadInfo, error) {
	return uc.repo.PlanDetailPreloadById(ctx, id)
}
func (uc *PlanUsecase) PlanPreload(ctx context.Context, id string) (*model.PlanPreload, error) {
	return uc.repo.PlanPreload(ctx, id)
}

func (uc *PlanUsecase) PlanCreateWithResource(ctx context.Context, tbPlan *chdb.TbPlan, relates []chdb.TbPlanRelate, items []chdb.TbWorkItem, tbRelates []chdb.TbRelate) (*chdb.TbPlan, error) {
	uc.log.WithContext(ctx).Infof("PlanCreateWithResource")
	return uc.repo.PlanCreateWithResource(ctx, tbPlan, relates, items, tbRelates)
}

func (uc *PlanUsecase) PlanUpdateWithResource(ctx context.Context, planResource *model.ToUpdatePlanResource) (*chdb.TbPlan, error) {
	uc.log.WithContext(ctx).Infof("PlanUpdateWithResource")
	return uc.repo.PlanUpdateWithResource(ctx, planResource)
}

func (uc *PlanUsecase) SaPlanDetailPreloadById(ctx context.Context, id string) (*v1.SaPlanDetailPreload, error) {
	uc.log.WithContext(ctx).Infof("SaPlanDetailPreloadById")
	return uc.repo.SaPlanDetailPreloadById(ctx, id)
}

func (uc *PlanUsecase) QueryPlanByProjectId(ctx context.Context, projectId string) (*chdb.TbPlan, error) {
	uc.log.WithContext(ctx).Infof("QueryPlanByProjectId")
	return uc.repo.QueryPlanByProjectId(ctx, projectId)
}
