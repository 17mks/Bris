package chdb

import "time"

// TbFormColumn 表单列
type TbFormColumn struct {
	ID             string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'表单字段编码'"`                        // 表单字段编码
	Name           string     `gorm:"uniqueIndex:name_UNIQUE;column:name;type:varchar(128);default:null;comment:'表单字段名称'"`    // 表单字段名称
	GroupName      string     `gorm:"column:group_name;type:varchar(128);default:null;comment:'表单字段分组名称'"`                    // 表单字段分组名称
	DataType       string     `gorm:"column:data_type;type:varchar(45);default:null;comment:'表单字段数据类型'"`                      // 表单字段数据类型
	ViewType       string     `gorm:"column:view_type;type:varchar(45);default:null;comment:'表单字段View类型'"`                    // 表单字段View类型
	AvailableValue string     `gorm:"column:available_value;type:varchar(512);default:null;comment:'表单字段可用值(列表)'"`            // 表单字段可用值(列表)
	DefaultValue   string     `gorm:"column:default_value;type:varchar(45);default:null;comment:'表单字段默认值'"`                   // 表单字段默认值
	Regexp         string     `gorm:"column:regexp;type:varchar(512);default:null;comment:'表单字段数据校验正则'"`                      // 表单字段数据校验正则
	SortIndex      int32      `gorm:"column:sort_index;type:int(11);default:null;comment:'字段排序'"`                             // 字段排序
	Required       bool       `gorm:"column:required;type:tinyint(1);default:null;comment:'是否必填项'"`                           // 是否必填项
	Description    string     `gorm:"column:description;type:varchar(512);default:null;comment:'表单字段描述'"`                     // 表单字段描述
	Visible        bool       `gorm:"column:visible;type:tinyint(1);default:null;default:1;comment:'是否可见'"`                   // 是否可见
	Editable       bool       `gorm:"column:editable;type:tinyint(1);default:null;default:1;comment:'是否可编辑'"`                 // 是否可编辑
	CreateTime     *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"` // 创建时间
	UpdateTime     *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"` // 更新时间
	TbFormID       string     `gorm:"uniqueIndex:name_UNIQUE;column:tb_form_id;type:varchar(45);not null"`
}

// TableName get sql table name.获取数据库表名
func (m *TbFormColumn) TableName() string {
	return "tb_form_column"
}
