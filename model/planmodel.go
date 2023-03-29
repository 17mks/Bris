package model

import "followup/gencode/chdb"

// PlanDetailPreloadInfo 方案详情
type PlanDetailPreloadInfo struct {
	chdb.TbPlan
	TbPlanRelates []chdb.TbPlanRelate `gorm:"foreignKey:tb_plan_id"`
}

// ToUpdatePlanResource 方案待更新资源信息
type ToUpdatePlanResource struct {
	Plan *chdb.TbPlan // 待更新方案信息

	ToCreateWorkItem    []chdb.TbWorkItem // 待写入的工作项
	ToDeleteWorkItemIds []string          // 待删除的工作项
	ToUpdateWorkItem    []chdb.TbWorkItem // 待删除的工作项

	ToCreateRelate            []chdb.TbRelate // 待写入的工作项关联
	ToDeleteWorkItemRelateIds []string        // 待删除的工作项关联
	ToUpdateRelate            []chdb.TbRelate // 待更新的工作项关联

	ToCreatePlanRelate []chdb.TbPlanRelate // 待写入的方案关联
	ToDeletePlanRelate []string            // 待删除的方案关联
	ToUpdatePlanRelate []chdb.TbPlanRelate // 待更新的方案关联
}

type PlanPreload struct {
	chdb.TbPlan
	RelateWorkItems    []PlanRelateWorkItem    `gorm:"foreignKey:tb_plan_id"` // 方案关联的工作项
	RelateDiseases     []PlanRelateDisease     `gorm:"foreignKey:tb_plan_id"` // 方案关联的病种
	RelateDysfunctions []PlanRelateDysfunction `gorm:"foreignKey:tb_plan_id"` // 方案关联的功能障碍
	TbPlanRelates      []chdb.TbPlanRelate     `gorm:"foreignKey:tb_plan_id"`
}

type PlanRelateWorkItem struct {
	chdb.TbPlanRelate
	PlanWorkItem PlanWorkItem `gorm:"foreignKey:id;references:ResourceID"`
}

type PlanWorkItem struct {
	chdb.TbWorkItem
	TbRelates []chdb.TbRelate `gorm:"foreignKey:tb_work_item_id"`
}

type PlanRelateDisease struct {
	chdb.TbPlanRelate
	TbDisease *chdb.TbDisease `gorm:"foreignKey:id;references:ResourceID"`
}
type PlanRelateDysfunction struct {
	chdb.TbPlanRelate
	TbDisFunc *chdb.TbDisFunc `gorm:"foreignKey:id;references:ResourceID"`
}
