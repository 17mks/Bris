package chdb

import "time"

// TbProject 项目表
type TbProject struct {
	ID                string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'项目编码'"`                          // 项目编码
	Name              string     `gorm:"column:name;type:varchar(255);default:null;comment:'项目名称'"`                              // 项目名称
	Description       string     `gorm:"column:description;type:varchar(255);default:null;comment:'项目描述'"`                       // 项目描述
	Creator           string     `gorm:"column:creator;type:varchar(45);not null;comment:'项目创建者'"`                               // 项目创建者
	InformedConsentID string     `gorm:"column:informed_consent_id;type:varchar(45);default:null;comment:'知情同意书表单编码'"`           // 知情同意书表单编码
	CreateTime        *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"` // 创建时间
	UpdateTime        *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"` // 更新时间
	TbOrgID           string     `gorm:"column:tb_org_id;type:varchar(45);not null"`
}

// TableName get sql table name.获取数据库表名
func (m *TbProject) TableName() string {
	return "tb_project"
}
