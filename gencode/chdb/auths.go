package chdb

import "time"

// Auths [...]
type Auths struct {
	ID        uint64     `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint(20) unsigned;not null"`
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime(3);default:null"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime(3);default:null"`
	DeletedAt *time.Time `gorm:"index:idx_auths_deleted_at;column:deleted_at;type:datetime(3);default:null"`
	UserID    string     `gorm:"column:user_id;type:varchar(500);default:null"`
	UserName  string     `gorm:"column:user_name;type:varchar(500);default:null"`
	AppID     string     `gorm:"column:app_id;type:varchar(500);default:null"`
	Token     string     `gorm:"column:token;type:varchar(500);default:null"`
}

// TableName get sql table name.获取数据库表名
func (m *Auths) TableName() string {
	return "auths"
}
