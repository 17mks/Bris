package chdb

import "time"

// TbFormRow 表单行数据
type TbFormRow struct {
	ID            string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'表单编码'"`                                              // 表单编码
	Title         string     `gorm:"column:title;type:varchar(255);default:null;comment:'表单名称'"`                                                 // 表单名称
	Status        string     `gorm:"column:status;type:enum('OPEN','LOCKED');not null;default:OPEN;comment:'表单状态(OPEN：开放,允许修改 LOCKED：锁定,禁止修改)'"` // 表单状态(OPEN：开放,允许修改 LOCKED：锁定,禁止修改)
	Data          string     `gorm:"column:data;type:text;default:null;comment:'提交的表单数据(JSON格式)'"`                                               // 提交的表单数据(JSON格式)
	Remark        string     `gorm:"column:remark;type:varchar(512);default:null;comment:'表单备注'"`                                                // 表单备注
	Submitter     string     `gorm:"column:submitter;type:varchar(45);not null;comment:'提交者用户编码'"`                                               // 提交者用户编码
	TbFormID      string     `gorm:"column:tb_form_id;type:varchar(45);not null;comment:'表单编码'"`                                                 // 表单编码
	TbRelateID    string     `gorm:"column:tb_relate_id;type:varchar(45);not null;comment:'关联编码'"`                                               // 关联编码
	SignaturePath string     `gorm:"column:signature_path;type:varchar(255);default:null;comment:'填表者签名存储路径'"`                                   // 填表者签名存储路径
	SignerID      string     `gorm:"column:signer_id;type:varchar(45);default:null;comment:'签名人编码'"`                                             // 签名人编码
	SignerName    string     `gorm:"column:signer_name;type:varchar(45);default:null;comment:'签名人姓名'"`                                           // 签名人姓名
	SignOnBehalf  bool       `gorm:"column:sign_on_behalf;type:tinyint(1);default:null;comment:'是否是代签署'"`                                        // 是否是代签署
	CreateTime    *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                     // 创建时间
	UpdateTime    *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                     // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *TbFormRow) TableName() string {
	return "tb_form_row"
}
