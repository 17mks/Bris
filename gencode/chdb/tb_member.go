package chdb

import "time"

// TbMember 成员表
type TbMember struct {
	ID         string     `gorm:"primaryKey;unique;column:id;type:varchar(45);not null;comment:'成员编码'"`                                                                             // 成员编码
	Name       string     `gorm:"column:name;type:varchar(128);not null;comment:'成员姓名'"`                                                                                            // 成员姓名
	Mobile     string     `gorm:"unique;column:mobile;type:varchar(45);default:null;comment:'成员手机号'"`                                                                               // 成员手机号
	Email      string     `gorm:"unique;column:email;type:varchar(255);default:null;comment:'成员邮箱'"`                                                                                // 成员邮箱
	EmployNum  string     `gorm:"unique;column:employ_num;type:varchar(45);default:null;comment:'成员工号'"`                                                                            // 成员工号
	Status     string     `gorm:"column:status;type:enum('WORKING','RESIGNED','HOLIDAY');default:null;default:WORKING;comment:'成员状态('WORKING 正常在职','RESIGNED 已离职','HOLIDAY 休假中')'"` // 成员状态('WORKING 正常在职','RESIGNED 已离职','HOLIDAY 休假中')
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                                                           // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                                                           // 更新时间
	TbUserID   string     `gorm:"uniqueIndex:tb_user_id_UNIQUE;column:tb_user_id;type:varchar(45);not null;comment:'成员用户编码'"`                                                       // 成员用户编码
	TbOrgID    string     `gorm:"uniqueIndex:tb_user_id_UNIQUE;column:tb_org_id;type:varchar(45);not null;comment:'组织编码'"`                                                          // 组织编码
}

// TableName get sql table name.获取数据库表名
func (m *TbMember) TableName() string {
	return "tb_member"
}
