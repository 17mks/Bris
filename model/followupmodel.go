package model

import (
	v1 "followup/api"
	"followup/gencode/chdb"
)

type FollowupWiCreateApiReq struct {
	HeaderParams
	Data v1.FollowupFilterRequest
}

type FollowupWiDelApiReq struct {
	HeaderParams
}
type FollowupWiDetailApiReq struct {
	HeaderParams
}
type FollowupWiFilterApiReq struct {
	HeaderParams
	Data v1.FollowupFilterRequest
}

type FollowupWiFilterPreload struct {
	chdb.TbWorkItem
	TbRelate chdb.TbRelate `gorm:"foreignKey:tb_work_item_id"`
	//TbPatient chdb.TbPatient `gorm:"foreignKey:id;references:AssignedTo"`
}

// FollowupWiDetailPreload 随访工作项详情预加载
type FollowupWiDetailPreload struct {
	chdb.TbWorkItem
	TbRelate           chdb.TbRelate     `gorm:"foreignKey:tb_work_item_id"` // 随访工作项的方案管理 - 目前一个随访关联一个方案
	FollowupWiChildren []FollowupWiChild `gorm:"foreignKey:pid;references:ID"`
}

type FollowupWiChild struct {
	chdb.TbWorkItem
	WiRelatePreloads []WiRelatePreload `gorm:"foreignKey:tb_work_item_id;references:ID"`
}
type WiRelatePreload struct {
	chdb.TbRelate
	TbFormRow      *chdb.TbFormRow      `gorm:"foreignKey:tb_relate_id"`
	TbFormCSS      *chdb.TbFormCSS      `gorm:"foreignKey:tb_form_id;references:ResourceID"`
	TbFormWarnings []chdb.TbFormWarning `gorm:"foreignKey:tb_form_row_id"`
}
