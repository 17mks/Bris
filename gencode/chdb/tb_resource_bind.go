package chdb

import "time"

// TbResourceBind 主体与资源绑定表
type TbResourceBind struct {
	ID              string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'文件资源绑定编码'"`                                                                                                                                    // 文件资源绑定编码
	BodyType        string     `gorm:"uniqueIndex:body_type_UNIQUE;column:body_type;type:varchar(45);not null;comment:'主体类型(USER,ORG,TEAM,GROUP,TASK,TARGET,PLAN,PROGRAM)'"`                                                                 // 主体类型(USER,ORG,TEAM,GROUP,TASK,TARGET,PLAN,PROGRAM)
	BodyID          string     `gorm:"uniqueIndex:body_type_UNIQUE;column:body_id;type:varchar(45);not null;comment:'主体编码'"`                                                                                                                 // 主体编码
	ResourceType    string     `gorm:"uniqueIndex:body_type_UNIQUE;column:resource_type;type:varchar(45);not null;comment:'资源类型('ROLE', 'ORG', 'CERTIFICATE', 'PATIENT', 'FILE', 'TEAM', 'GROUP', 'TASK', 'PROJECT', 'MEMBER', 'PROGRAM')'"` // 资源类型('ROLE', 'ORG', 'CERTIFICATE', 'PATIENT', 'FILE', 'TEAM', 'GROUP', 'TASK', 'PROJECT', 'MEMBER', 'PROGRAM')
	ResourceID      string     `gorm:"uniqueIndex:body_type_UNIQUE;column:resource_id;type:varchar(45);not null;comment:'资源编码(内部资源)'"`                                                                                                       // 资源编码(内部资源)
	OuterResourceID string     `gorm:"column:outer_resource_id;type:varchar(45);default:null;comment:'外部资源编码'"`                                                                                                                              // 外部资源编码
	Creator         string     `gorm:"column:creator;type:varchar(45);default:null;comment:'创建者'"`                                                                                                                                           // 创建者
	CreateTime      *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                                                                                                               // 创建时间
	UpdateTime      *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                                                                                                               // 更新时间
	DeletedAt       *time.Time `gorm:"uniqueIndex:body_type_UNIQUE;index:delete_at_INDEX;column:deleted_at;type:datetime;default:null;comment:'删除时间'"`                                                                                       // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m *TbResourceBind) TableName() string {
	return "tb_resource_bind"
}
