package model

import "followup/gencode/chdb"

type WorkItemDetailBean struct {
	chdb.TbWorkItem
	TbRelates []*chdb.TbRelate `gorm:"foreignKey:tb_work_item_id"`
}

type WorkItemPreloadInfo struct {
	TbWorkItem chdb.TbWorkItem
	TbRelate   chdb.TbRelate
}
