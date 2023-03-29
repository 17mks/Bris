package chdb

import "time"

// TbPatientInfo 患者信息表
type TbPatientInfo struct {
	ID               string     `gorm:"column:id;type:varchar(45);not null"`
	Province         string     `gorm:"column:province;type:varchar(255);default:null;comment:'省（自治区、直辖市）'"`    // 省（自治区、直辖市）
	City             string     `gorm:"column:city;type:varchar(255);default:null;comment:'市（地区）'"`             // 市（地区）
	County           string     `gorm:"column:county;type:varchar(255);default:null;comment:'县（区）'"`            // 县（区）
	Township         string     `gorm:"column:township;type:varchar(255);default:null;comment:'乡（镇、街道、路等）'"`    // 乡（镇、街道、路等）
	Village          string     `gorm:"column:village;type:varchar(255);default:null;comment:'村（街、路等）'"`        // 村（街、路等）
	HouseNumber      string     `gorm:"column:house_number;type:varchar(255);default:null;comment:'门牌号码'"`      // 门牌号码
	ResidentialAddr  string     `gorm:"column:residential_addr;type:text;default:null;comment:'现居住地址'"`         // 现居住地址
	PostalCode       string     `gorm:"column:postal_code;type:varchar(45);default:null;comment:'邮政编码'"`        // 邮政编码
	ContactNumber    string     `gorm:"column:contact_number;type:varchar(45);default:null;comment:'联系电话（座机）'"` // 联系电话（座机）
	MaritalStatus    string     `gorm:"column:marital_status;type:varchar(45);default:null;comment:'婚姻状况'"`     // 婚姻状况
	WorkUnitName     string     `gorm:"column:work_unit_name;type:varchar(255);default:null;comment:'工作单位名称'"`  // 工作单位名称
	EducationCode    string     `gorm:"column:education_code;type:varchar(45);default:null;comment:'教育编码'"`     // 教育编码
	OccStatus        string     `gorm:"column:occ_status;type:varchar(45);default:null"`
	OccCode          string     `gorm:"column:occ_code;type:varchar(45);default:null"`
	HealthRecordCode string     `gorm:"column:health_record_code;type:varchar(45);default:null;comment:'健康档案编码'"`         // 健康档案编码
	MobileCertified  bool       `gorm:"column:mobile_certified;type:tinyint(1);default:null;default:0;comment:'手机号是否验证'"` // 手机号是否验证
	EmailCertified   bool       `gorm:"column:email_certified;type:tinyint(1);default:null;default:0;comment:'手机号是否验证'"`  // 手机号是否验证
	TbPatientID      string     `gorm:"column:tb_patient_id;type:varchar(45);not null"`
	CreateTime       *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"` // 创建时间
	UpdateTime       *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *TbPatientInfo) TableName() string {
	return "tb_patient_info"
}
