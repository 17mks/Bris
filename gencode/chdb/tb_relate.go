package chdb

import "time"

// TbRelate 工作项关联表
type TbRelate struct {
	ID           string    `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'工作项关联编码'"`                                                                                                                                                                                                                                             // 工作项关联编码
	Title        string    `gorm:"column:title;type:varchar(255);not null;comment:'标题(根据关联资源类型决定)'"`                                                                                                                                                                                                                                             // 标题(根据关联资源类型决定)
	Status       string    `gorm:"column:status;type:enum('NEW','ACTIVE','CLOSED');not null;default:NEW;comment:'关联资源状态('NEW 待处理', 'ACTIVE 正在处理 'CLOSED 已关闭')'"`                                                                                                                                                                                 // 关联资源状态('NEW 待处理', 'ACTIVE 正在处理 'CLOSED 已关闭')
	ResourceType string    `gorm:"column:resource_type;type:enum('PLAN','FORM','WORK_ITEM','ARTICLE','FU_WC','FU_M','RM_FZ','RM_YH','RM_HY','RM_SS','NONE');not null;comment:'工作项关联资源类型ENUM('PLAN 方案', 'FORM 表单', 'WORK_ITEM 工作项', 'ARTICLE 宣教文章', 'FU_WC 微信随访', 'FU_M 电话随访', 'RM_FZ 复诊提醒', 'RM_YH 用药提醒', 'RM_HY 换药提醒', 'RM_SS 手术提醒', 'NONE')'"` // 工作项关联资源类型ENUM('PLAN 方案', 'FORM 表单', 'WORK_ITEM 工作项', 'ARTICLE 宣教文章', 'FU_WC 微信随访', 'FU_M 电话随访', 'RM_FZ 复诊提醒', 'RM_YH 用药提醒', 'RM_HY 换药提醒', 'RM_SS 手术提醒', 'NONE')
	ResourceID   string    `gorm:"column:resource_id;type:varchar(45);not null;comment:'关联资源编码'"`                                                                                                                                                                                                                                                // 关联资源编码
	Conclusion   string    `gorm:"column:conclusion;type:text;default:null;comment:'结论'"`                                                                                                                                                                                                                                                        // 结论
	Suggestion   string    `gorm:"column:suggestion;type:text;default:null;comment:'建议'"`                                                                                                                                                                                                                                                        // 建议
	Comments     string    `gorm:"column:comments;type:varchar(255);default:null;comment:'描述'"`                                                                                                                                                                                                                                                  // 描述
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                                                                                                                                                                                                                           // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                                                                                                                                                                                                                           // 更新时间
	TbWorkItemID string    `gorm:"column:tb_work_item_id;type:varchar(45);not null;comment:'工作项编码'"`                                                                                                                                                                                                                                             // 工作项编码
}

// TableName get sql table name.获取数据库表名
func (m *TbRelate) TableName() string {
	return "tb_relate"
}
