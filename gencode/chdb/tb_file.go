package chdb

import "time"

// TbFile 文件表
type TbFile struct {
	ID              string     `gorm:"primaryKey;column:id;type:varchar(45);not null"`
	Type            string     `gorm:"column:type;type:enum('DIR','IMAGE','VIDEO','ZIP');not null;comment:'文件类型(DIR,IMAGE,VIDEO)'"`                                                                                    // 文件类型(DIR,IMAGE,VIDEO)
	Name            string     `gorm:"column:name;type:varchar(255);not null;comment:'文件名'"`                                                                                                                           // 文件名
	Pid             string     `gorm:"column:pid;type:varchar(45);default:null;comment:'父级编码(如果为空则表示存放到根目录)'"`                                                                                                         // 父级编码(如果为空则表示存放到根目录)
	StoragePlatform string     `gorm:"column:storage_platform;type:varchar(45);default:null;comment:'资源存储平台(AliOSS)'"`                                                                                                 // 资源存储平台(AliOSS)
	FileURL         string     `gorm:"column:file_url;type:text;default:null;comment:'文件路径(如果为第三方存储则为全路径,否则为相对路径)'"`                                                                                                   // 文件路径(如果为第三方存储则为全路径,否则为相对路径)
	Size            int32      `gorm:"column:size;type:int(11);default:null;comment:'文件大小(byte)'"`                                                                                                                     // 文件大小(byte)
	ServiceModule   string     `gorm:"column:service_module;type:enum('ARTICLE','PROGRAM','TARGET','PLAN','CERTIFICATE','FORM');default:null;comment:'文件所属业务('ARTICLE', 'PROGRAM', 'TARGET', 'PLAN', 'CERTIFICATE')'"` // 文件所属业务('ARTICLE', 'PROGRAM', 'TARGET', 'PLAN', 'CERTIFICATE')
	Sha256          string     `gorm:"column:sha256;type:varchar(255);default:null;comment:'文件sha256值'"`                                                                                                               // 文件sha256值
	Suffix          string     `gorm:"column:suffix;type:varchar(45);default:null;comment:'文件后缀名'"`                                                                                                                    // 文件后缀名
	Uploader        string     `gorm:"column:uploader;type:varchar(45);default:null;comment:'上传者'"`                                                                                                                    // 上传者
	CreateTime      *time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                                                                                             // 创建时间
	UpdateTime      *time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                                                                                             // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *TbFile) TableName() string {
	return "tb_file"
}
