package chdb

import "time"

// TbDisease [...]
type TbDisease struct {
	ID          string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:''病种主键编码''"`                                      // '病种主键编码'
	Code        string     `gorm:"unique;column:code;type:varchar(45);default:null;comment:''疾病编码(e.g. A00-B99 C00.1)''"`                  // '疾病编码(e.g. A00-B99 C00.1)'
	Name        string     `gorm:"column:name;type:varchar(512);not null;comment:''疾病名称''"`                                                // '疾病名称'
	NameJp      string     `gorm:"column:name_jp;type:varchar(512);not null;comment:''疾病名称首字母拼音''"`                                        // '疾病名称首字母拼音'
	NameQp      string     `gorm:"column:name_qp;type:varchar(512);not null;comment:''疾病名称拼音全拼''"`                                         // '疾病名称拼音全拼'
	Version     string     `gorm:"column:version;type:varchar(45);not null;comment:''版本(e.g. 'ICD10')''"`                                  // '版本(e.g. 'ICD10')'
	Status      string     `gorm:"column:status;type:enum('ENABLED','DISABLED');default:null;comment:''状态('ENABLED 启用', 'DISABLED 禁用')''"` // '状态('ENABLED 启用', 'DISABLED 禁用')'
	Tag         string     `gorm:"column:tag;type:varchar(512);default:null;comment:''备用标记''"`                                             // '备用标记'
	Description string     `gorm:"column:description;type:varchar(512);default:null;comment:''疾病描述''"`                                     // '疾病描述'
	Pid         string     `gorm:"column:pid;type:varchar(45);default:null;comment:''父级编码''"`                                              // '父级编码'
	CreateTime  *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:''创建时间''"`               // '创建时间'
	UpdateTime  *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:''更新时间''"`               // '更新时间'
	DeleteAt    *time.Time `gorm:"column:delete_at;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:''删除时间''"`                 // '删除时间'
}

// TableName get sql table name.获取数据库表名
func (m *TbDisease) TableName() string {
	return "tb_disease"
}
