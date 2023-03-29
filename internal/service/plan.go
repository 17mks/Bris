package service

import (
	"context"
	"errors"
	"fmt"
	v1 "followup/api"
	"followup/api/models"
	"followup/gencode/chdb"
	"followup/internal/biz"
	"followup/internal/modelconv"
	"followup/model"
	"followup/protocol"
	"followup/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sort"
	"strings"
	"time"
)

type PlanService struct {
	v1.UnimplementedPlanServer
	planUseCase    *biz.PlanUsecase
	diseaseUseCase *biz.DiseaseUseCase
	DisFunction    *biz.DisFunctionUseCase
	WorkItem       *biz.WorkItemUsecase
	log            *log.Helper
}

func NewPlanService(plan *biz.PlanUsecase, disease *biz.DiseaseUseCase,
	disFunction *biz.DisFunctionUseCase, workItem *biz.WorkItemUsecase, logger log.Logger) *PlanService {
	return &PlanService{
		planUseCase:    plan,
		diseaseUseCase: disease,
		DisFunction:    disFunction,
		WorkItem:       workItem,
		log:            log.NewHelper(logger),
	}
}

func (service *PlanService) PlanCreate(ctx context.Context, apiReq *v1.PlanCreateRequest) (resp *v1.PlanCreateResponse, errs error) {
	contextHeader, err := NewSBasisService().ParseContextHeader(ctx)
	if err != nil {
		return nil, err
	}

	//组装方案信息
	tbPlan := ConvertPlanCreateReqToTbPlanModel(apiReq)
	tbPlan.CreatorID = contextHeader.TokenInfo.UserId
	tbPlan.CreatorName = contextHeader.TokenInfo.UserName
	//组装关联病种数据
	tbPlanRelates := make([]chdb.TbPlanRelate, 0)
	if len(apiReq.RelateDiseases) > 0 {
		tbDiseases, err := service.diseaseUseCase.QueryDiseasesByIds(ctx, apiReq.RelateDiseases)
		if err != nil {
			return nil, err
		}
		// 关联资源组装
		diseaseStr := ""
		for _, tbDisease := range tbDiseases {
			tbPlanRelate := chdb.TbPlanRelate{
				ID:                GenNewPlanRelateId(),
				Title:             tbDisease.Name,
				ResourceType:      protocol.PlanRelateResTypeDisease, // 方案关联的病种 ,
				ResourceID:        tbDisease.ID,
				FrequencyInterval: 0,
				FrequencyOffset:   0,
				Times:             0,
				SortNum:           0,
				CreateTime:        nil,
				UpdateTime:        nil,
				TbPlanID:          tbPlan.ID,
				DeleteAt:          nil,
			}
			tbPlanRelates = append(tbPlanRelates, tbPlanRelate)
			if "" == diseaseStr {
				diseaseStr = tbDisease.Name
			} else {
				diseaseStr = fmt.Sprintf("%s,%s", diseaseStr, tbDisease.Name)
			}
		}
		tbPlan.ApplyDisease = diseaseStr
	}
	// 组装关联功能障碍数据
	if len(apiReq.RelateDysfunction) > 0 {
		//查询待添加的功能障碍信息
		tbDisFunctions, err := service.DisFunction.QueryDisFunByIds(ctx, apiReq.RelateDysfunction)
		if err != nil {
			return nil, err
		}
		applyDysfunction := ""
		//关联资源组装
		for _, tbDisFunc := range tbDisFunctions {
			tbPlanRelate := chdb.TbPlanRelate{
				ID:                GenNewPlanRelateId(),
				Title:             tbDisFunc.ID,
				ResourceType:      protocol.PlanRelateResTypDysfunction, // 方案关联的功能障碍
				ResourceID:        tbDisFunc.ID,
				FrequencyInterval: 0,
				FrequencyOffset:   0,
				Times:             0,
				SortNum:           0,
				CreateTime:        nil,
				UpdateTime:        nil,
				TbPlanID:          tbPlan.ID,
			}
			tbPlanRelates = append(tbPlanRelates, tbPlanRelate)
			if "" == applyDysfunction {
				applyDysfunction = tbDisFunc.Name
			} else {
				applyDysfunction = fmt.Sprintf("%s,%s", applyDysfunction, tbDisFunc.Name)
			}
		}
		tbPlan.ApplyDysfunction = applyDysfunction
	}
	//组装方案适用年龄
	if len(apiReq.ApplyAges) > 0 {
		for _, applyAge := range apiReq.ApplyAges {
			if "" == tbPlan.ApplyAges {
				tbPlan.ApplyAges = applyAge
			} else {
				tbPlan.ApplyAges = fmt.Sprintf("%s,%s", tbPlan.ApplyAges, applyAge)
			}
		}
	}

	tbWorkItems := make([]chdb.TbWorkItem, 0)
	tbWorkItemRelates := make([]chdb.TbRelate, 0)

	for _, workItem := range apiReq.WorkItems {
		tbWorkItem := chdb.TbWorkItem{
			ID:     GenNewWorkItemId(), // 工作项编码
			Title:  *workItem.WorkItem.Title,
			Type:   workItem.WorkItem.Type,
			Status: workItem.WorkItem.Status, // 只能是MODEL，模型
			//BelongType:        workItem.WorkItem.BelongType,
			SortNum:           *workItem.WorkItem.SortNum,
			Event:             workItem.WorkItem.Event,
			FrequencyInterval: *workItem.WorkItem.FrequencyInterval,
			FrequencyUnit:     *workItem.WorkItem.FrequencyUnit,
		}
		for _, relate := range workItem.Relates {
			tbRelate := chdb.TbRelate{
				ID:           GenNewRelateId(),
				Title:        relate.Title,
				Status:       relate.Status,
				ResourceType: relate.ResourceType,
				ResourceID:   relate.ResourceId,
				Conclusion:   "",
				Suggestion:   "",
				Comments:     "",

				CreateTime:   time.Time{},
				UpdateTime:   time.Time{},
				TbWorkItemID: tbWorkItem.ID,
			}
			tbWorkItemRelates = append(tbWorkItemRelates, tbRelate)
		}
		tbWorkItems = append(tbWorkItems, tbWorkItem)
	}
	for _, tbWorkItem := range tbWorkItems {
		tbPlanRelate := chdb.TbPlanRelate{
			ID:                GenNewPlanRelateId(),
			Title:             tbWorkItem.Title,
			ResourceType:      "WI",
			ResourceID:        tbWorkItem.ID,
			FrequencyInterval: 0,
			FrequencyOffset:   0,
			Times:             0,
			SortNum:           0,
			CreateTime:        nil,
			UpdateTime:        nil,
			TbPlanID:          tbPlan.ID,
		}
		tbPlanRelates = append(tbPlanRelates, tbPlanRelate)
	}
	//将方案信息写入数据库
	newPlan, err := service.planUseCase.PlanCreateWithResource(ctx, &tbPlan, tbPlanRelates, tbWorkItems, tbWorkItemRelates)
	if err != nil {
		return nil, err
	}
	return &v1.PlanCreateResponse{
		Id: newPlan.ID,
	}, err
}
func (service *PlanService) PlanDelById(ctx context.Context, req *v1.PlanDelByIdRequest) (*v1.PlanDelByIdResponse, error) {
	deletedID, err := service.planUseCase.PlanDelete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.PlanDelByIdResponse{
		Id: deletedID,
	}, nil
}
func (service *PlanService) PlanUpdate(ctx context.Context, apiReq *v1.PlanUpdateRequest) (*v1.PlanUpdateResponse, error) {
	// 获取Token信息
	//contextHeader, err := NewSBasisService().ParseContextHeader(ctx)
	//if err != nil {
	//	return nil, err
	//}
	// 校验请求参数
	if err := service.checkPlanUpdateRequest(apiReq); err != nil {
		return nil, err
	}

	// 查询方案详情 - 包括方案关联资源的详情
	oldPlanPreload, err := service.planUseCase.PlanPreload(ctx, apiReq.Id)
	if err != nil {
		return nil, err
	}

	toUpdate := model.ToUpdatePlanResource{
		Plan:                      nil,
		ToCreateWorkItem:          []chdb.TbWorkItem{},
		ToDeleteWorkItemIds:       []string{},
		ToUpdateWorkItem:          []chdb.TbWorkItem{},
		ToCreateRelate:            []chdb.TbRelate{},
		ToDeleteWorkItemRelateIds: []string{},
		ToCreatePlanRelate:        []chdb.TbPlanRelate{},
		ToDeletePlanRelate:        []string{},
		ToUpdatePlanRelate:        []chdb.TbPlanRelate{},
	}
	// 方案基本信息 -
	toUpdatePlan := oldPlanPreload.TbPlan
	planBasisUpdate := apiReq.Data.PlanBasisUpdate // 接口请求信息 - 方案基本信息
	toUpdatePlan.Name = planBasisUpdate.Name       // 更新方案名称
	toUpdatePlan.Event = planBasisUpdate.Event     // 更新事件
	toUpdatePlan.UpdateTime = GetNowTimeAddr()     // 更新时间
	// 更新适用年龄段
	applyAgeStr := ""
	for _, applyAge := range planBasisUpdate.ApplyAges {
		if "" == applyAgeStr {
			applyAgeStr = applyAge
		} else {
			applyAgeStr = fmt.Sprintf("%s,%s", applyAgeStr, applyAge)
		}
	}
	toUpdatePlan.ApplyAges = applyAgeStr // 更新适用年龄段
	// 更新功能障碍 - 1.删除旧的功能障碍 2. 添加新的功能障碍
	for _, relateDysfunction := range oldPlanPreload.RelateDysfunctions {
		toUpdate.ToDeletePlanRelate = append(toUpdate.ToDeletePlanRelate, relateDysfunction.TbPlanRelate.ID)
	}
	// 更新功能障碍 - 2. 添加新的功能障碍
	if len(planBasisUpdate.RelateDysfunction) > 0 {
		//查询待添加的功能障碍信息
		tbDisFunctions, err := service.DisFunction.QueryDisFunByIds(ctx, planBasisUpdate.RelateDysfunction)
		if err != nil {
			return nil, err
		}
		applyDysfunction := ""
		//关联资源组装
		for _, disFunction := range tbDisFunctions {
			tbPlanRelate := chdb.TbPlanRelate{
				ID:                GenNewPlanRelateId(),
				Title:             disFunction.Name,
				ResourceType:      protocol.PlanRelateResTypDysfunction, // 方案关联的功能障碍
				ResourceID:        disFunction.ID,
				FrequencyInterval: 0,
				FrequencyOffset:   0,
				Times:             0,
				SortNum:           0,
				CreateTime:        nil,
				UpdateTime:        nil,
				TbPlanID:          toUpdatePlan.ID,
			}
			toUpdate.ToCreatePlanRelate = append(toUpdate.ToCreatePlanRelate, tbPlanRelate) // 待创建的方案关联
			if "" == applyDysfunction {
				applyDysfunction = disFunction.Name
			} else {
				applyDysfunction = fmt.Sprintf("%s,%s", applyDysfunction, disFunction.Name)
			}
		}
		toUpdatePlan.ApplyDysfunction = applyDysfunction // 更新关联功能障碍
	}
	// 更新关联病种 - 1. 删除原有关联病种
	for _, relateDisease := range oldPlanPreload.RelateDiseases {
		toUpdate.ToDeletePlanRelate = append(toUpdate.ToDeletePlanRelate, relateDisease.TbPlanRelate.ID)
	}
	//  更新关联病种 - 2. 添加接口上传的病种
	if len(planBasisUpdate.RelateDiseases) > 0 {
		tbDiseases, err := service.diseaseUseCase.QueryDiseasesByIds(ctx, planBasisUpdate.RelateDiseases)
		if err != nil {
			return nil, err
		}
		// 关联资源组装
		diseaseStr := ""
		for _, tbDisease := range tbDiseases {
			tbPlanRelate := chdb.TbPlanRelate{
				ID:                GenNewPlanRelateId(),
				Title:             tbDisease.Name,
				ResourceType:      protocol.PlanRelateResTypeDisease, // 方案关联的病种 ,
				ResourceID:        tbDisease.ID,
				FrequencyInterval: 0,
				FrequencyOffset:   0,
				Times:             0,
				SortNum:           0,
				CreateTime:        nil,
				UpdateTime:        nil,
				TbPlanID:          toUpdatePlan.ID,
				DeleteAt:          nil,
			}
			toUpdate.ToCreatePlanRelate = append(toUpdate.ToCreatePlanRelate, tbPlanRelate)
			if "" == diseaseStr {
				diseaseStr = tbDisease.Name
			} else {
				diseaseStr = fmt.Sprintf("%s,%s", diseaseStr, tbDisease.Name)
			}
		}
		toUpdatePlan.ApplyDisease = diseaseStr // 更新方案关联病种
	}

	// 获取需要添加&更新的工作项
	planWorkItemUpdate := apiReq.GetData().PlanWorkItemUpdate
	for _, workItem := range planWorkItemUpdate.WorkItems {
		if nil == workItem.WorkItem {
			continue
		}
		var tbWorkItem *chdb.TbWorkItem
		if "" == workItem.WorkItem.GetId() {
			// 没有工作项ID，是待添加的工作项
			tbWorkItem = &chdb.TbWorkItem{
				ID:                GenNewWorkItemId(),
				Title:             workItem.WorkItem.GetTitle(),
				Type:              workItem.WorkItem.Type,
				Status:            workItem.WorkItem.Status, // 只能是MODEL，模型
				PrincipalType:     "NONE",
				Tag:               workItem.WorkItem.GetTag(),
				AssignedType:      "NONE",
				BelongType:        "NONE",
				SortNum:           workItem.WorkItem.GetSortNum(),
				Event:             workItem.WorkItem.Event,
				FrequencyInterval: workItem.WorkItem.GetFrequencyInterval(),
				FrequencyUnit:     workItem.WorkItem.GetFrequencyUnit(),
				AppID:             "",
				NotifyOffsetUnit:  "NONE",
				ExecArea:          "NONE",
				Description:       workItem.WorkItem.GetDescription(),
			}
			toUpdate.ToCreateWorkItem = append(toUpdate.ToCreateWorkItem, *tbWorkItem)
			// 添加工作项关联
			relateWorkItem := chdb.TbPlanRelate{
				ID:                GenNewUUID(),
				Title:             tbWorkItem.Title,
				ResourceType:      protocol.PlanRelateResTypeWorkItem,
				ResourceID:        tbWorkItem.ID,
				FrequencyInterval: 0,
				FrequencyOffset:   0,
				Times:             0,
				SortNum:           0,
				CreateTime:        nil,
				UpdateTime:        nil,
				TbPlanID:          oldPlanPreload.TbPlan.ID,
				DeleteAt:          nil,
			}
			toUpdate.ToCreatePlanRelate = append(toUpdate.ToCreatePlanRelate, relateWorkItem)

			for _, relate := range workItem.Relates {
				// 新的工作项关联
				tbRelate := chdb.TbRelate{
					ID:           GenNewUUID(),
					Title:        relate.Title,
					Status:       relate.Status,
					ResourceType: relate.ResourceType,
					ResourceID:   relate.ResourceId,
					Conclusion:   "",
					Suggestion:   "",
					Comments:     "",
					TbWorkItemID: tbWorkItem.ID,
				}
				toUpdate.ToCreateRelate = append(toUpdate.ToCreateRelate, tbRelate)

			}

		} else {
			// 有工作项ID，是待更新的工作项
			for _, planRelateWorkItem := range oldPlanPreload.RelateWorkItems {
				oldWID := planRelateWorkItem.PlanWorkItem.TbWorkItem.ID
				if oldWID == *workItem.WorkItem.Id {
					tbWorkItem = &planRelateWorkItem.PlanWorkItem.TbWorkItem
				}
			}
			if nil != tbWorkItem {
				// 更新工作项
				tbWorkItem.Event = workItem.WorkItem.Event
				tbWorkItem.FrequencyInterval = workItem.WorkItem.GetFrequencyInterval()
				tbWorkItem.FrequencyUnit = workItem.WorkItem.GetFrequencyUnit()
				tbWorkItem.SortNum = workItem.WorkItem.GetSortNum()
				tbWorkItem.Title = workItem.WorkItem.GetTitle()
				tbWorkItem.Status = workItem.WorkItem.GetStatus()
				tbWorkItem.Tag = workItem.WorkItem.GetTag()
				tbWorkItem.Description = workItem.WorkItem.GetDescription()
				tbWorkItem.Type = workItem.WorkItem.GetType()
				tbWorkItem.UpdateTime = GetNowTimeAddr()

				toUpdate.ToUpdateWorkItem = append(toUpdate.ToUpdateWorkItem, *tbWorkItem)

				// 处理工作项关联
				for _, relate := range workItem.Relates {
					if "" == relate.GetId() {
						// 新的工作项关联
						tbRelate := chdb.TbRelate{
							ID:           GenNewUUID(),
							Title:        relate.Title,
							Status:       relate.Status,
							ResourceType: relate.ResourceType,
							ResourceID:   relate.ResourceId,
							Conclusion:   "",
							Suggestion:   "",
							Comments:     "",
							TbWorkItemID: workItem.WorkItem.GetId(),
						}
						toUpdate.ToCreateRelate = append(toUpdate.ToCreateRelate, tbRelate)
					} else {
						var toUpdateRelate *chdb.TbRelate
						for _, relateWorkItem := range oldPlanPreload.RelateWorkItems {
							for _, tbRelate := range relateWorkItem.PlanWorkItem.TbRelates {
								if *relate.Id == tbRelate.ID {
									toUpdateRelate = &tbRelate
								}
							}
						}
						// 待更新的工作项关联
						if nil != toUpdateRelate {
							toUpdateRelate.TbWorkItemID = workItem.WorkItem.GetId()
							toUpdateRelate.Title = relate.Title
							toUpdateRelate.Status = relate.Status
							toUpdateRelate.ResourceType = relate.ResourceType
							toUpdateRelate.ResourceID = relate.ResourceId
							toUpdateRelate.UpdateTime = *GetNowTimeAddr()

							toUpdate.ToUpdateRelate = append(toUpdate.ToUpdateRelate, *toUpdateRelate)
						}
					}
				}

			}

		}

	}
	// 将被删除的工作项与方案的关联删除
	for _, workItemId := range planWorkItemUpdate.PlanDeletedWorkItemInfo.WorkItemIds {
		for _, planRelate := range oldPlanPreload.TbPlanRelates {
			if planRelate.ResourceID == workItemId {
				toUpdate.ToDeletePlanRelate = append(toUpdate.ToDeletePlanRelate, planRelate.ID)
			}
		}
		toUpdate.ToDeleteWorkItemIds = append(toUpdate.ToDeleteWorkItemIds, workItemId)
	}
	for _, workItemRelateId := range planWorkItemUpdate.PlanDeletedWorkItemInfo.WorkItemRelateIds {
		toUpdate.ToDeleteWorkItemRelateIds = append(toUpdate.ToDeleteWorkItemRelateIds, workItemRelateId)
	}

	// 执行更新
	toUpdate.Plan = &toUpdatePlan

	tbPlan, err := service.planUseCase.PlanUpdateWithResource(ctx, &toUpdate)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	// 结果赋值
	planUpdateResponse := ConvertTbPlanToUpdateApiResp(tbPlan)
	return planUpdateResponse, nil

}
func (service *PlanService) PlanDetail(ctx context.Context, apiReq *v1.PlanDetailRequest) (*v1.PlanDetailResponse, error) {
	// TODO 解析TOKEN，这里需要统一处理【暂时这样吧...】
	serverContext, ok := transport.FromServerContext(ctx)
	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}
	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Debug(Token)

	// 查询方案及方案关联资源
	planPreload, err := service.planUseCase.PlanPreload(ctx, apiReq.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	planDetailResponse := v1.PlanDetailResponse{
		Plan:               modelconv.ConvTbPlanToModelsPlan(planPreload.TbPlan),
		RelateWorkItems:    make([]*v1.PlanRelateWorkItem, 0),
		RelateDiseases:     make([]*v1.PlanRelateDisease, 0),
		RelateDysfunctions: make([]*v1.PlanRelateDysfunction, 0),
	}
	for _, relateWorkItem := range planPreload.RelateWorkItems {
		planRelateWorkItem := v1.PlanRelateWorkItem{
			PlanRelate: modelconv.ConvTbPlanRelateToModelsPlanRelate(relateWorkItem.TbPlanRelate),
			WorkItem:   modelconv.ConvTbWorkItemToModelsWorkItem(relateWorkItem.PlanWorkItem.TbWorkItem),
			Relates:    make([]*models.Relate, 0),
		}
		for _, tbRelate := range relateWorkItem.PlanWorkItem.TbRelates {
			relate := modelconv.ConvTbRelateToModelsRelate(tbRelate)
			planRelateWorkItem.Relates = append(planRelateWorkItem.Relates, relate)
		}

		planDetailResponse.RelateWorkItems = append(planDetailResponse.RelateWorkItems, &planRelateWorkItem)
	}
	for _, relateDisease := range planPreload.RelateDiseases {
		planRelateDisease := v1.PlanRelateDisease{
			PlanRelate: modelconv.ConvTbPlanRelateToModelsPlanRelate(relateDisease.TbPlanRelate),
			Disease:    modelconv.ConvTbDiseaseToModelsDisease(relateDisease.TbDisease),
		}
		planDetailResponse.RelateDiseases = append(planDetailResponse.RelateDiseases, &planRelateDisease)
	}
	for _, planRelateDysfunction := range planPreload.RelateDysfunctions {
		planRelateDisease := v1.PlanRelateDysfunction{
			PlanRelate: modelconv.ConvTbPlanRelateToModelsPlanRelate(planRelateDysfunction.TbPlanRelate),
			DisFunc:    modelconv.ConvTbDisFuncToModelsDisFunc(planRelateDysfunction.TbDisFunc),
		}
		planDetailResponse.RelateDysfunctions = append(planDetailResponse.RelateDysfunctions, &planRelateDisease)
	}

	return &planDetailResponse, nil
}

