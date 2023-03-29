package biz

import (
	"context"
	v1 "followup/api"
	"followup/gencode/chdb"
	"followup/model"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// TbWorkItem 工作项表
type TbWorkItem struct {
	ID                string
	Title             string
	Type              string
	Status            string
	PrincipalType     string
	PrincipalID       string
	PrincipalName     string
	Participant       string
	Cc                string
	Tag               string
	Pid               string
	PatientID         string
	PatientName       string
	PlanID            string
	EventStartTime    string
	AssignedType      string
	AssignedTo        string
	AssignedToName    string
	PlanStartTime     time.Time
	PlanEndTime       time.Time
	ActualStartTime   time.Time
	ActualEndTime     time.Time
	Description       string
	CreatBy           string
	CreatByName       string
	UpdateBy          string
	BelongType        string
	BelongTo          string
	SortNum           int32
	Event             string
	FrequencyInterval int32
	FrequencyUnit     string
	AppID             string
	ExecArea          string
	NotifyLeftDate    time.Time
	NotifyLeftOffset  int32
	NotifyNode        string
	NotifyOffsetUnit  string
	NotifyRightDate   time.Time
	NotifyRightOffset int32
	CreateTime        time.Time
	UpdateTime        time.Time
	DeleteAt          time.Time
}

// 定义Followup 的操作接口
type FollowupRepo interface {
	FollowupWiCreate(ctx context.Context, req *chdb.TbWorkItem) (*chdb.TbWorkItem, error)
	FollowupWiDelById(ctx context.Context, id string) (string, error)
	FollowupWiDetailQueryById(ctx context.Context, id string) (*model.FollowupWiDetailPreload, error)
	FollowupWiFilterQuery(ctx context.Context, apiReq *v1.FollowupFilterRequest) ([]model.FollowupWiFilterPreload, int64, error)

	FollowupWiDelByIdRecursion(ctx context.Context, id string) ([]string, error) // 递归删除工作项
}

type FollowupUseCase struct {
	repo FollowupRepo
	log  *log.Helper
}

func NewFollowupUseCase(repo FollowupRepo, logger log.Logger) *FollowupUseCase {
	return &FollowupUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *FollowupUseCase) FollowupWiCreate(ctx context.Context, req *chdb.TbWorkItem) (*chdb.TbWorkItem, error) {
	uc.log.WithContext(ctx).Infof("FollowupWiCreate")
	return uc.repo.FollowupWiCreate(ctx, req)
}

func (uc *FollowupUseCase) FollowupWiDelById(ctx context.Context, id string) (string, error) {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.FollowupWiDelById(ctx, id)
}

func (uc *FollowupUseCase) FollowupWiDetailQueryById(ctx context.Context, id string) (*model.FollowupWiDetailPreload, error) {
	uc.log.WithContext(ctx).Infof("Get: %v", id)
	return uc.repo.FollowupWiDetailQueryById(ctx, id)
}

func (uc *FollowupUseCase) FollowupWiFilterQuery(ctx context.Context, apiReq *v1.FollowupFilterRequest) ([]model.FollowupWiFilterPreload, int64, error) {
	uc.log.WithContext(ctx).Infof("filter")
	return uc.repo.FollowupWiFilterQuery(ctx, apiReq)
}

func (uc *FollowupUseCase) FollowupWiDelByIdRecursion(ctx context.Context, id string) ([]string, error) {
	uc.log.WithContext(ctx).Infof("FollowupWiDelByIdRecursion")
	return uc.repo.FollowupWiDelByIdRecursion(ctx, id)
}
