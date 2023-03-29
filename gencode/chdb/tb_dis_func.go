package chdb

import "time"

// TbDisFunc [...]
type TbDisFunc struct {
	ID          string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:''功能障碍编码''"`                      // '功能障碍编码'
	Name        string     `gorm:"column:name;type:varchar(512);not null;comment:''功能障碍名称''"`                              // '功能障碍名称'
	Description string     `gorm:"column:description;type:varchar(512);default:null;comment:''描述''"`                       // '描述'
	Py          *time.Time `gorm:"column:py;type:datetime;default:null;comment:''功能障碍名称拼音''"`                              // '功能障碍名称拼音'
	CreateTime  *time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:''创建时间''"`   // '创建时间'
	UpdateTime  *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:''更新时间''"`   // '更新时间'
	DeleteAt    *time.Time `gorm:"column:delete_at;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:''删除时间''"` // '删除时间'
}

// TableName get sql table name.获取数据库表名
func (m *TbDisFunc) TableName() string {
	return "tb_dis_func"
}
