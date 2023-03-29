package chdb

import "time"

// TbUserInfo 用户基础信息表
type TbUserInfo struct {
	ID                 string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'用户详细编码'"`                                                                   // 用户详细编码
	Name               string     `gorm:"column:name;type:varchar(128);default:null;comment:'姓名'"`                                                                           // 姓名
	Gender             string     `gorm:"column:gender;type:enum('MALE','FEMALE','OTHER','NONE');default:null;default:NONE;comment:'性别('MALE', 'FEMALE', 'OTHER', 'NONE')'"` // 性别('MALE', 'FEMALE', 'OTHER', 'NONE')
	Nation             string     `gorm:"column:nation;type:varchar(128);default:null;comment:'民族'"`                                                                         // 民族
	DateOfBirth        *time.Time `gorm:"column:date_of_birth;type:date;default:null;comment:'出生日期(yyyy-MM-DD)'"`                                                            // 出生日期(yyyy-MM-DD)
	IDNumber           string     `gorm:"column:id_number;type:varchar(45);default:null;comment:'身份证号码'"`                                                                    // 身份证号码
	Address            string     `gorm:"column:address;type:varchar(512);default:null;comment:'户籍地址'"`                                                                      // 户籍地址
	IssuingAuthority   string     `gorm:"column:issuing_authority;type:varchar(255);default:null;comment:'签发机关(派出所或发证机构)'"`                                                  // 签发机关(派出所或发证机构)
	EffectiveDate      *time.Time `gorm:"column:effective_date;type:datetime;default:null;comment:'证件生效日期'"`                                                                 // 证件生效日期
	ExpirationDate     *time.Time `gorm:"column:expiration_date;type:datetime;default:null;comment:'证件失效日期'"`                                                                // 证件失效日期
	ProfilePhoto       string     `gorm:"column:profile_photo;type:varchar(512);default:null;comment:'证件头像'"`                                                                // 证件头像
	Certified          bool       `gorm:"column:certified;type:tinyint(1);default:null;default:0;comment:'是否实名认证(身份证和姓名二元素认证)'"`                                             // 是否实名认证(身份证和姓名二元素认证)
	CertifiedTime      *time.Time `gorm:"column:certified_time;type:datetime;default:null;comment:'实名认证认证时间'"`                                                               // 实名认证认证时间
	Avatar             string     `gorm:"column:avatar;type:varchar(512);default:null;comment:'用户头像'"`                                                                       // 用户头像
	ResidentialAddr    string     `gorm:"column:residential_addr;type:varchar(512);default:null;comment:'居住地址'"`                                                             // 居住地址
	FormerName         string     `gorm:"column:former_name;type:varchar(128);default:null;comment:'曾用名'"`                                                                   // 曾用名
	Nickname           string     `gorm:"column:nickname;type:varchar(45);default:null;comment:'昵称'"`                                                                        // 昵称
	MobileVerified     bool       `gorm:"column:mobile_verified;type:tinyint(1);default:null;default:0;comment:'手机号是否验证'"`                                                   // 手机号是否验证
	MobileVerifiedTime *time.Time `gorm:"column:mobile_verified_time;type:datetime;default:null;comment:'手机号认证时间'"`                                                          // 手机号认证时间
	EmailVerified      bool       `gorm:"column:email_verified;type:tinyint(1);default:null;default:0;comment:'邮箱地址是否验证'"`                                                   // 邮箱地址是否验证
	EmailVerifiedTime  string     `gorm:"column:email_verified_time;type:varchar(45);default:null;comment:'邮箱认证时间'"`                                                         // 邮箱认证时间
	CreateTime         *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                                            // 创建时间
	UpdateTime         *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                                            // 更新时间
	TbUserID           string     `gorm:"column:tb_user_id;type:varchar(45);not null"`
}

// TableName get sql table name.获取数据库表名
func (m *TbUserInfo) TableName() string {
	return "tb_user_info"
}
