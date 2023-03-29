package biz

import (
	"context"
	v1 "followup/api"
	"followup/gencode/chdb"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// TbDisease 病种表
type TbDisease struct {
	ID          string
	Code        string
	Name        string
	NameJp      string
	NameQp      string
	Version     string
	Status      string
	Tag         string
	Description string
	Pid         string
	CreateTime  *time.Time
	UpdateTime  *time.Time
	//DeleteAt    gorm.DeletedAt `gorm:"column:delete_at;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'删除时间'"`                 // 删除时间
}

// 定义Disease 的操作接口
type DiseaseRepo interface {
	Save(ctx context.Context, disease *chdb.TbDisease) error
	DeleteByID(context.Context, string) (string, error)
	Update(context.Context, *chdb.TbDisease) (*chdb.TbDisease, error)
	GetByID(context.Context, string) (*chdb.TbDisease, error)
	Filter(ctx context.Context, req *v1.DiseaseFilterRequest) ([]*chdb.TbDisease, int64, error)

	QueryDiseasesByIds(ctx context.Context, ids []string) ([]chdb.TbDisease, error)
	QueryDisFunByIds(ctx context.Context, ids []string) ([]chdb.TbDisFunc, error)
}

type DiseaseUseCase struct {
	repo DiseaseRepo
	log  *log.Helper
}

func NewDiseaseUseCase(repo DiseaseRepo, logger log.Logger) *DiseaseUseCase {
	return &DiseaseUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *DiseaseUseCase) CreateDisease(ctx context.Context, disease *chdb.TbDisease) (*chdb.TbDisease, error) {
	d := &chdb.TbDisease{
		ID:          disease.ID,
		Code:        disease.Code,
		Name:        disease.Name,
		NameJp:      disease.NameJp,
		NameQp:      disease.NameQp,
		Version:     disease.Version,
		Status:      disease.Status,
		Tag:         disease.Tag,
		Description: disease.Description,
		Pid:         disease.Pid,
		CreateTime:  disease.CreateTime,
		UpdateTime:  disease.UpdateTime,
	}
	if err := uc.repo.Save(ctx, d); err != nil {
		return nil, err
	}
	return &chdb.TbDisease{
		ID: disease.ID,
	}, nil
}

func (uc *DiseaseUseCase) DeleteDisease(ctx context.Context, id string) (string, error) {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.DeleteByID(ctx, id)
}

func (uc *DiseaseUseCase) UpdateDisease(ctx context.Context, disease *chdb.TbDisease) (*chdb.TbDisease, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", disease.ID)
	return uc.repo.Update(ctx, disease)
}

func (uc *DiseaseUseCase) DetailDisease(ctx context.Context, id string) (*chdb.TbDisease, error) {
	uc.log.WithContext(ctx).Infof("Get: %v", id)
	return uc.repo.GetByID(ctx, id)
}

func (uc *DiseaseUseCase) FilterDisease(ctx context.Context, req *v1.DiseaseFilterRequest) ([]*chdb.TbDisease, int64, error) {
	uc.log.WithContext(ctx).Infof("filter")
	return uc.repo.Filter(ctx, req)
}

func (uc *DiseaseUseCase) QueryDiseasesByIds(ctx context.Context, ids []string) ([]chdb.TbDisease, error) {
	uc.log.WithContext(ctx).Infof("QueryDiseasesByIds")
	return uc.repo.QueryDiseasesByIds(ctx, ids)
}

func (uc *DiseaseUseCase) QueryDisFunByIds(ctx context.Context, ids []string) ([]chdb.TbDisFunc, error) {
	uc.log.WithContext(ctx).Infof("QueryDisFunByIds")
	return uc.repo.QueryDisFunByIds(ctx, ids)
}
