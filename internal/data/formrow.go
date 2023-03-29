package data

import (
	"context"
	v1 "followup/api"
	"followup/gencode/chdb"
	"followup/internal/biz"
	"followup/protocol"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type FormRowRepo struct {
	data *Data
	log  *log.Helper
}

func NewFormRowRepo(data *Data, logger log.Logger) biz.FormRowRepo {
	return &FormRowRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *FormRowRepo) Save(ctx context.Context, tbFormRow *chdb.TbFormRow) (*chdb.TbFormRow, error) {
	// 查询关联数据
	tbRelate := chdb.TbRelate{}
	if err := r.data.gormDB.Model(&chdb.TbRelate{}).Where("id = ?", tbFormRow.TbRelateID).First(&tbRelate).Error; err != nil {
		return nil, err
	}
	// 查询工作项数据
	tbWorkItem := chdb.TbWorkItem{}
	if err := r.data.gormDB.Model(&chdb.TbWorkItem{}).Where("id = ?", tbRelate.TbWorkItemID).First(&tbWorkItem).Error; err != nil {
		return nil, err
	}

	if err := r.data.gormDB.Transaction(func(tx *gorm.DB) error {
		// 写入表单行数据
		if err := tx.Create(tbFormRow).Error; err != nil {
			return err
		}
		// 写入表单预警信息
		//if len(tbColumnWarnings) > 0 {
		//	if err := tx.Create(&tbColumnWarnings).Error; err != nil {
		//		return err
		//	}
		//}
		// 更新关联项状态为正在处理
		tbRelate.Status = protocol.RelateStatusActive
		if err := tx.Save(&tbRelate).Error; err != nil {
			return err
		}
		// 更新工作项状态
		tbWorkItem.Status = protocol.WorkItemStatusActive

		if err := tx.Save(&tbWorkItem).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return tbFormRow, nil
}

func (r *FormRowRepo) DeleteByID(ctx context.Context, id string) (string, error) {
	tbFormRow := chdb.TbFormRow{ID: id}
	if err := r.data.gormDB.Delete(&tbFormRow).Error; err != nil {
		return "", err
	}
	return tbFormRow.ID, nil
}

func (r *FormRowRepo) Update(ctx context.Context, req *chdb.TbFormRow) (res *chdb.TbFormRow, err error) {
	tbFormRow := new(chdb.TbFormRow)
	err = r.data.gormDB.Where("ID", req.ID).First(tbFormRow).Error

	if err != nil {
		return nil, err
	}
	err = r.data.gormDB.Model(&tbFormRow).Updates(&chdb.TbFormRow{
		ID:            req.ID,
		Title:         req.Title,
		Status:        req.Status,
		Data:          req.Data,
		Remark:        req.Remark,
		Submitter:     req.Submitter,
		TbFormID:      req.TbFormID,
		TbRelateID:    req.TbRelateID,
		SignaturePath: req.SignaturePath,
		SignerID:      req.SignerID,
		SignerName:    req.SignerName,
		SignOnBehalf:  req.SignOnBehalf,
		CreateTime:    req.CreateTime,
		UpdateTime:    req.UpdateTime,
	}).Error

	return &chdb.TbFormRow{
		ID:            tbFormRow.ID,
		Title:         tbFormRow.Title,
		Status:        tbFormRow.Status,
		Data:          tbFormRow.Data,
		Remark:        tbFormRow.Remark,
		Submitter:     tbFormRow.Submitter,
		TbFormID:      tbFormRow.TbFormID,
		TbRelateID:    tbFormRow.TbRelateID,
		SignaturePath: tbFormRow.SignaturePath,
		SignerID:      tbFormRow.SignerID,
		SignerName:    tbFormRow.SignerName,
		SignOnBehalf:  tbFormRow.SignOnBehalf,
		CreateTime:    tbFormRow.CreateTime,
		UpdateTime:    GetNowTimeAddr(),
	}, nil
}

func (r *FormRowRepo) GetByID(ctx context.Context, id string) (*chdb.TbFormRow, error) {
	tbFormRow := chdb.TbFormRow{}
	if err := r.data.gormDB.Where("id = ?", id).First(&tbFormRow).Error; err != nil {
		return nil, err
	}
	return &tbFormRow, nil
}

func (r *FormRowRepo) Filter(ctx context.Context, req *v1.FormRowFilterRequest) ([]chdb.TbFormRow, int64, error) {
	tbFormRows := make([]chdb.TbFormRow, 0)
	var count int64 = 0

	filter := req.Filter

	// 条件查询
	tx := r.data.gormDB.Model(&chdb.TbFormRow{})

	if filter != nil {
		if "" != filter.Key {
			keyLike := AddLikeCharToStr(filter.Key)
			tx.Where("title like ?", keyLike)
		}
		if "" != filter.TbRelateId {
			tx.Where("tb_relate_id = ?", filter.TbRelateId)
		}
		if "" != filter.TbFormId {
			tx.Where("tb_form_id = ?", filter.TbFormId)
		}
		if len(filter.Ids) > 0 {
			tx.Where("id in ?", filter.Ids)
		}
	}
	tx.Order("create_time DESC")
	tx.Count(&count)

	if req.Page != -1 {
		tx.Limit(int(req.PerPage)).Offset(int((req.Page - 1) * req.PerPage))
	}
	if err := tx.Find(&tbFormRows).Error; err != nil {
		return nil, 0, err
	}
	return tbFormRows, count, nil
}
