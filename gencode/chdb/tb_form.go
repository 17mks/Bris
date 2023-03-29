package chdb

import "time"

// TbForm [...]
type TbForm struct {
	ID          string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:''表单模板编码''"`                                                              // '表单模板编码'
	Name        string     `gorm:"column:name;type:varchar(256);not null;comment:''表单模板名称''"`                                                                      // '表单模板名称'
	Type        string     `gorm:"column:type;type:enum('SAF','IC','UP','SP','FOLLOW_UP');not null;comment:''表单类型('SAF 症状评估表', 'IC 知情同意书','UP 用户协议','SP 隐私协议')''"` // '表单类型('SAF 症状评估表', 'IC 知情同意书','UP 用户协议','SP 隐私协议')'
	Status      string     `gorm:"column:status;type:enum('DESIGNING','ENABLE','DISABLE');not null;default:DESIGNING;comment:''表单模板状态''"`                          // '表单模板状态'
	BranchLogic bool       `gorm:"column:branch_logic;type:tinyint(1);default:null;default:0;comment:''是否启用分支逻辑功能''"`                                              // '是否启用分支逻辑功能'
	BelongType  string     `gorm:"column:belong_type;type:enum('ORG','PROJECT','TEAM','PLAN','NONE');default:null;comment:''归属类型''"`                               // '归属类型'
	BelongTo    string     `gorm:"column:belong_to;type:varchar(45);default:null;comment:''归属资源编码''"`                                                              // '归属资源编码'
	Description string     `gorm:"column:description;type:varchar(512);default:null;comment:''表单模板描述''"`                                                           // '表单模板描述'
	AppID       string     `gorm:"column:app_id;type:varchar(512);default:null;comment:''应用编码''"`                                                                  // '应用编码'
	VersionName string     `gorm:"column:version_name;type:varchar(512);default:null;comment:''版本名称''"`                                                            // '版本名称'
	VersionCode int32      `gorm:"column:version_code;type:int(11);default:null;comment:''版本号''"`                                                                  // '版本号'
	CreateTime  *time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:''创建时间''"`                                           // '创建时间'
	UpdateTime  *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:''更新时间''"`                                           // '更新时间'
	DeleteAt    *time.Time `gorm:"column:delete_at;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:''删除时间''"`                                         // '删除时间'
}

// TableName get sql table name.获取数据库表名
func (m *TbForm) TableName() string {
	return "tb_form"
}
