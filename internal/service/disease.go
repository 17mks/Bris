package service

import (
	"context"
	"fmt"
	v1 "followup/api"
	"followup/gencode/chdb"
	"followup/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	Log "log"

	"time"
)

func ConvertDiseaseToTbDisease(req *chdb.TbDisease) *v1.Diseases {
	return &v1.Diseases{
		Code:        req.Code,
		CreateTime:  *TimeToString(req.CreateTime),
		Description: req.Description,
		Id:          req.ID,
		Name:        req.Name,
		NameJp:      req.NameJp,
		NameQp:      req.NameQp,
		Pid:         req.Pid,
		Status:      req.Status,
		Tag:         req.Tag,
		UpdateTime:  *TimeToString(req.UpdateTime),
		Version:     req.Version,
	}
}

func ConvertTbDiseaseToGenModel(req *chdb.TbDisease) *v1.Diseases {
	return &v1.Diseases{
		Code:        req.Code,
		CreateTime:  *TimeToString(req.CreateTime),
		Description: req.Description,
		Id:          req.ID,
		Name:        req.Name,
		NameJp:      req.NameJp,
		NameQp:      req.NameQp,
		Pid:         req.Pid,
		Status:      req.Status,
		Tag:         req.Tag,
		UpdateTime:  *TimeToString(req.UpdateTime),
		Version:     req.Version,
	}
}

func StringToTime(timeStr string) *time.Time {
	var layout = "2006-01-02 15:04" //转换的时间字符串带秒则 为 2006-01-02 15:04:05
	timeVal, errByTimeConver := time.ParseInLocation(layout, timeStr, time.Local)
	if errByTimeConver != nil {
		log.Error("TimeStr To Time Error.....", errByTimeConver)
	}
	return &timeVal
}

// TimeToString 将时间转换为sql查询的字符串格式(防止框架出现增加8小时的问题)
func TimeToString(time *time.Time) *string {
	val := ""
	if nil == time || time.IsZero() {
		return &val
	}
	val = time.Format("2006-01-02 15:04:05")
	return &val
}

func GetNowTimeAddr() *time.Time {
	nTime := time.Now()
	return &nTime
}

func (s *DiseaseService) DiseaseCreate(ctx context.Context, req *v1.DiseaseCreateRequest) (*v1.DiseaseCreateResponse, error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	Log.Println(Token)

	ds, err := s.Disease.CreateDisease(ctx, &chdb.TbDisease{
		ID:          GenNewUUID(),
		Code:        req.Code,
		Name:        req.Name,
		NameJp:      req.NameJp,
		NameQp:      req.NameQp,
		Version:     req.Version,
		Status:      req.Status,
		Tag:         req.Tag,
		Description: req.Description,
		Pid:         req.Pid,
	})

	if err != nil {
		return nil, err
	}
	return &v1.DiseaseCreateResponse{
		Id: ds.ID,
	}, nil
}

func (s *DiseaseService) DiseaseDelete(ctx context.Context, req *v1.DiseaseDeleteRequest) (*v1.DiseaseDeleteResponse, error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	Log.Println(Token)

	deletedID, err := s.Disease.DeleteDisease(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DiseaseDeleteResponse{
		Id: deletedID,
	}, nil
}

func (s *DiseaseService) DiseaseUpdate(ctx context.Context, req *v1.DiseaseUpdateRequest) (res *v1.DiseaseUpdateResponse, errs error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	Log.Println(Token)

	disease, err := s.Disease.UpdateDisease(ctx, &chdb.TbDisease{
		ID:          req.Id,
		Code:        req.Body.Code,
		Description: req.Body.Description,
		Name:        req.Body.Name,
		Pid:         req.Body.Pid,
		Status:      req.Body.Status,
		Tag:         req.Body.Tag,
		Version:     req.Body.Version,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DiseaseUpdateResponse{
		Code:        disease.Code,
		Description: disease.Description,
		Id:          disease.ID,
		Name:        disease.Name,
		NameJp:      disease.NameJp,
		NameQp:      disease.NameQp,
		Pid:         disease.Pid,
		Status:      disease.Status,
		Tag:         disease.Tag,
		Version:     disease.Version,
		CreateTime:  *TimeToString(disease.CreateTime),
		UpdateTime:  *TimeToString(GetNowTimeAddr()),
	}, nil

}

func (s *DiseaseService) DiseaseDetail(ctx context.Context, req *v1.DiseaseDetailRequest) (*v1.DiseaseDetailResponse, error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	Log.Println(Token)

	disease, err := s.Disease.DetailDisease(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DiseaseDetailResponse{
		Code:        disease.Code,
		Description: disease.Description,
		Id:          disease.ID,
		Name:        disease.Name,
		NameJp:      disease.NameJp,
		NameQp:      disease.NameQp,
		Pid:         disease.Pid,
		Status:      disease.Status,
		Tag:         disease.Tag,
		Version:     disease.Version,
		CreateTime:  *TimeToString(disease.CreateTime),
		UpdateTime:  *TimeToString(GetNowTimeAddr()),
	}, nil
}

func (s *DiseaseService) DiseaseFilter(ctx context.Context, apiReq *v1.DiseaseFilterRequest) (resp *v1.DiseaseFilterResponse, errs error) {

	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	Log.Println(Token)

	//resp.Page = apiReq.Page
	//resp.PerPage = apiReq.PerPage
	//resp.Results = make([]*v1.Diseases, 0)
	//
	//tbDiseases, count, err := s.diseaseUseCase.FilterDisease(ctx, apiReq)
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		return nil, err
	//	}
	//	return nil, err
	//}
	//
	//resp.Total = count
	//for _, project := range tbDiseases {
	//	resp.Results = append(resp.Results, ConvertTbDiseaseToGenModel(project))
	//}
	//
	//return &v1.DiseaseFilterResponse{
	//	Page:    resp.Page,
	//	PerPage: resp.PerPage,
	//	Results: resp.Results,
	//	Total:   count,
	//}, err

	diseaselist, count, err := s.Disease.FilterDisease(ctx, apiReq)
	if err != nil {
		return nil, err
	}
	var diseaselistRes []*v1.Diseases

	//循环转换对象
	for _, ds := range diseaselist {
		diseaselistRes = append(diseaselistRes,
			&v1.Diseases{
				Code:        ds.Code,
				Description: ds.Description,
				Id:          ds.ID,
				Name:        ds.Name,
				NameJp:      ds.NameJp,
				NameQp:      ds.NameQp,
				Pid:         ds.Pid,
				Status:      ds.Status,
				Tag:         ds.Tag,
				Version:     ds.Version,
				CreateTime:  *TimeToString(ds.CreateTime),
				UpdateTime:  *TimeToString(ds.UpdateTime),
			})
	}

	return &v1.DiseaseFilterResponse{
		Page:    apiReq.Page,
		PerPage: apiReq.PerPage,
		Results: diseaselistRes,
		Total:   count,
	}, nil
}
