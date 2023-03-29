package biz

import (
	"context"
	v1 "followup/api"
	"followup/gencode/chdb"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// a Article model.
type TbArticle struct {
	ID          string
	Type        string
	Title       string
	ThumbImgUrl string
	Summary     string
	Author      string
	Status      string
	ContentUrl  string
	Contents    string
	CreateTime  *time.Time
	UpdateTime  *time.Time
}

// 定义 Article 的操作接口
type ArticleRepo interface {
	Save(ctx context.Context, article *chdb.TbArticle) (*chdb.TbArticle, error)
	DeleteByID(context.Context, string) (string, error)
	Update(context.Context, *chdb.TbArticle) (*chdb.TbArticle, error)
	GetByID(context.Context, string) (*chdb.TbArticle, error)
	Filter(ctx context.Context, req *v1.ArticleFilterRequest) ([]chdb.TbArticle, int64, error)
}

type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{
		repo: repo,
		log:  log.NewHelper(logger)}
}

// TimeToString 将时间转换为sql查询的字符串格式(防止框架出现增加8小时的问题)
func TimeToString(time *time.Time) string {
	if nil == time || time.IsZero() {
		return ""
	}
	return time.Format("2006-01-02 15:04:05")
}

func (uc *ArticleUsecase) CreateArticle(ctx context.Context, req *chdb.TbArticle) (*chdb.TbArticle, error) {
	uc.log.WithContext(ctx).Infof("CreateArticle")
	return uc.repo.Save(ctx, req)
}
func (uc *ArticleUsecase) DeleteArticle(ctx context.Context, id string) (string, error) {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.DeleteByID(ctx, id)
}

func (uc *ArticleUsecase) UpdateArticle(ctx context.Context, Article *chdb.TbArticle) (*chdb.TbArticle, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", Article.ID)
	return uc.repo.Update(ctx, Article)
}

func (uc *ArticleUsecase) DetailArticle(ctx context.Context, id string) (*chdb.TbArticle, error) {
	uc.log.WithContext(ctx).Infof("Get: %d", id)
	return uc.repo.GetByID(ctx, id)
}

func (uc *ArticleUsecase) FilterArticle(ctx context.Context, req *v1.ArticleFilterRequest) ([]chdb.TbArticle, int64, error) {
	uc.log.WithContext(ctx).Infof("filter")
	return uc.repo.Filter(ctx, req)
}