func (service *PlanService) PlanFilterQuery(ctx context.Context, apiReq *v1.PlanFilterQueryRequest) (resp *v1.PlanFilterQueryResponse, errs error) {
	serverContext, ok := transport.FromServerContext(ctx)
	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Debug(Token)

	planList, count, err := service.planUseCase.FilterPlan(ctx, apiReq)
	if err != nil {
		return nil, err
	}
	planFilterQueryResponse := v1.PlanFilterQueryResponse{
		Page:    apiReq.Page,
		PerPage: apiReq.PerPage,
		Results: make([]*models.Plan, 0),
		Total:   count,
	}
	data := make([]*models.Plan, 0)

	// TODO for遍历和for-range遍历的区别

	//循环转换对象
	for i := 0; i < len(planList); i++ {
		pl := planList[i]
		plan := models.Plan{
			Id:               pl.ID,
			AppId:            &pl.AppID,
			NotifyNode:       &pl.NotifyNode,
			Name:             pl.Name,
			Type:             pl.Type,
			Status:           pl.Status,
			BelongType:       pl.BelongType,
			BelongTo:         &pl.BelongTo,
			ApplyAges:        &pl.ApplyAges,
			ApplyDisease:     &pl.ApplyDisease,
			ApplyDysfunction: &pl.ApplyDysfunction,
			CreatorId:        &pl.CreatorID,
			CreatorName:      &pl.CreatorName,
			Event:            &pl.Event,
			CreateTime:       TimeToString(pl.CreateTime),
			UpdateTime:       TimeToString(GetNowTimeAddr()),
		}
		data = append(data, &plan)
	}
	planFilterQueryResponse.Results = append(planFilterQueryResponse.Results, data...)
	return &planFilterQueryResponse, nil
}
func (service *PlanService) QueryPlanWorkItem(ctx context.Context, planDetailPreloadInfo *model.PlanDetailPreloadInfo) ([]*v1.WorkItemDetailGenInfo, error) {
	workItemIds := make([]string, 0)
	for _, tbPlanRelate := range planDetailPreloadInfo.TbPlanRelates {
		switch tbPlanRelate.ResourceType {
		case "WI":
			if "" != tbPlanRelate.ResourceID {
				workItemIds = append(workItemIds, tbPlanRelate.ResourceID)
			}
		default:
			// 其余资源类型先暂时忽略
		}
	}

	workItemDetailBeans, err := service.WorkItem.QueryPlanWorkItemByIds(ctx, workItemIds)
	if err != nil {
		return nil, err
	}

	// 根据sortNum排序
	sort.SliceStable(workItemDetailBeans, func(i, j int) bool {
		return workItemDetailBeans[i].TbWorkItem.SortNum < workItemDetailBeans[j].TbWorkItem.SortNum
	})
	itemDetailGenInfos := make([]*v1.WorkItemDetailGenInfo, 0)
	for _, tbDisease := range workItemDetailBeans {
		workItemDetailGenInfo := ConWorkItemDetailBenToGenModel(&tbDisease)
		itemDetailGenInfos = append(itemDetailGenInfos, workItemDetailGenInfo)
	}

	return itemDetailGenInfos, nil
}
func (service *PlanService) QueryPlanDisease(ctx context.Context, planDetailPreloadInfo *model.PlanDetailPreloadInfo) ([]*v1.Diseases, error) {
	diseaseIds := make([]string, 0)
	for _, tbPlanRelate := range planDetailPreloadInfo.TbPlanRelates {
		switch tbPlanRelate.ResourceType {
		case "BZ":
			if "" != tbPlanRelate.ResourceID {
				diseaseIds = append(diseaseIds, tbPlanRelate.ResourceID)
			}
		default:
			// 其余资源类型先暂时忽略
		}
	}

	tbDiseases, err := service.diseaseUseCase.QueryDiseasesByIds(ctx, diseaseIds)
	if err != nil {
		return nil, err
	}
	diseases := make([]*v1.Diseases, 0)
	for _, tbDisease := range tbDiseases {
		disease := ConvertDiseaseToTbDisease(&tbDisease)
		diseases = append(diseases, disease)
	}

	return diseases, nil
}

