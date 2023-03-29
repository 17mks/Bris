package chdb

import "time"

// TbColumnThreshold 表单字段预警阈值表
type TbColumnThreshold struct {
	ID             string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'表单字段阈值编码'"`                               // 表单字段阈值编码
	Min            string     `gorm:"uniqueIndex:min_UNIQUE;column:min;type:varchar(45);default:null;comment:'表单字段阈值最小值'"`             // 表单字段阈值最小值
	Max            string     `gorm:"uniqueIndex:min_UNIQUE;column:max;type:varchar(45);default:null;comment:'表单字段阈值最大值'"`             // 表单字段阈值最大值
	Reverse        bool       `gorm:"uniqueIndex:min_UNIQUE;column:reverse;type:tinyint(1);default:null;comment:'是否反转'"`               // 是否反转
	Level          int32      `gorm:"column:level;type:int(11);default:null;comment:'表单字段阈值等级'"`                                       // 表单字段阈值等级
	WarningInfo    string     `gorm:"column:warning_info;type:varchar(128);default:null;comment:'表单字段预警提示信息'"`                         // 表单字段预警提示信息
	WarningRegex   string     `gorm:"column:warning_regex;type:varchar(512);default:null;comment:'预警信息组装规则(e.g. message{{key}},age)'"` // 预警信息组装规则(e.g. message{{key}},age)
	CreateTime     *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`          // 创建时间
	UpdateTime     *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`          // 更新时间
	TbFormColumnID string     `gorm:"uniqueIndex:min_UNIQUE;column:tb_form_column_id;type:varchar(45);not null"`
	TbProjectID    string     `gorm:"column:tb_project_id;type:varchar(45);not null"`
}

// TableName get sql table name.获取数据库表名
func (m *TbColumnThreshold) TableName() string {
	return "tb_column_threshold"
}
