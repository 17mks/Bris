package chdb

import "time"

// TbComboResource 套餐资源表
type TbComboResource struct {
	ID         string     `gorm:"primaryKey;column:id;type:varchar(45);not null"`
	Type       string     `gorm:"column:type;type:varchar(45);default:null;comment:'套餐资源类型'"`                             // 套餐资源类型
	ResourceID string     `gorm:"column:resource_id;type:varchar(45);default:null;comment:'套餐资源编码'"`                      // 套餐资源编码
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"` // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"` // 更新时间
	TbComboID  string     `gorm:"column:tb_combo_id;type:varchar(45);not null"`
}

// TableName get sql table name.获取数据库表名
func (m *TbComboResource) TableName() string {
	return "tb_combo_resource"
}
