package chdb

import "time"

// TbArticle [...]
type TbArticle struct {
	ID          string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:''文章编码''"`                                                   // '文章编码'
	Type        string     `gorm:"column:type;type:enum('MISSION','KNOWLEDGE');not null;comment:''文章类型(MISSION：宣教 KNOWLEDGE：专家知识)''"`                 // '文章类型(MISSION：宣教 KNOWLEDGE：专家知识)'
	Title       string     `gorm:"index:title_UNIQUE;column:title;type:varchar(255);not null;comment:''文章标题''"`                                       // '文章标题'
	ThumbImgURL string     `gorm:"column:thumb_img_url;type:text;default:null;comment:''缩略图地址''"`                                                     // '缩略图地址'
	Summary     string     `gorm:"column:summary;type:text;default:null;comment:''文章概要''"`                                                            // '文章概要'
	Author      string     `gorm:"column:author;type:varchar(45);not null;comment:''文章作者''"`                                                          // '文章作者'
	Status      string     `gorm:"column:status;type:enum('DRAFT','ONLINE','OFFLINE');default:null;comment:''文章状态(DRAFT：草稿 ONLINE：上线  OFFLINE：下线)''"` // '文章状态(DRAFT：草稿 ONLINE：上线  OFFLINE：下线)'
	ContentURL  string     `gorm:"column:content_url;type:text;default:null;comment:''文章内容地址''"`                                                      // '文章内容地址'
	Contents    string     `gorm:"column:contents;type:longtext;default:null;comment:''文章内容''"`                                                       // '文章内容'
	CreateTime  *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:''创建时间''"`                          // '创建时间'
	UpdateTime  *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:''更新时间''"`                          // '更新时间'
}

// TableName get sql table name.获取数据库表名
func (m *TbArticle) TableName() string {
	return "tb_article"
}
