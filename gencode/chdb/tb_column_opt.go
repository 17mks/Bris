package chdb

import "time"

// TbColumnOpt 表单字段可选值选项
type TbColumnOpt struct {
	ID             string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'表单字段选项编码'"`                      // 表单字段选项编码
	SortIndex      int32      `gorm:"column:sort_index;type:int(11);default:null;comment:'表单字段选项排序索引'"`                       // 表单字段选项排序索引
	Value          string     `gorm:"column:value;type:varchar(512);default:null;comment:'表单字段选项值'"`                          // 表单字段选项值
	NextFiledID    string     `gorm:"column:next_filed_id;type:varchar(45);default:null;comment:'下一个表单字段编码(如果为空则按表单字段顺序展示)'"` // 下一个表单字段编码(如果为空则按表单字段顺序展示)
	Score          int32      `gorm:"column:score;type:int(11);default:null;comment:'得分'"`                                    // 得分
	CreateTime     *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"` // 创建时间
	UpdateTime     *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"` // 更新时间
	TbFormColumnID string     `gorm:"column:tb_form_column_id;type:varchar(45);not null"`
}

// TableName get sql table name.获取数据库表名
func (m *TbColumnOpt) TableName() string {
	return "tb_column_opt"
}
