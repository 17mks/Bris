package chdb

import "time"

// TbPlans [...]
type TbPlans struct {
	ID               string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:''方案编码''"`                                                                        // '方案编码'
	Name             string     `gorm:"uniqueIndex:name_UNIQUE;column:name;type:varchar(128);not null;comment:''方案名称''"`                                                        // '方案名称'
	Type             string     `gorm:"column:type;type:enum('SA','FOLLOWUP');not null;comment:''方案类型(SA 症状评估, FOLLOWUP 随访)''"`                                                 // '方案类型(SA 症状评估, FOLLOWUP 随访)'
	Status           string     `gorm:"column:status;type:enum('DRAFT','ENABLED','DISABLED');not null;default:DRAFT;comment:''方案状态('DRAFT 草稿', 'ENABLED 启用', 'DISABLED 禁用')''"` // '方案状态('DRAFT 草稿', 'ENABLED 启用', 'DISABLED 禁用')'
	BelongType       string     `gorm:"column:belong_type;type:enum('ORG','PROJECT','TEAM','GROUP','NONE');not null;default:NONE;comment:''归属类型(组织、项目、团队、组等)''"`                // '归属类型(组织、项目、团队、组等)'
	BelongTo         string     `gorm:"uniqueIndex:name_UNIQUE;column:belong_to;type:varchar(45);default:null;comment:''规则资源编码(e.g. 如果归属类型时是项目则填项目编码)''"`                       // '规则资源编码(e.g. 如果归属类型时是项目则填项目编码)'
	ApplyDisease     string     `gorm:"column:apply_disease;type:text;default:null;comment:''适用病种''"`                                                                           // '适用病种'
	ApplyDysfunction string     `gorm:"column:apply_dysfunction;type:text;default:null;comment:''适用功能障碍(多个功能障碍用','分隔)''"`                                                       // '适用功能障碍(多个功能障碍用','分隔)'
	ApplyAges        string     `gorm:"column:apply_ages;type:text;default:null;comment:''适用年龄段(多个年龄段用','分隔)''"`                                                                // '适用年龄段(多个年龄段用','分隔)'
	Event            string     `gorm:"column:event;type:text;default:null;comment:''事件(随访开始触发事件)''"`                                                                           // '事件(随访开始触发事件)'
	CreatorID        string     `gorm:"column:creator_id;type:varchar(45);default:null;comment:''创建人编码''"`                                                                      // '创建人编码'
	CreatorName      string     `gorm:"column:creator_name;type:varchar(45);default:null;comment:''创建人名称''"`                                                                    // '创建人名称'
	AppID            string     `gorm:"column:app_id;type:varchar(45);default:null;comment:''应用编码''"`                                                                           // '应用编码'
	CreateTime       *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:''创建时间''"`                                               // '创建时间'
	UpdateTime       *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:''更新时间''"`                                               // '更新时间'
	DeleteAt         *time.Time `gorm:"column:delete_at;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:''删除时间''"`                                                 // '删除时间'
}

// TableName get sql table name.获取数据库表名
func (m *TbPlans) TableName() string {
	return "tb_plans"
}
