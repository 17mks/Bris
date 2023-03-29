package data

import (
	"context"
	"errors"
	"fmt"
	"followup/gencode/chdb"
	"followup/internal/biz"
	"followup/model"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type WorkItemRepo struct {
	data *Data
	log  *log.Helper
}

func NewWorkItemRepo(data *Data, logger log.Logger) biz.WorkItemRepo {
	return &WorkItemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *WorkItemRepo) QueryPlanWorkItemByIds(ctx context.Context, ids []string) ([]model.WorkItemDetailBean, error) {
	tbWorkItems := make([]model.WorkItemDetailBean, 0)
	if len(ids) == 0 {
		return tbWorkItems, nil
	}

	tx := r.data.gormDB.Model(&model.WorkItemDetailBean{}).Where("id in ?", ids)

	tx.Preload("TbRelates")

	if err := tx.Find(&tbWorkItems).Error; err != nil {
		return nil, err
	}
	return tbWorkItems, nil
}

func (r *WorkItemRepo) WorkItemCreateWithRelates(ctx context.Context, tbWorkItems []chdb.TbWorkItem, tbRelates []chdb.TbRelate) error {
	if err := r.data.gormDB.Transaction(func(tx *gorm.DB) error {
		if len(tbWorkItems) > 0 {
			if err := tx.Create(&tbWorkItems).Error; err != nil {
				return err
			}
		}
		if len(tbRelates) > 0 {
			if err := tx.Create(&tbRelates).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *WorkItemRepo) WorkItemPreloadByRelateId(ctx context.Context, relateId string) (*model.WorkItemPreloadInfo, error) {
	if "" == relateId {
		return nil, errors.New("param relateId is nil")
	}
	workItemPreloadInfo := model.WorkItemPreloadInfo{}
	// 查询关联信息
	tbRelate := chdb.TbRelate{}
	if err := r.data.gormDB.First(&tbRelate, "id = ?", relateId).Error; err != nil {
		return nil, fmt.Errorf("relate '%s' not exists", relateId)
	}
	workItemPreloadInfo.TbRelate = tbRelate
	// 查询工作项信息
	detailBean, err := r.WorkItemDetailPreloadById(ctx, tbRelate.TbWorkItemID)
	if err != nil {
		return nil, err
	}
	workItemPreloadInfo.TbWorkItem = detailBean.TbWorkItem
	// TODO 查询工作项相关人员信息

	return &workItemPreloadInfo, nil
}

func (r *WorkItemRepo) WorkItemDetailPreloadById(ctx context.Context, id string) (*model.WorkItemDetailBean, error) {
	tbWorkItem := model.WorkItemDetailBean{}

	tx := r.data.gormDB.Model(&model.WorkItemDetailBean{}).Where("id = ?", id)

	tx.Preload("TbRelates")

	if err := tx.First(&tbWorkItem).Error; err != nil {
		return nil, err
	}
	return &tbWorkItem, nil
}
