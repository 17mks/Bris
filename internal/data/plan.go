package data

import (
	"context"
	Err "errors"
	v1 "followup/api"
	"followup/gencode/chdb"
	"followup/internal/biz"
	"followup/model"
	"followup/protocol"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"sort"
	"time"
)

type TbPlan struct {
	ID               string     `gorm:"primaryKey;column:id;type:varchar(45);not null;AUTO_INCREMENT;comment:'方案编码'"`                                                         // 方案编码
	Name             string     `gorm:"uniqueIndex:name_UNIQUE;column:name;type:varchar(128);not null;comment:'方案名称'"`                                                        // 方案名称
	Type             string     `gorm:"column:type;type:enum('SA','FOLLOWUP');not null;comment:'方案类型(SA 症状评估, FOLLOWUP 随访)'"`                                                 // 方案类型(SA 症状评估, FOLLOWUP 随访)
	Status           string     `gorm:"column:status;type:enum('DRAFT','ENABLED','DISABLED');not null;default:DRAFT;comment:'方案状态('DRAFT 草稿', 'ENABLED 启用', 'DISABLED 禁用')'"` // 方案状态('DRAFT 草稿', 'ENABLED 启用', 'DISABLED 禁用')
	BelongType       string     `gorm:"column:belong_type;type:enum('ORG','PROJECT','TEAM','GROUP','NONE');not null;default:NONE;comment:'归属类型(组织、项目、团队、组等)'"`                // 归属类型(组织、项目、团队、组等)
	BelongTo         string     `gorm:"uniqueIndex:name_UNIQUE;column:belong_to;type:varchar(45);default:null;comment:'规则资源编码(e.g. 如果归属类型时是项目则填项目编码)'"`                       // 规则资源编码(e.g. 如果归属类型时是项目则填项目编码)
	ApplyDisease     string     `gorm:"column:apply_disease;type:text;default:null;comment:'适用病种'"`                                                                           // 适用病种
	ApplyDysfunction string     `gorm:"column:apply_dysfunction;type:text;default:null;comment:'适用功能障碍(多个功能障碍用','分隔)'"`                                                       // 适用功能障碍(多个功能障碍用','分隔)
	ApplyAges        string     `gorm:"column:apply_ages;type:text;default:null;comment:'适用年龄段(多个年龄段用','分隔)'"`                                                                // 适用年龄段(多个年龄段用','分隔)
	Event            string     `gorm:"column:event;type:text;default:null;comment:'事件(随访开始触发事件)'"`                                                                           // 事件(随访开始触发事件)
	CreatorID        string     `gorm:"column:creator_id;type:varchar(45);default:null;comment:'创建人编码'"`                                                                      // 创建人编码
	CreatorName      string     `gorm:"column:creator_name;type:varchar(45);default:null;comment:'创建人名称'"`                                                                    // 创建人编码
	AppID            string     `gorm:"column:app_id;type:varchar(45);default:null;comment:'应用编码'"`
	CreateTime       *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"` // 创建时间
	UpdateTime       *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	DeletedAt        *time.Time `gorm:"column:delete_at;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'删除时间'"`
}

type planRepo struct {
	data *Data
	log  *log.Helper
}

