package chdb

import "time"

// TbRole 角色表
type TbRole struct {
	ID          string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'角色编码'"`                      // 角色编码
	Name        string     `gorm:"column:name;type:varchar(45);not null;comment:'角色名称'"`                               // 角色名称
	Level       string     `gorm:"column:level;type:enum('SYSTEM','ORG','USER');not null;comment:'角色级别(SYSTEM,USER)'"` // 角色级别(SYSTEM,USER)
	Status      string     `gorm:"column:status;type:enum('ENABLE','DISABLE');not null;comment:'角色状态(启用,停用)'"`         // 角色状态(启用,停用)
	Tag         string     `gorm:"column:tag;type:varchar(45);default:null;comment:'备用'"`                              // 备用
	Description string     `gorm:"column:description;type:varchar(255);default:null;comment:'角色描述'"`                   // 角色描述
	TbOrgID     string     `gorm:"column:tb_org_id;type:varchar(45);not null"`
	CreateTime  *time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"` // 创建时间
	UpdateTime  *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *TbRole) TableName() string {
	return "tb_role"
}
