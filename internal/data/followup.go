package data

import (
	"context"
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
	PlanID            string
	AssignedType      string
	PatientID         string
	PatientName       string
	EventStartTime    string
	AssignedTo        string
	AssignedToName    string
	PlanStartTime     time.Time
	PlanEndTime       time.Time
	ActualStartTime   time.Time
	ActualEndTime     time.Time
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
	NotifyLeftOffset  int32
	NotifyRightOffset int32
	NotifyOffsetUnit  string
	NotifyLeftDate    time.Time
	NotifyRightDate   time.Time
	NotifyNode        string
	ExecArea          string
	Description       string
	CreateTime        time.Time
	UpdateTime        time.Time
	DeletedAt         time.Time
}
type FollowupRepo struct {
	data *Data
	log  *log.Helper
}

func NewFollowupRepo(data *Data, logger log.Logger) biz.FollowupRepo {
	return &FollowupRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *FollowupRepo) FollowupWiCreate(ctx context.Context, tbFollowupWi *chdb.TbWorkItem) (*chdb.TbWorkItem, error) {
	if err := r.data.gormDB.Model(&chdb.TbWorkItem{}).Create(tbFollowupWi).Error; err != nil {
		return nil, err
	}
	return tbFollowupWi, nil
}

func (r *FollowupRepo) FollowupWiDelById(ctx context.Context, id string) (string, error) {
	tbFollowupWi := chdb.TbWorkItem{ID: id}
	if err := r.data.gormDB.Delete(&tbFollowupWi).Error; err != nil {
		return "", err
	}
	return tbFollowupWi.ID, nil
}

func (r *FollowupRepo) FollowupWiDetailQueryById(ctx context.Context, id string) (*model.FollowupWiDetailPreload, error) {

	wiDetailPreload := model.FollowupWiDetailPreload{}
	tx := r.data.gormDB.Debug().Model(&model.FollowupWiDetailPreload{}).Where("id = ?", id)

	tx.Preload("TbRelate", "resource_type = ?", protocol.ResourceTypePlan)
	tx.Preload("FollowupWiChildren")
	//tx.Preload("FollowupWiChildren.ChildRelates")

	tx.Preload("FollowupWiChildren.WiRelatePreloads")
	tx.Preload("FollowupWiChildren.WiRelatePreloads.TbFormRow")
	tx.Preload("FollowupWiChildren.WiRelatePreloads.TbFormCSS")

	if err := tx.First(&wiDetailPreload).Error; err != nil {
		return nil, err
	}

	// 根据formRow查询所有此次任务相关的预警信息
	formRowIds := make([]string, 0)
	for _, child := range wiDetailPreload.FollowupWiChildren {
		for _, wiRelatePreload := range child.WiRelatePreloads {
			if nil == wiRelatePreload.TbFormRow {
				continue
			}
			if "" != wiRelatePreload.TbFormRow.ID {
				formRowIds = append(formRowIds, wiRelatePreload.TbFormRow.ID)
			}
		}
	}
	// 通过formRowId查询预警信息
	formWarnings := make([]chdb.TbFormWarning, 0)
	if len(formRowIds) > 0 {
		if err := r.data.gormDB.Model(&chdb.TbFormWarning{}).Where("tb_form_row_id in ?", formRowIds).Find(&formWarnings).Error; err != nil {
			return nil, err
		}
	}
	// 根据sortNum排序

	sort.Slice(wiDetailPreload.FollowupWiChildren, func(i, j int) bool {
		return wiDetailPreload.FollowupWiChildren[i].TbWorkItem.SortNum < wiDetailPreload.FollowupWiChildren[j].TbWorkItem.SortNum
	})
	for i := 0; i < len(wiDetailPreload.FollowupWiChildren); i++ {
		for j := 0; j < len(wiDetailPreload.FollowupWiChildren[i].WiRelatePreloads); j++ {
			if nil == wiDetailPreload.FollowupWiChildren[i].WiRelatePreloads[j].TbFormRow {
				continue
			}

			if "" != wiDetailPreload.FollowupWiChildren[i].WiRelatePreloads[j].TbFormRow.ID {
				for _, formWarning := range formWarnings {
					if formWarning.TbFormRowID == wiDetailPreload.FollowupWiChildren[i].WiRelatePreloads[j].TbFormRow.ID {
						if nil == wiDetailPreload.FollowupWiChildren[i].WiRelatePreloads[j].TbFormWarnings {
							wiDetailPreload.FollowupWiChildren[i].WiRelatePreloads[j].TbFormWarnings = make([]chdb.TbFormWarning, 0)
						}
						wiDetailPreload.FollowupWiChildren[i].WiRelatePreloads[j].TbFormWarnings = append(wiDetailPreload.FollowupWiChildren[i].WiRelatePreloads[j].TbFormWarnings, formWarning)
					}
				}
			}
		}
	}

	return &wiDetailPreload, nil
}

