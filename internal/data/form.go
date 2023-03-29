package data

import (
	"context"
	v1 "followup/api"
	"followup/gencode/chdb"
	"followup/internal/biz"
	"followup/model"
	"followup/protocol"
	"followup/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	Log "log"
	"time"
)

// TbForm 病种表
type TbForm struct {
	ID          string    `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'表单模板编码'"`                                                  // 表单模板编码
	Name        string    `gorm:"column:name;type:varchar(256);not null;comment:'表单模板名称'"`                                                          // 表单模板名称
	Type        string    `gorm:"column:type;type:enum('SAF','IC','UP','SP');not null;comment:'表单类型('SAF 症状评估表', 'IC 知情同意书','UP 用户协议','SP 隐私协议')'"` // 表单类型('SAF 症状评估表', 'IC 知情同意书','UP 用户协议','SP 隐私协议')
	Status      string    `gorm:"column:status;type:enum('DESIGNING','ENABLE','DISABLE');not null;default:DESIGNING;comment:'表单模板状态'"`              // 表单模板状态
	BranchLogic bool      `gorm:"column:branch_logic;type:tinyint(1);default:null;default:0;comment:'是否启用分支逻辑功能'"`                                  // 是否启用分支逻辑功能
	BelongType  string    `gorm:"column:belong_type;type:enum('ORG','PROJECT','TEAM','PLAN','NONE');default:null;comment:'归属类型'"`                   // 归属类型
	BelongTo    string    `gorm:"column:belong_to;type:varchar(45);default:null;comment:'归属资源编码'"`                                                  // 归属资源编码
	Description string    `gorm:"column:description;type:varchar(512);default:null;comment:'表单模板描述'"`                                               // 表单模板描述
	AppID       string    `gorm:"column:appid;type:varchar(512);default:null;comment:'应用编码'"`
	VersionName string    `gorm:"column:version_name;type:varchar(512);default:null;comment:'版本名称'"`
	VersionCode int32     `gorm:"column:version_code;type:int(11);default:null;comment:'版本号'"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	DeleteAt    time.Time `gorm:"column:delete_at;type:datetime;default:not null;default:CURRENT_TIMESTAMP;comment:'删除时间'"`
}

type FormRepo struct {
	data *Data
	log  *log.Helper
}

func NewFormRepo(data *Data, logger log.Logger) biz.FormRepo {
	return &FormRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *FormRepo) Save(ctx context.Context, req *chdb.TbForm) error {
	Form := chdb.TbForm{
		ID:          req.ID,
		Name:        req.Name,
		Status:      req.Status,
		Type:        req.Type,
		BelongTo:    req.BelongTo,
		BelongType:  req.BelongType,
		BranchLogic: req.BranchLogic,
		VersionCode: req.VersionCode,
		VersionName: req.VersionName,
		Description: req.Description,
	}
	//验证是否已创建
	result := r.data.gormDB.Where(&chdb.TbForm{ID: req.ID}).First(&Form)

	if result.RowsAffected == 1 {
		return status.Errorf(codes.AlreadyExists, "该病种已存在")
	}
	Form.ID = req.ID
	Form.Name = req.Name
	Form.Description = req.Description
	Form.Status = req.Status
	Form.Type = req.Type
	Form.BelongTo = req.BelongTo
	Form.BelongType = req.BelongType
	Form.BranchLogic = req.BranchLogic
	Form.VersionCode = req.VersionCode
	Form.VersionName = req.VersionName
	res := r.data.gormDB.Create(&Form)
	return res.Error
}

func (r *FormRepo) Delete(ctx context.Context, id string) (string, error) {
	Form := chdb.TbForm{ID: id}
	if err := r.data.gormDB.Delete(&Form).Error; err != nil {
		return "", nil
	}
	return Form.ID, nil
}

func (r *FormRepo) Update(ctx context.Context, req *chdb.TbForm) (res *chdb.TbForm, err error) {
	Form := new(chdb.TbForm)
	err = r.data.gormDB.Where("ID", req.ID).First(Form).Error

	if err != nil {
		return nil, err
	}
	err = r.data.gormDB.Model(&Form).Updates(&chdb.TbForm{
		ID:          req.ID,
		Name:        req.Name,
		Status:      req.Status,
		Type:        req.Type,
		BelongTo:    req.BelongTo,
		BelongType:  req.BelongType,
		BranchLogic: req.BranchLogic,
		VersionCode: req.VersionCode,
		VersionName: req.VersionName,
		Description: req.Description,
	}).Error

	return &chdb.TbForm{
		ID:          Form.ID,
		Name:        Form.Name,
		Status:      Form.Status,
		Type:        Form.Type,
		BelongTo:    Form.BelongTo,
		BelongType:  Form.BelongType,
		BranchLogic: Form.BranchLogic,
		VersionCode: Form.VersionCode,
		VersionName: Form.VersionName,
		Description: Form.Description,
		CreateTime:  req.CreateTime,
		UpdateTime:  req.UpdateTime,
		DeleteAt:    req.DeleteAt,
	}, nil
}

func (r *FormRepo) GetByID(ctx context.Context, id string) (rv *chdb.TbForm, errs error) {
	tbForm := new(chdb.TbForm)
	if err := r.data.gormDB.Where("id = ?", id).First(&tbForm).Error; err != nil {
		return nil, err
	}
	return &chdb.TbForm{
		ID:          tbForm.ID,
		Name:        tbForm.Name,
		Status:      tbForm.Status,
		Type:        tbForm.Type,
		BelongTo:    tbForm.BelongTo,
		BelongType:  tbForm.BelongType,
		BranchLogic: tbForm.BranchLogic,
		VersionCode: tbForm.VersionCode,
		VersionName: tbForm.VersionName,
		Description: tbForm.Description,
		CreateTime:  tbForm.CreateTime,
		UpdateTime:  tbForm.UpdateTime,
		DeleteAt:    tbForm.DeleteAt,
	}, nil
}

func (r *FormRepo) Filter(ctx context.Context, apiReq *v1.FormFilterRequest) ([]model.FormDetailPreloadInfo, int64, error) {
	serverContext, _ := transport.FromServerContext(ctx)

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	Log.Println(Token)

	tbForms := make([]model.FormDetailPreloadInfo, 0)
	var count int64 = 0

	filter := apiReq.Filter

	// 仅条件查询
	tx := r.data.gormDB.Model(&model.FormDetailPreloadInfo{})

	if filter != nil {
		if "" != filter.Key {
			keyLike := protocol.AddLikeCharToStr(filter.Key)
			tx.Where("name like ?", keyLike)
		}
		if len(filter.Ids) > 0 {
			tx.Where("id in ?", filter.Ids)
		}
		if len(filter.Status) > 0 {
			tx.Where("status in ?", filter.Status)
		}
		if "" != filter.Type {
			tx.Where("type = ?", filter.Type)
		}

		//if "" != apiReq.TokenInfo.AppId {
		//	tx.Where("app_id = ?", apiReq.TokenInfo.AppId)
		//}
		if len(filter.NotIn.Types) > 0 {
			tx.Where("type not in ?", filter.NotIn.Types)
		}
	}

	//if "" != filter.NotIn.ProjectId {
	//	// 过滤掉已和项目绑定的表单
	//	proFormIds, err := r.QueryProjectBindFormIds(filter.NotIn.ProjectId)
	//	if err != nil {
	//		return nil, 0, err
	//	}
	//	if len(proFormIds) == 0 {
	//		return tbForms, 0, nil
	//	}
	//	tx.Where("id not in ?", proFormIds)
	//
	//}

	tx.Order("update_time DESC")
	tx.Count(&count)
	if apiReq.Page != -1 {
		tx.Limit(int(apiReq.PerPage)).Offset(int((apiReq.Page - 1) * apiReq.PerPage))
	}

	tx.Preload("TbFormCSS")
	tx.Preload("TbFormColumns")
	if err := tx.Find(&tbForms).Error; err != nil {
		return nil, 0, err
	}

	return tbForms, count, nil

}

func (r *FormRepo) Import(ctx context.Context, formInfo *model.FormInfo) (*model.FormInfo, error) {
	if err := r.data.gormDB.Transaction(func(tx *gorm.DB) error {
		// 表单
		if err := tx.Create(&formInfo.TbForm).Error; err != nil {
			return err
		}
		// 表单样式
		if "" != formInfo.TbFormCSS.ID {
			if err := tx.Create(&formInfo.TbFormCSS).Error; err != nil {
				return err
			}
		}

		for _, formColumnInfo := range formInfo.FormColumnInfos {
			// 表单字段
			if err := tx.Create(&formColumnInfo.TbFormColumn).Error; err != nil {
				return err
			}
			// 字段选项
			if len(formColumnInfo.TbColumnOpts) > 0 {
				if err := tx.Create(&formColumnInfo.TbColumnOpts).Error; err != nil {
					return err
				}
			}
			// 字段预警
			if len(formColumnInfo.TbColumnThresholds) > 0 {
				if err := tx.Create(&formColumnInfo.TbColumnThresholds).Error; err != nil {
					return err
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return formInfo, nil
}

func (r *FormRepo) FormInfoPreloadByIds(ctx context.Context, ids []string, projectId string) ([]model.FormInfo, error) {
	formInfos := make([]model.FormInfo, 0)
	if len(ids) == 0 {
		return formInfos, nil
	}
	// 查询表单详情
	detailPreloadInfos := make([]model.FormDetailPreloadInfo, 0)
	tx := r.data.gormDB.Model(&model.FormDetailPreloadInfo{}).Where("id in ?", ids)
	tx.Preload("TbFormCSS")
	tx.Preload("TbFormColumns")
	if err := tx.Find(&detailPreloadInfos).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return formInfos, nil
		}
		return nil, err
	}
	// 查询表单字段详情
	for _, formDetailPreloadInfo := range detailPreloadInfos {
		formInfo := model.FormInfo{
			TbForm:          formDetailPreloadInfo.TbForm,
			TbFormCSS:       formDetailPreloadInfo.TbFormCSS,
			FormColumnInfos: make([]model.FormColumnInfo, 0),
		}
		// 查询表单模板字段详情
		columnIds := make([]string, 0)
		for _, tbFormColumn := range formDetailPreloadInfo.TbFormColumns {
			columnIds = append(columnIds, tbFormColumn.ID)
		}
		if len(columnIds) > 0 {
			formColumnInfos := make([]model.FormColumnInfo, 0)
			tx := r.data.gormDB.Model(&model.FormColumnInfo{}).Where("id in ?", columnIds)
			tx.Preload("TbColumnOpts")
			if "" == projectId {
				tx.Preload("TbColumnThresholds", "tb_project_id is null")
			} else {
				tx.Preload("TbColumnThresholds", "tb_project_id = ?", projectId)
			}

			tx.Order("sort_index")
			if err := tx.Find(&formColumnInfos).Error; err != nil {
				if err != gorm.ErrRecordNotFound {
					return nil, err
				}
			}
			if len(formColumnInfos) > 0 {
				formInfo.FormColumnInfos = append(formInfo.FormColumnInfos, formColumnInfos...)
			}
		}
		formInfos = append(formInfos, formInfo)
	}
	return formInfos, nil
}

func (r *FormRepo) FormQueryByIds(ctx context.Context, ids []string) ([]chdb.TbForm, error) {
	tbForms := make([]chdb.TbForm, 0)
	if err := r.data.gormDB.Model(&chdb.TbForm{}).Where("id in ?", ids).Find(&tbForms).Error; err != nil {
		return nil, err
	}
	return tbForms, nil
}

func (r *FormRepo) FormPseudoDelById(ctx context.Context, id string) (string, error) {
	if err := r.data.gormDB.Model(&chdb.TbForm{}).Where("id = ?", id).Update("delete_at", time.Now()).Error; err != nil {
		return id, err
	}
	return id, nil
}

func (r *FormRepo) QueryLatestFormByName(ctx context.Context, name string) (*chdb.TbForm, error) {
	tbForm := chdb.TbForm{}

	if err := r.data.gormDB.Model(&chdb.TbForm{}).Where("name = ?", name).Order("version_code desc").First(&tbForm).Error; err != nil {
		return nil, err
	}
	return &tbForm, nil
}

func (r *FormRepo) QueryFormByType(ctx context.Context, fType string) ([]chdb.TbForm, error) {
	tbForms := make([]chdb.TbForm, 0)

	if err := r.data.gormDB.Model(&v1.FormBean{}).Where("type = ?", fType).Find(&tbForms).Error; err != nil {
		return nil, err
	}
	return tbForms, nil
}

func (r *FormRepo) FormDetailPreloadById(ctx context.Context, id string) (*model.FormDetailPreloadInfo, error) {
	detailPreloadInfo := model.FormDetailPreloadInfo{}

	tx := r.data.gormDB.Model(&model.FormDetailPreloadInfo{}).Where("id = ?", id)
	tx.Preload("TbFormColumns")
	tx.Preload("TbFormCSS")
	if err := tx.First(&detailPreloadInfo).Error; err != nil {
		return nil, err
	}
	return &detailPreloadInfo, nil
}

func (r *FormRepo) FormInfoPreloadById(ctx context.Context, id string, projectId string) (*model.FormInfo, error) {

	formInfo := model.FormInfo{}
	// 查询表单详情
	tx := r.data.gormDB.Debug().Model(&model.FormInfo{}).Where("id = ?", id)
	tx.Preload("TbFormCSS")
	tx.Preload("FormColumnInfos", "sort_index is not null order by sort_index")
	//tx.Preload("FormColumnInfos")
	tx.Preload("FormColumnInfos.TbColumnOpts")

	if "" == projectId {
		tx.Preload("FormColumnInfos.TbColumnThresholds", "tb_project_id is null or tb_project_id = ''")
	} else {
		// 查询项目与表单绑定ID，取过滤表单在这个项目下的预警规则
		proFormId := ""
		if err := r.data.gormDB.Model(&chdb.TbResourceBind{}).Select("id").Where("body_id = ? and resource_id = ?",
			projectId, id).First(&proFormId).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return nil, err
			}
		}
		if "" != proFormId {
			tx.Preload("FormColumnInfos.TbColumnThresholds", "tb_project_id = ?", proFormId)
		}

	}

	if err := tx.First(&formInfo).Error; err != nil {
		return nil, err
	}

	return &formInfo, nil

	//if "" == id {
	//	return nil, fmt.Errorf("param is empty")
	//}
	//formInfos, err := r.FormInfoPreloadByIds(ctx, []string{id}, projectId)
	//if err != nil {
	//	return nil, err
	//}
	//if len(formInfos) != 1 {
	//	return nil, fmt.Errorf("form id '%s' not exits", id)
	//}
	//return &formInfos[0], nil
}

//func (r *FormRepo) QueryProjectBindFormIds(ctx context.Context, projectId string) ([]string, error) {
//	tbResourceBinds, err := r.resource.ResourcesQuery(ctx, protocol.BodyTypeProject, projectId, protocol.ResourceTypeForm, "")
//	if err != nil {
//		return nil, err
//	}
//	proFormIds := make([]string, 0)
//	for _, tbResourceBind := range tbResourceBinds {
//		proFormIds = append(proFormIds, tbResourceBind.ResourceID)
//	}
//
//	return proFormIds, nil
//
//}
