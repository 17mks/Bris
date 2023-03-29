package chdb

import "time"

// TbEvent 事件表
type TbEvent struct {
	ID          string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'事件编码'"`                                                    // 事件编码
	Name        string     `gorm:"column:name;type:varchar(255);not null;comment:'事件名称'"`                                                            // 事件名称
	Type        string     `gorm:"column:type;type:varchar(45);not null;comment:'事件类型'"`                                                             // 事件类型
	Status      string     `gorm:"column:status;type:enum('ENABLED','DISABLED');not null;default:ENABLED;comment:'状态('ENABLED 启用', 'DISABLED 禁用')'"` // 状态('ENABLED 启用', 'DISABLED 禁用')
	SortNum     int32      `gorm:"column:sort_num;type:int(11);default:null;comment:'排序'"`                                                           // 排序
	Description string     `gorm:"column:description;type:varchar(255);default:null;comment:'事件描述'"`                                                 // 事件描述
	CreateTime  *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                           // 创建时间
	UpdateTime  *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                           // 更新时间
	DeletedAt   *time.Time `gorm:"index:deleted_at_INDEX;column:deleted_at;type:datetime;default:null;comment:'删除时间'"`                               // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m *TbEvent) TableName() string {
	return "tb_event"
}