// FollowupWiFilterQuery 随访工作项条件查询
func (r *FollowupRepo) FollowupWiFilterQuery(ctx context.Context, apiReq *v1.FollowupFilterRequest) ([]model.FollowupWiFilterPreload, int64, error) {
	tbFollowupWis := make([]model.FollowupWiFilterPreload, 0)
	var count int64 = 0

	filter := apiReq.Filter

	// 仅条件查询
	tx := r.data.gormDB.Model(&model.FollowupWiFilterPreload{}).Where("type = ?", protocol.WorkItemTypeFollowUp)
	// 此处仅返回根任务，不返回子任务
	tx.Where("pid is NULL")
	// 过滤掉方案模板和草稿
	notInStatus := make([]string, 0)
	notInStatus = append(notInStatus, protocol.WorkItemStatusModel)
	notInStatus = append(notInStatus, protocol.WorkItemStatusDraft)
	tx.Where("status not in ?", notInStatus)

	if filter != nil {
		if "" != filter.Key {
			keyLike := AddLikeCharToStr(filter.Key)
			tx.Where("title like ? or principal_name like ? or assigned_to_name like ?", keyLike, keyLike, keyLike)
		}
		if "" != filter.PatientId {
			tx.Where("assigned_to = ?", filter.PatientId)
		}

		// 任务预计开始时间条件查询
		// 计划开始时间过滤
		planStartTime := filter.PlanStartTime
		if nil != planStartTime && ("" != planStartTime.Start || "" != planStartTime.End) {
			if "" != planStartTime.Start && "" != planStartTime.End {
				tx.Where("plan_start_time is not null and plan_start_time between ? and ?",
					planStartTime.Start, planStartTime.End)
			}
			if "" != planStartTime.Start && "" == planStartTime.End {
				tx.Where("plan_start_time is not null and unix_timestamp(plan_start_time) > unix_timestamp(?)",
					planStartTime.Start)
			}
			if "" == planStartTime.Start && "" != planStartTime.End {
				tx.Where("plan_start_time is not null and unix_timestamp(plan_start_time) < unix_timestamp(?)",
					planStartTime.End)
			}
		}
	}

	tx.Preload("TbRelate")
	//tx.Preload("TbPatient")

	tx.Order("plan_start_time")
	tx.Count(&count)
	if apiReq.Page != -1 {
		tx.Limit(int(apiReq.PerPage)).Offset(int((apiReq.Page - 1) * apiReq.PerPage))
	}
	if err := tx.Find(&tbFollowupWis).Error; err != nil {
		return nil, 0, err
	}

	return tbFollowupWis, count, nil
}

func (r *FollowupRepo) QueryFollowupWiChildrenById(id string, recursion bool) ([]chdb.TbWorkItem, error) {
	workItems := make([]chdb.TbWorkItem, 0)
	if recursion {
		// 递归查询子集
		if err := r.data.gormDB.Raw("WITH RECURSIVE t as ( select * from tb_work_item where id = ? UNION ALL select tmp.* from tb_work_item tmp join t on tmp.pid = t.id ) SELECT * FROM t;",
			id).Scan(&workItems).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return nil, err
			}
		}
	} else {
		if err := r.data.gormDB.Model(&chdb.TbWorkItem{}).Where("pid = ?", id).Find(&workItems).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return nil, err
			}
		}
	}

	return workItems, nil
}

func (r *FollowupRepo) UpdateFollowupWiById(tbFollowupWi *chdb.TbWorkItem) (*chdb.TbWorkItem, error) {
	if err := r.data.gormDB.Save(tbFollowupWi).Error; err != nil {
		return nil, err
	}
	return tbFollowupWi, nil
}

func (r *FollowupRepo) FollowupWiDelByIdRecursion(ctx context.Context, id string) ([]string, error) {

	children, err := r.QueryFollowupWiChildrenById(id, true)
	if err != nil {
		return nil, err
	}

	toDelIds := make([]string, 0)
	for _, child := range children {
		if "" == child.ID {
			continue
		}
		toDelIds = append(toDelIds, child.ID)
	}

	if len(toDelIds) == 0 {
		return toDelIds, nil
	}

	if err := r.data.gormDB.Where("id in ?", toDelIds).Delete(&chdb.TbWorkItem{}).Error; err != nil {
		return nil, err
	}
	return toDelIds, nil
}