func NewPlanRepo(data *Data, logger log.Logger) biz.PlanRepo {
	return &planRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func convertPlan(x TbPlan) *biz.TbPlan {
	return &biz.TbPlan{
		ID:               x.ID,
		Name:             x.Name,
		Type:             x.Type,
		Status:           x.Status,
		BelongType:       x.BelongType,
		BelongTo:         x.BelongTo,
		ApplyAges:        x.ApplyAges,
		ApplyDisease:     x.ApplyDisease,
		ApplyDysfunction: x.ApplyDysfunction,
		Event:            x.Event,
		CreatorID:        x.CreatorID,
		CreatorName:      x.CreatorName,
	}
}

func (repo *planRepo) Save(ctx context.Context, tbPlan *chdb.TbPlan) (*chdb.TbPlan, error) {
	if err := repo.data.gormDB.Model(&chdb.TbPlan{}).Create(tbPlan).Error; err != nil {
		return nil, err
	}
	return tbPlan, nil
}

func (repo *planRepo) Delete(ctx context.Context, id string) (string, error) {
	Plan := chdb.TbPlan{ID: id}
	if err := repo.data.gormDB.Delete(&Plan).Error; err != nil {
		return "", nil
	}
	return Plan.ID, nil
}

func (repo *planRepo) Update(ctx context.Context, tbPlan *chdb.TbPlan) (res *chdb.TbPlan, err error) {
	u := new(chdb.TbPlan)
	err = repo.data.gormDB.Where("ID", tbPlan.ID).First(u).Error

	if err != nil {
		return nil, err
	}
	err = repo.data.gormDB.Model(&u).Updates(&chdb.TbPlan{
		ID:     tbPlan.ID,
		Name:   tbPlan.Name,
		Type:   tbPlan.Type,
		Status: tbPlan.Status,
	}).Error

	return &chdb.TbPlan{
		ID:         u.ID,
		Name:       u.Name,
		Type:       u.Type,
		Status:     u.Status,
		BelongType: u.BelongType,
		BelongTo:   u.BelongTo,
	}, nil
}

func (repo *planRepo) GetByID(ctx context.Context, id string) (*chdb.TbPlan, error) {
	tbPlan := chdb.TbPlan{}
	if err := repo.data.gormDB.Where("id = ?", id).First(&tbPlan).Error; err != nil {
		return nil, err
	}
	return &tbPlan, nil
}

func (repo *planRepo) Filter(ctx context.Context, apiReq *v1.PlanFilterQueryRequest) ([]chdb.TbPlan, int64, error) {
	planlist := make([]chdb.TbPlan, 0)
	var count int64 = 0

	filter := apiReq.Filter
	tx := repo.data.gormDB.Model(&chdb.TbPlan{})

	if filter != nil {
		if "" != filter.Key {
			keyLike := AddLikeCharToStr(filter.Key)
			tx.Where("name like ? or apply_disease like ? or apply_dysfunction like ?", keyLike, keyLike, keyLike)
		}
		if len(filter.Ids) > 0 {
			tx.Where("id in ?", filter.Ids)
		}

		//if "" != apiReq.HeaderParams.TokenInfo.AppID {
		//	tx.Where("app_id = ?", apiReq.HeaderParams.TokenInfo.AppID)
		//}
		if "" != filter.Status {
			tx.Where("status = ?", filter.Status)
		}
		if "" != filter.Type {
			tx.Where("type = ?", filter.Type)
		}
		if "" != filter.BelongType {
			tx.Where("belong_type = ?", filter.BelongType)
		}

		if "" != filter.BelongTo {
			tx.Where("belong_to = ?", filter.BelongTo)
		}
	}

	tx.Order("update_time DESC")
	tx.Count(&count)

	if apiReq.Page != -1 {
		tx.Limit(int(apiReq.PerPage)).Offset(int((apiReq.Page - 1) * apiReq.PerPage))
	}

	result := tx.Find(&planlist)
	return planlist, count, result.Error
}

func (repo *planRepo) PlanDetailPreloadById(ctx context.Context, id string) (*model.PlanDetailPreloadInfo, error) {
	planDetailPreloadInfo := model.PlanDetailPreloadInfo{}

	tx := repo.data.gormDB.Model(&model.PlanDetailPreloadInfo{}).Where("id = ?", id)
	tx.Preload("TbPlanRelates")
	if err := tx.First(&planDetailPreloadInfo).Error; err != nil {
		return nil, err
	}
	return &planDetailPreloadInfo, nil
}

func (repo *planRepo) PlanCreateWithResource(ctx context.Context, tbPlan *chdb.TbPlan, relates []chdb.TbPlanRelate,
	tbWorkItems []chdb.TbWorkItem, tbRelates []chdb.TbRelate) (*chdb.TbPlan, error) {
	if nil == tbPlan {
		return nil, Err.New("param plan is nil")
	}
	if err := repo.data.gormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(tbPlan).Error; err != nil {
			return err
		}
		if len(relates) > 0 {
			if err := tx.Create(relates).Error; err != nil {
				return err
			}
		}
		if len(tbWorkItems) > 0 {
			if err := tx.Create(tbWorkItems).Error; err != nil {
				return err
			}
		}
		if len(tbRelates) > 0 {
			if err := tx.Create(tbRelates).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return tbPlan, nil
}

func (repo *planRepo) PlanUpdateWithResource(ctx context.Context, planResource *model.ToUpdatePlanResource) (*chdb.TbPlan, error) {
	if nil == planResource {
		return nil, Err.New("param resource is nil")
	}
	if err := repo.data.gormDB.Transaction(func(tx *gorm.DB) error {
		//更新方案基本信息
		if err := tx.Save(planResource.Plan).Error; err != nil {
			return err
		}
		//写入数据
		if len(planResource.ToCreateWorkItem) > 0 {
			if err := tx.Create(planResource.ToCreateWorkItem).Error; err != nil {
				return err
			}
		}
		if len(planResource.ToCreateRelate) > 0 {
			if err := tx.Create(planResource.ToCreateRelate).Error; err != nil {
				return err
			}
		}
		if len(planResource.ToCreatePlanRelate) > 0 {
			if err := tx.Create(planResource.ToCreatePlanRelate).Error; err != nil {
				return err
			}
		}
		//更新数据
		for _, WorkItem := range planResource.ToUpdateWorkItem {
			if err := tx.Save(&WorkItem).Error; err != nil {
				return err
			}
		}
		for _, tbRelate := range planResource.ToUpdateRelate {
			if err := tx.Save(&tbRelate).Error; err != nil {
				return err
			}
		}
		for _, tbPlanRelate := range planResource.ToUpdatePlanRelate {
			if err := tx.Save(&tbPlanRelate).Error; err != nil {
				return err
			}
		}
		// 删除数据
		if len(planResource.ToDeleteWorkItemIds) > 0 {
			if err := tx.Where("id in ?", planResource.ToDeleteWorkItemIds).Delete(&chdb.TbWorkItem{}).Error; err != nil {
				return err
			}
		}
		if len(planResource.ToDeleteWorkItemRelateIds) > 0 {

			if err := tx.Where("id in ?", planResource.ToDeleteWorkItemRelateIds).Delete(&chdb.TbRelate{}).Error; err != nil {
				return err
			}
		}
		if len(planResource.ToDeletePlanRelate) > 0 {
			if err := tx.Where("id in ?", planResource.ToDeletePlanRelate).Delete(&chdb.TbPlanRelate{}).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return planResource.Plan, nil
}

func (repo *planRepo) SaPlanDetailPreloadById(ctx context.Context, id string) (*v1.SaPlanDetailPreload, error) {
	saPlanDetailPreload := v1.SaPlanDetailPreload{}

	tx := repo.data.gormDB.Model(&v1.SaPlanDetailPreload{}).Where("id = ?", id)
	tx.Preload("SaPlanRelates", "resource_type = ?", "WI") // 这里仅查询方案关联的工作项
	tx.Unscoped().Preload("SaPlanRelates.SaPlanWorkItemRelate")
	tx.Preload("SaPlanRelates.SaPlanWorkItemRelate.TbRelates")

	tx.Order("create_time")
	if err := tx.First(&saPlanDetailPreload).Error; err != nil {
		return nil, err
	}

	// 根据偏移进行排序
	sort.Slice(saPlanDetailPreload.SaPlanRelates, func(i, j int) bool {
		if saPlanDetailPreload.SaPlanRelates[i].SaPlanWorkItemRelate.WorkItem.Event > saPlanDetailPreload.SaPlanRelates[j].SaPlanWorkItemRelate.WorkItem.Event {
			return true
		}

		return saPlanDetailPreload.SaPlanRelates[i].SaPlanWorkItemRelate.WorkItem.Event == saPlanDetailPreload.SaPlanRelates[j].SaPlanWorkItemRelate.WorkItem.Event &&
			saPlanDetailPreload.SaPlanRelates[i].SaPlanWorkItemRelate.WorkItem.FrequencyInterval < saPlanDetailPreload.SaPlanRelates[j].SaPlanWorkItemRelate.WorkItem.FrequencyInterval
	})

	return &saPlanDetailPreload, nil
}

func (repo *planRepo) QueryPlanByProjectId(ctx context.Context, projectId string) (*chdb.TbPlan, error) {
	tbPlan := chdb.TbPlan{}
	if err := repo.data.gormDB.Model(&chdb.TbPlan{}).Where("belong_type = ? and belong_to = ?",
		protocol.PlanBelongTypeProject, projectId).First(&tbPlan).Error; err != nil {
		return nil, err
	}

	return &tbPlan, nil
}

func (repo *planRepo) PlanPreload(ctx context.Context, id string) (*model.PlanPreload, error) {
	planPreload := model.PlanPreload{}

	tx := repo.data.gormDB.Model(&model.PlanPreload{}).Where("id = ?", id)
	tx.Preload("TbPlanRelates")
	// 加载方案关联的工作项
	tx.Preload("RelateWorkItems", "resource_type = ?", "WI")
	tx.Preload("RelateWorkItems.PlanWorkItem")
	tx.Preload("RelateWorkItems.PlanWorkItem.TbRelates")
	// 加载方案关联的病种
	tx.Preload("RelateDiseases", "resource_type = ?", "BZ")
	tx.Preload("RelateDiseases.TbDisease", "id is not null")
	// 加载方案关联的功能障碍
	tx.Preload("RelateDysfunctions", "resource_type = ?", "GNZA")
	tx.Preload("RelateDysfunctions.TbDisFunc")

	if err := tx.First(&planPreload).Error; err != nil {
		return nil, err
	}

	// 对工作项进行排序
	sort.Slice(planPreload.RelateWorkItems, func(i, j int) bool {
		return planPreload.RelateWorkItems[i].PlanWorkItem.TbWorkItem.FrequencyInterval < planPreload.RelateWorkItems[j].PlanWorkItem.TbWorkItem.FrequencyInterval
	})

	return &planPreload, nil
}
