package chdb

import "time"

// TbThirdAccount 第三方平台账号
type TbThirdAccount struct {
	ID          string     `gorm:"primaryKey;column:id;type:varchar(45);not null"`
	Platform    string     `gorm:"uniqueIndex:platform_UNIQUE;column:platform;type:enum('WECHAT','BRIS');not null;comment:'第三方平台(WC：微信 BR：BRis)'"`                                                                              // 第三方平台(WC：微信 BR：BRis)
	AccountType string     `gorm:"uniqueIndex:platform_UNIQUE;column:account_type;type:enum('MOBILE','EMAIL','ID_NUMBER','EMPLOYEE_NUM','OPEN_ID');not null;comment:'第三方账号类型('MOBILE', 'EMAIL', 'ID_NUMBER', 'EMPLOYEE_NUM')'"` // 第三方账号类型('MOBILE', 'EMAIL', 'ID_NUMBER', 'EMPLOYEE_NUM')
	Account     string     `gorm:"uniqueIndex:platform_UNIQUE;column:account;type:varchar(255);not null;comment:'第三方平台账号'"`                                                                                                     // 第三方平台账号
	OuterUserID string     `gorm:"column:outer_user_id;type:varchar(45);not null;comment:'外部系统用户编码'"`                                                                                                                           // 外部系统用户编码
	CreateTime  *time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                                                                                                          // 创建时间
	UpdateTime  *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                                                                                                          // 更新时间
	TbUserID    string     `gorm:"column:tb_user_id;type:varchar(45);not null"`
}

// TableName get sql table name.获取数据库表名
func (m *TbThirdAccount) TableName() string {
	return "tb_third_account"
}