func (service *PlanService) checkPlanUpdateRequest(request *v1.PlanUpdateRequest) error {
	if nil == request.Data {
		return errors.New("data can't be empty")
	}
	return nil
}

func GenNewUUID() string {
	userId := strings.ReplaceAll(uuid.NewString(), "-", "")
	return strings.ToLower(userId)
}

// GenNewPlanRelateId 生成方案关联编码
func GenNewPlanRelateId() string {
	userId := strings.ReplaceAll(uuid.NewString(), "-", "")
	return strings.ToUpper(fmt.Sprintf("PLR%s", userId))
}

func GenNewWorkItemId() string {
	userId := strings.ReplaceAll(uuid.NewString(), "-", "")
	return strings.ToUpper(fmt.Sprintf("WI%s", userId))
}

func GenNewRelateId() string {
	userId := strings.ReplaceAll(uuid.NewString(), "-", "")
	return strings.ToUpper(fmt.Sprintf("RL%s", userId))
}

func ConvertPlanCreateReqToTbPlanModel(req *v1.PlanCreateRequest) chdb.TbPlan {
	return chdb.TbPlan{
		ID:               GenNewUUID(),
		Name:             req.Name,
		Type:             req.Type,
		Status:           req.Status,
		BelongType:       req.BelongType,
		BelongTo:         *req.BelongTo,
		ApplyDisease:     "",
		ApplyDysfunction: "",
		ApplyAges:        "",
		Event:            *req.Event,
		CreatorID:        "",
		CreatorName:      "",
		AppID:            "",
		NotifyNode:       "",
		CreateTime:       nil,
		UpdateTime:       nil,
		DeleteAt:         nil,
	}
}

