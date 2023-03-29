package service

import (
	"context"
	"fmt"
	v1 "followup/api"
	"followup/gencode/chdb"
	"followup/utils"
	"log"

	"github.com/go-kratos/kratos/v2/transport"
)

func (s *ArticleService) ArticleCreate(ctx context.Context, req *v1.ArticleCreateRequest) (resp *v1.ArticleCreateResponse, err error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	toInsertArticle := chdb.TbArticle{
		ID:          GenNewUUID(),
		Type:        req.Type,
		Title:       req.Title,
		ThumbImgURL: req.ThumbImgUrl,
		Summary:     req.Summary,
		Author:      req.Author,
		Status:      req.Status,
		ContentURL:  req.ContentUrl,
		Contents:    req.Contents,
	}
	newArticle, err := s.Article.CreateArticle(ctx, &toInsertArticle)
	if err != nil {
		return nil, err
	}

	return &v1.ArticleCreateResponse{
		Id: newArticle.ID,
	}, nil
}

func (s *ArticleService) ArticleDelete(ctx context.Context, req *v1.ArticleDeleteRequest) (*v1.ArticleDeleteResponse, error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	deletedId, err := s.Article.DeleteArticle(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ArticleDeleteResponse{
		Id: deletedId,
	}, nil
}

func (s *ArticleService) ArticleUpdate(ctx context.Context, req *v1.ArticleUpdateRequest) (res *v1.ArticleUpdateResponse, errs error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	at, err := s.Article.UpdateArticle(ctx, &chdb.TbArticle{
		ID:          req.Id,
		Type:        req.Body.Type,
		Status:      req.Body.Status,
		Title:       req.Body.Title,
		Summary:     req.Body.Summary,
		ThumbImgURL: req.Body.ThumbImgUrl,
		Author:      req.Body.Author,
		Contents:    req.Body.Contents,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ArticleUpdateResponse{
		Id:          at.ID,
		Type:        at.Type,
		Title:       at.Title,
		ThumbImgUrl: at.ThumbImgURL,
		Summary:     at.Summary,
		Author:      at.Author,
		Status:      at.Status,
		Contents:    at.Contents,
		ContentUrl:  at.ContentURL,
	}, nil
}

func (s *ArticleService) ArticleDetail(ctx context.Context, apiReq *v1.ArticleDetailRequest) (resp *v1.ArticleDetailResponse, errs error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	at, err := s.Article.DetailArticle(ctx, apiReq.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ArticleDetailResponse{
		Id:          at.ID,
		Type:        at.Type,
		Title:       at.Title,
		ThumbImgUrl: at.ThumbImgURL,
		Summary:     at.Summary,
		Author:      at.Author,
		Status:      at.Status,
		Contents:    at.Contents,
		ContentUrl:  at.ContentURL,
	}, nil
}

func (s *ArticleService) ArticleFilter(ctx context.Context, apiReq *v1.ArticleFilterRequest) (resp *v1.ArticleFilterResponse, errs error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	articleList, count, err := s.Article.FilterArticle(ctx, apiReq)
	if err != nil {
		return nil, err
	}
	var articleListRes []*v1.Articles
	for _, at := range articleList {
		articleListRes = append(articleListRes,
			&v1.Articles{
				Id:          at.ID,
				Type:        at.Type,
				Title:       at.Title,
				ThumbImgUrl: at.ThumbImgURL,
				Summary:     at.Summary,
				Author:      at.Author,
				Status:      at.Status,
				Contents:    at.Contents,
				ContentUrl:  at.ContentURL,
				CreateTime:  *TimeToString(at.CreateTime),
				UpdateTime:  *TimeToString(at.UpdateTime),
			})
	}
	return &v1.ArticleFilterResponse{
		Total:     count,
		TotalPage: count,
		Page:      apiReq.Page,
		PerPage:   apiReq.PerPage,
		Results:   articleListRes,
	}, err
}
