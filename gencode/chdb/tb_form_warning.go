package chdb

import "time"

// TbFormWarning 表单预警
type TbFormWarning struct {
	ID                  string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'字段预警编码'"`                                                               // 字段预警编码
	PrincipalType       string     `gorm:"column:principal_type;type:enum('USER','MEMBER','NONE');default:null;default:MEMBER;comment:'负责人类型('USER', 'MEMBER', 'NONE')'"` // 负责人类型('USER', 'MEMBER', 'NONE')
	PrincipalID         string     `gorm:"column:principal_id;type:varchar(45);default:null;comment:'负责人成员编码'"`                                                           // 负责人成员编码
	PrincipalName       string     `gorm:"column:principal_name;type:varchar(128);default:null;comment:'负责人姓名'"`                                                          // 负责人姓名
	AssignedType        string     `gorm:"column:assigned_type;type:enum('USER','MEMBER','PATIENT','NONE');default:null;default:NONE;comment:'指派类型'"`                     // 指派类型
	AssignedTo          string     `gorm:"column:assigned_to;type:varchar(45);default:null;comment:'指派给谁'"`                                                               // 指派给谁
	AssignedToName      string     `gorm:"column:assigned_to_name;type:varchar(128);default:null;comment:'被指派人姓名'"`                                                       // 被指派人姓名
	FieldValue          float64    `gorm:"column:field_value;type:double;default:null;comment:'字段值'"`                                                                     // 字段值
	WarningInfo         string     `gorm:"column:warning_info;type:varchar(512);default:null;comment:'动态组装提示信息'"`                                                         // 动态组装提示信息
	ProjectID           string     `gorm:"column:project_id;type:varchar(45);default:null;comment:'项目编码'"`                                                                // 项目编码
	TbFormRowID         string     `gorm:"column:tb_form_row_id;type:varchar(45);not null;comment:'表单行编码'"`                                                               // 表单行编码
	TbColumnThresholdID string     `gorm:"column:tb_column_threshold_id;type:varchar(45);not null;comment:'字段阈值编码'"`                                                      // 字段阈值编码
	TbFormColumnID      string     `gorm:"column:tb_form_column_id;type:varchar(45);not null;comment:'字段编码(是哪个字段的预警)'"`                                                   // 字段编码(是哪个字段的预警)
	CreateTime          *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                                        // 创建时间
	UpdateTime          *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                                        // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *TbFormWarning) TableName() string {
	return "tb_form_warning"
}