func ConvertPlanUpdateReqToTbPlanModel(req *v1.PlanUpdateReq) chdb.TbPlan {
	return chdb.TbPlan{
		ID:         "",
		Type:       req.Type,
		Name:       req.Name,
		BelongType: req.BelongType,
		BelongTo:   req.BelongTo,
		CreateTime: nil,
		UpdateTime: nil,
	}
}

func ConvertTbPlanToUpdateApiResp(req *chdb.TbPlan) *v1.PlanUpdateResponse {
	return &v1.PlanUpdateResponse{
		BelongTo:   req.BelongTo,
		BelongType: req.BelongType,
		CreateTime: *TimeToString(req.CreateTime),
		Id:         req.ID,
		Name:       req.Name,
		Status:     req.Status,
		Type:       req.Type,
		UpdateTime: *TimeToString(req.UpdateTime),
	}
}
func ConWorkItemDetailBenToGenModel(req *model.WorkItemDetailBean) *v1.WorkItemDetailGenInfo {
	result := &v1.WorkItemDetailGenInfo{
		WorkItems: &v1.WorkItems{
			Id:                req.TbWorkItem.ID,
			Title:             req.TbWorkItem.Title,
			Type:              req.TbWorkItem.Type,
			Status:            req.TbWorkItem.Status,
			PrincipalType:     req.TbWorkItem.PrincipalType,
			PrincipalId:       req.TbWorkItem.PrincipalID,
			PrincipalName:     req.TbWorkItem.PrincipalName,
			Participant:       req.TbWorkItem.Participant,
			Cc:                req.TbWorkItem.Cc,
			Tag:               req.TbWorkItem.Tag,
			Pid:               req.TbWorkItem.Pid,
			AssignedType:      req.TbWorkItem.AssignedType,
			AssignedTo:        req.TbWorkItem.AssignedTo,
			AssignedToName:    req.TbWorkItem.AssignedToName,
			PlanStartTime:     *TimeToString(req.TbWorkItem.PlanStartTime),
			PlanEndTime:       *TimeToString(req.TbWorkItem.PlanEndTime),
			ActualStartTime:   *TimeToString(req.TbWorkItem.ActualStartTime),
			ActualEndTime:     *TimeToString(req.TbWorkItem.ActualEndTime),
			CreatBy:           req.TbWorkItem.CreatBy,
			CreatByName:       req.TbWorkItem.CreatByName,
			UpdateBy:          req.TbWorkItem.UpdateBy,
			BelongType:        req.TbWorkItem.BelongType,
			BelongTo:          req.TbWorkItem.BelongTo,
			SortNum:           req.TbWorkItem.SortNum,
			Event:             req.TbWorkItem.Event,
			FrequencyInterval: req.TbWorkItem.FrequencyInterval,
			FrequencyUnit:     req.TbWorkItem.FrequencyUnit,
			NotifyLeftOffset:  req.TbWorkItem.NotifyLeftOffset,
			NotifyRightOffset: req.TbWorkItem.NotifyRightOffset,
			NotifyOffsetUnit:  req.TbWorkItem.NotifyOffsetUnit,
			NotifyLeftDate:    *TimeToString(req.TbWorkItem.NotifyLeftDate),
			NotifyRightDate:   *TimeToString(req.TbWorkItem.NotifyRightDate),
			NotifyNode:        req.TbWorkItem.NotifyNode,
			ExecArea:          req.TbWorkItem.ExecArea,
			Description:       req.TbWorkItem.Description,
			CreateTime:        *TimeToString(req.TbWorkItem.CreateTime),
			UpdateTime:        *TimeToString(req.TbWorkItem.UpdateTime),
			DeletedAt:         *TimeToString(req.TbWorkItem.DeletedAt),
		},
		Relates: make([]*v1.Relates, 0),
	}
	for _, relates := range req.TbRelates {
		relate := v1.Relates{
			Id:           relates.ID,
			Title:        relates.Title,
			Status:       relates.Status,
			Conclusion:   relates.Conclusion,
			ResourceId:   relates.ResourceID,
			ResourceType: relates.ResourceType,
			Suggestion:   relates.Suggestion,
			TbWorkItemId: relates.TbWorkItemID,
			CreateTime:   *TimeToString(&relates.CreateTime),
			UpdateTime:   *TimeToString(&relates.UpdateTime),
		}
		result.Relates = append(result.Relates, &relate)
	}

	return result
}
