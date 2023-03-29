package chdb

import "time"

// TbPatient 患者表(HIS导入或新建)
type TbPatient struct {
	ID              string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'患者编码'"`                                                      // 患者编码
	Name            string     `gorm:"column:name;type:varchar(128);not null;comment:'姓名(身份证)'"`                                                           // 姓名(身份证)
	Gender          string     `gorm:"column:gender;type:enum('MALE','FEMALE','NOT_SET');default:null;default:NOT_SET;comment:'性别(MALE, FEMALE,NOT_SET)'"` // 性别(MALE, FEMALE,NOT_SET)
	Nation          string     `gorm:"column:nation;type:varchar(45);default:null;comment:'民族'"`                                                           // 民族
	IDNumber        string     `gorm:"unique;column:id_number;type:varchar(45);default:null;comment:'身份证号码'"`                                              // 身份证号码
	Mobile          string     `gorm:"unique;column:mobile;type:varchar(45);default:null;comment:'手机号'"`                                                   // 手机号
	Email           string     `gorm:"unique;column:email;type:varchar(255);default:null;comment:'邮箱'"`                                                    // 邮箱
	Certified       bool       `gorm:"column:certified;type:tinyint(1);default:null;default:0;comment:'是否实名认证(身份证、姓名、性别)'"`                                // 是否实名认证(身份证、姓名、性别)
	BirthdayDate    *time.Time `gorm:"column:birthday_date;type:date;default:null;comment:'出生日期'"`                                                         // 出生日期
	OuterPlatform   string     `gorm:"uniqueIndex:outer_platform_UNIQUE;column:outer_platform;type:varchar(45);default:null;comment:'外部系统(BRis3.0,HIS等)'"` // 外部系统(BRis3.0,HIS等)
	OuterPatientOrg string     `gorm:"uniqueIndex:outer_platform_UNIQUE;column:outer_patient_org;type:varchar(255);default:null;comment:'外部患者管理机构名称'"`     // 外部患者管理机构名称
	OuterPatientID  string     `gorm:"uniqueIndex:outer_platform_UNIQUE;column:outer_patient_id;type:varchar(45);default:null;comment:'外部患者编码'"`           // 外部患者编码
	NamePy          string     `gorm:"column:name_py;type:varchar(512);default:null;comment:'姓名拼音'"`                                                       // 姓名拼音
	InHospital      bool       `gorm:"column:in_hospital;type:tinyint(1);default:null;comment:'患者是否在院'"`                                                   // 患者是否在院
	Outpno          string     `gorm:"column:outpno;type:varchar(45);default:null;comment:'门诊号'"`                                                          // 门诊号
	Inpno           string     `gorm:"column:inpno;type:varchar(45);default:null;comment:'住院号'"`                                                           // 住院号
	VisitCardNo     string     `gorm:"column:visit_card_no;type:varchar(125);default:null;comment:'就诊卡号'"`                                                 // 就诊卡号
	PatVisitType    string     `gorm:"column:pat_visit_type;type:varchar(45);default:null;comment:'病人来源'"`                                                 // 病人来源
	DeptID          string     `gorm:"column:dept_id;type:varchar(45);default:null;comment:'入院科室id'"`                                                      // 入院科室id
	DeptName        string     `gorm:"column:dept_name;type:varchar(45);default:null;comment:'入院科室'"`                                                      // 入院科室
	InpWardID       string     `gorm:"column:inp_ward_id;type:varchar(45);default:null;comment:'入院病区id'"`                                                  // 入院病区id
	InpWardName     string     `gorm:"column:inp_ward_name;type:varchar(45);default:null;comment:'入院病区'"`                                                  // 入院病区
	InpBedNo        string     `gorm:"column:inp_bed_no;type:varchar(45);default:null;comment:'入院病床'"`                                                     // 入院病床
	InpPnursID      string     `gorm:"column:inp_pnurs_id;type:varchar(45);default:null;comment:'责任护士id'"`                                                 // 责任护士id
	InpPnurs        string     `gorm:"column:inp_pnurs;type:varchar(45);default:null;comment:'责任护士'"`                                                      // 责任护士
	PatAtdpscn      string     `gorm:"column:pat_atdpscn;type:varchar(45);default:null;comment:'住院医师id'"`                                                  // 住院医师id
	PatAtdpscnName  string     `gorm:"column:pat_atdpscn_name;type:varchar(45);default:null;comment:'住院医师'"`                                               // 住院医师
	AdtaTime        *time.Time `gorm:"column:adta_time;type:datetime;default:null;comment:'入院日期'"`                                                         // 入院日期
	AdtdTime        *time.Time `gorm:"column:adtd_time;type:datetime;default:null;comment:'出院日期'"`                                                         // 出院日期
	PatInpTimes     string     `gorm:"column:pat_inp_times;type:varchar(45);default:null;comment:'住院次数'"`                                                  // 住院次数
	IndeptTime      *time.Time `gorm:"column:indept_time;type:datetime;default:null;comment:'入科时间'"`                                                       // 入科时间
	SurgeryTime     *time.Time `gorm:"column:surgery_time;type:datetime;default:null;comment:'手术时间'"`                                                      // 手术时间
	CreateTime      *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                             // 创建时间
	UpdateTime      *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                             // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *TbPatient) TableName() string {
	return "tb_patient"
}
