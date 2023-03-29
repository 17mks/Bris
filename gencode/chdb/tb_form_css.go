package chdb

import "time"

// TbFormCSS 表单样式表
type TbFormCSS struct {
	ID          string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'表单样式编码'"`                // 表单样式编码
	Name        string     `gorm:"column:name;type:varchar(255);default:null;comment:'表单样式名称'"`                    // 表单样式名称
	Status      string     `gorm:"column:status;type:varchar(45);default:null;comment:'表单样式状态'"`                   // 表单样式状态
	CSSCode     string     `gorm:"column:css_code;type:longtext;default:null;comment:'表单样式代码'"`                    // 表单样式代码
	CSSURL      string     `gorm:"column:css_url;type:varchar(512);default:null;comment:'表单模板通过/form/model访问的地址'"` // 表单模板通过/form/model访问的地址
	OssURL      string     `gorm:"column:oss_url;type:varchar(512);default:null;comment:'表单模板在对象存储服务器的地址'"`        // 表单模板在对象存储服务器的地址
	Description string     `gorm:"column:description;type:varchar(255);default:null;comment:'表单样式描述'"`             // 表单样式描述
	VersionName string     `gorm:"column:version_name;type:varchar(45);default:null"`
	VersionCode int32      `gorm:"column:version_code;type:int(11);default:null"`
	CreateTime  *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"` // 创建时间
	UpdateTime  *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"` // 更新时间
	TbFormID    string     `gorm:"column:tb_form_id;type:varchar(45);not null"`
}

// TableName get sql table name.获取数据库表名
func (m *TbFormCSS) TableName() string {
	return "tb_form_css"
}
