package service

import (
	"context"
	"fmt"
	v1 "followup/api"
	"followup/gencode/chdb"
	"followup/utils"
	"github.com/go-kratos/kratos/v2/transport"
	"log"
)

func ConvertTbDisFuncToGenModel(req *v1.DisFunctions) *v1.DisFunctions {
	return &v1.DisFunctions{
		CreateTime:  req.CreateTime,
		DeletedAt:   req.DeletedAt,
		Description: req.Description,
		Id:          req.Id,
		Name:        req.Name,
		UpdateTime:  req.UpdateTime,
	}
}

func (s *DisFunctionService) DisFunctionCreate(ctx context.Context, req *v1.DisFunctionCreateRequest) (*v1.DisFunctionCreateResponse, error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	at, err := s.DisFunction.CreateDisFunction(ctx, &chdb.TbDisFunc{
		ID:          GenNewUUID(),
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DisFunctionCreateResponse{
		Id: at.ID,
	}, nil
}

func (s *DisFunctionService) DisFunctionDelete(ctx context.Context, req *v1.DisFunctionDeleteRequest) (*v1.DisFunctionDeleteResponse, error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	deletedId, err := s.DisFunction.DeleteDisFunction(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DisFunctionDeleteResponse{
		Id: deletedId,
	}, nil
}

func (s *DisFunctionService) DisFunctionUpdate(ctx context.Context, req *v1.DisFunctionUpdateRequest) (res *v1.DisFunctionUpdateResponse, errs error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	at, err := s.DisFunction.UpdateDisFunction(ctx, &chdb.TbDisFunc{
		ID:          req.Id,
		Name:        req.Body.Name,
		Description: req.Body.Description,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DisFunctionUpdateResponse{
		Id:          at.ID,
		Name:        at.Name,
		Description: at.Description,
		CreateTime:  *TimeToString(at.CreateTime),
		UpdateTime:  *TimeToString(GetNowTimeAddr()),
		DeletedAt:   *TimeToString(GetNowTimeAddr()),
	}, nil
}

func (s *DisFunctionService) DisFunctionDetail(ctx context.Context, req *v1.DisFunctionDetailRequest) (*v1.DisFunctionDetailResponse, error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	at, err := s.DisFunction.DetailDisFunction(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DisFunctionDetailResponse{
		Id:          at.ID,
		Name:        at.Name,
		Description: at.Description,
		CreateTime:  *TimeToString(at.CreateTime),
		UpdateTime:  *TimeToString(at.UpdateTime),
		DeletedAt:   *TimeToString(GetNowTimeAddr()),
	}, nil
}

func (s *DisFunctionService) DisFunctionFilter(ctx context.Context, apiReq *v1.DisFunctionFilterRequest) (resp *v1.DisFunctionFilterResponse, err error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	DisFunctionList, count, err := s.DisFunction.FilterDisFunction(ctx, apiReq)
	if err != nil {
		return nil, err
	}
	var DisFunctionListRes []*v1.DisFunctions
	for _, at := range DisFunctionList {
		DisFunctionListRes = append(DisFunctionListRes,
			&v1.DisFunctions{
				Id:          at.ID,
				Name:        at.Name,
				Description: at.Description,
				Py:          *TimeToString(at.Py),
				CreateTime:  *TimeToString(at.CreateTime),
				UpdateTime:  *TimeToString(at.UpdateTime),
				DeletedAt:   *TimeToString(GetNowTimeAddr()),
			})
	}
	return &v1.DisFunctionFilterResponse{
		Page:    apiReq.Page,
		PerPage: apiReq.PerPage,
		Results: DisFunctionListRes,
		Total:   count,
	}, nil
}
