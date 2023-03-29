package data

import (
	"context"
	v1 "followup/api"
	"followup/gencode/chdb"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
	"time"

	"followup/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// TbArticle 文章
type TbArticle struct {
	ID          string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'文章编码'"`                                                   // 文章编码
	Type        string     `gorm:"column:type;type:enum('MISSION','KNOWLEDGE');not null;comment:'文章类型(MISSION：宣教 KNOWLEDGE：专家知识)'"`                 // 文章类型(MISSION：宣教 KNOWLEDGE：专家知识)
	Title       string     `gorm:"column:title;type:varchar(255);not null;comment:'文章标题'"`                                                          // 文章标题
	ThumbImgURL string     `gorm:"column:thumb_img_url;type:text;default:null;comment:'缩略图地址'"`                                                     // 缩略图地址
	Summary     string     `gorm:"column:summary;type:text;default:null;comment:'文章概要'"`                                                            // 文章概要
	Author      string     `gorm:"uniqueIndex:title_UNIQUE;column:author;type:varchar(45);not null;comment:'文章作者'"`                                 // 文章作者
	Status      string     `gorm:"column:status;type:enum('DRAFT','ONLINE','OFFLINE');default:null;comment:'文章状态(DRAFT：草稿 ONLINE：上线  OFFLINE：下线)'"` // 文章状态(DRAFT：草稿 ONLINE：上线  OFFLINE：下线)
	ContentURL  string     `gorm:"column:content_url;type:text;default:null;comment:'文章内容地址'"`                                                      // 文章内容地址
	Contents    string     `gorm:"column:contents;type:longtext;default:null;comment:'文章内容'"`                                                       // 文章内容
	CreateTime  *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                          // 创建时间
	UpdateTime  *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                          // 更新时间
}

type ArticleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &ArticleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ArticleRepo) Save(ctx context.Context, at *chdb.TbArticle) (*chdb.TbArticle, error) {
	if err := r.data.gormDB.Create(at).Error; err != nil {
		return nil, err
	}
	return at, nil
}

func (r *ArticleRepo) DeleteByID(ctx context.Context, id string) (string, error) {
	tbArticle := chdb.TbArticle{ID: id}
	if err := r.data.gormDB.Delete(&tbArticle).Error; err != nil {
		return "", err
	}
	return tbArticle.ID, nil
}

func (r *ArticleRepo) Update(ctx context.Context, req *chdb.TbArticle) (res *chdb.TbArticle, err error) {
	Article := new(chdb.TbArticle)
	err = r.data.gormDB.Where("id = ?", req.ID).First(Article).Error
	if err != nil {
		return nil, err
	}
	err = r.data.gormDB.Model(&Article).Updates(&chdb.TbArticle{
		ID:          req.ID,
		Type:        req.Type,
		Status:      req.Status,
		Title:       req.Title,
		ThumbImgURL: req.ThumbImgURL,
		Summary:     req.Summary,
		Author:      req.Author,
		Contents:    req.Contents,
	}).Error
	return &chdb.TbArticle{
		ID:          Article.ID,
		Type:        Article.Type,
		Status:      Article.Status,
		Title:       Article.Title,
		ThumbImgURL: Article.ThumbImgURL,
		Summary:     Article.Summary,
		Author:      Article.Author,
		Contents:    Article.Contents,
	}, nil
}

func (r *ArticleRepo) GetByID(ctx context.Context, id string) (rv *chdb.TbArticle, err error) {
	Article := new(chdb.TbArticle)
	result := r.data.gormDB.Where("id = ?", id).First(Article)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("Article", "not found by id")
	}
	if result.Error != nil {
		return nil, err
	}
	return &chdb.TbArticle{
		ID:          Article.ID,
		Type:        Article.Type,
		Status:      Article.Status,
		Title:       Article.Title,
		ThumbImgURL: Article.ThumbImgURL,
		Summary:     Article.Summary,
		Author:      Article.Author,
		Contents:    Article.Contents,
	}, nil
}

func (r *ArticleRepo) Filter(ctx context.Context, apiReq *v1.ArticleFilterRequest) ([]chdb.TbArticle, int64, error) {

	articlelist := make([]chdb.TbArticle, 0)
	var count int64 = 0

	filter := apiReq.Filter
	tx := r.data.gormDB.Model(&chdb.TbArticle{})
	if filter != nil {

		if "" != filter.Key {
			keyLike := AddLikeCharToStr(filter.Key)
			tx.Where("title like  ? or summary like  ? or contents like  ?", keyLike, keyLike, keyLike)
		}

		if len(filter.Id) > 0 {
			tx.Where("id in ?", filter.Id)
		}
		if len(filter.Status) > 0 {
			tx.Where("status in ?", filter.Status)
		}
		if len(filter.Type) > 0 {
			tx.Where("type in ?", filter.Type)
		}
	}

	tx.Order("update_time DESC")
	tx.Count(&count)

	result := tx.Find(&articlelist)
	return articlelist, count, result.Error
}
