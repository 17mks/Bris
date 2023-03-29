package chdb

import "time"

// TbUser 用户表
type TbUser struct {
	ID          string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'用户编码'"`                                                                                              // 用户编码
	Mobile      string     `gorm:"unique;column:mobile;type:varchar(45);default:null;comment:'手机号'"`                                                                                           // 手机号
	Email       string     `gorm:"unique;column:email;type:varchar(255);default:null;comment:'邮箱地址'"`                                                                                          // 邮箱地址
	EmployeeNum string     `gorm:"unique;column:employee_num;type:varchar(45);default:null;comment:'内部工号'"`                                                                                    // 内部工号
	Password    string     `gorm:"column:password;type:varchar(512);default:null;comment:'密码'"`                                                                                                // 密码
	Status      string     `gorm:"column:status;type:enum('0','1');default:null;default:0;comment:'账号状态(0：正常1：锁定)'"`                                                                           // 账号状态(0：正常1：锁定)
	Source      string     `gorm:"column:source;type:enum('P_REG','M_INSERT','M_IMPORT','NONE');default:null;default:NONE;comment:'账号数据来源('P_REG 患者端注册', 'M_INSERT 管理员添加','M_IMPORT 管理员导入')'"` // 账号数据来源('P_REG 患者端注册', 'M_INSERT 管理员添加','M_IMPORT 管理员导入')
	CreateTime  *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                                                                     // 创建时间
	UpdateTime  *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                                                                     // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *TbUser) TableName() string {
	return "tb_user"
}
