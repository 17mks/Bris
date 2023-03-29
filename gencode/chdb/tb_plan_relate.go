package chdb

import "time"

// TbPlanRelate 方案关联表
type TbPlanRelate struct {
	ID                string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'方案信息编码'"`                                                           // 方案信息编码
	Title             string     `gorm:"column:title;type:varchar(255);not null;comment:'标题(根据关联资源类型决定)'"`                                                          // 标题(根据关联资源类型决定)
	ResourceType      string     `gorm:"column:resource_type;type:enum('FORM','WI','BZ','GNZA');not null;comment:'关联资源类型(‘FORM 表单’,'WI 工作项','BZ 病种','GNZA 功能障碍')'"` // 关联资源类型(‘FORM 表单’,'WI 工作项','BZ 病种','GNZA 功能障碍')
	ResourceID        string     `gorm:"column:resource_id;type:varchar(45);not null;comment:'关联资源编码'"`                                                             // 关联资源编码
	FrequencyInterval int32      `gorm:"column:frequency_interval;type:int(11);default:null;comment:'频次间隔(小时)'"`                                                    // 频次间隔(小时)
	FrequencyOffset   int32      `gorm:"column:frequency_offset;type:int(11);default:null;comment:'频次偏移(小时)'"`                                                      // 频次偏移(小时)
	Times             int32      `gorm:"column:times;type:int(11);default:null;comment:'次数'"`                                                                       // 次数
	SortNum           uint32     `gorm:"column:sort_num;type:int(10) unsigned;default:null;comment:'排序号码'"`                                                         // 排序号码
	CreateTime        *time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                                        // 创建时间
	UpdateTime        *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                                        // 更新时间
	TbPlanID          string     `gorm:"column:tb_plan_id;type:varchar(45);not null"`
	DeleteAt          *time.Time `gorm:"column:delete_at;type:datetime;default:null;comment:'删除时间'"` // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m *TbPlanRelate) TableName() string {
	return "tb_plan_relate"
}
