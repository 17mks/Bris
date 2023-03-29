package model

import "followup/gencode/chdb"

// UserDetailInfo 用户详情，外键关联; foreignKey表示当前表中的哪个字段与TbUser的主键ID对应
type UserDetailInfo struct {
	chdb.TbUser
	TbUserInfo      chdb.TbUserInfo       `gorm:"foreignKey:tb_user_id"`
	TbThirdAccounts []chdb.TbThirdAccount `gorm:"foreignKey:tb_user_id"`

	UserRolePreloadInfos []UserRolePreloadInfo `gorm:"foreignKey:body_id"` // 用于查询用户角色
}

type UserRolePreloadInfo struct {
	chdb.TbResourceBind
	TbRole chdb.TbRole `gorm:"foreignKey:id;references:ResourceID"`
}
