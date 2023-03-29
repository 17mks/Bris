package service

import (
	"context"
	"fmt"
	v1 "followup/api"
	"followup/api/models"
	"followup/gencode/chdb"
	"followup/model"
	"followup/protocol"
	"gorm.io/gorm"
	"log"
	"regexp"
	"strings"
	"time"
)

// ParseTime 将时间字符串转换为time.Time类型
func ParseTime(value string) (*time.Time, error) {
	if "" == value {
		return nil, nil
	}
	// 处理"1971-04-07 00:00:00.0"格式
	if strings.HasSuffix(value, ".0") {
		value = value[0 : len(value)-2]
	}

	layout := "2006-01-02 15:04:05"
	// 时间格式正则判断
	dateTimePatter := `^[1-9]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])\s+(20|21|22|23|[0-1]\d):[0-5]\d:[0-5]\d$`
	datePatter := `^[1-9]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$`

	if match, err := regexp.Match(dateTimePatter, []byte(value)); err != nil {
		return nil, err
	} else if match {
		layout = "2006-01-02 15:04:05"
	}
	if match, err := regexp.Match(datePatter, []byte(value)); err != nil {
		return nil, err
	} else if match {
		layout = "2006-01-02"
	}
	parse, err := time.Parse(layout, value)
	if err != nil {
		return nil, err
	}

	return &parse, nil
}

// ParseTimeIgnoreErr 将时间字符串转换为time.Time类型,忽略错误
func ParseTimeIgnoreErr(value string) *time.Time {
	timeStruct, err := ParseTime(value)
	if err != nil {
		log.Println(err)
		return nil
	}

	return timeStruct
}

func ConvertFollowupWiDetailPreloadToApiResp(preload *model.FollowupWiDetailPreload) *v1.FollowupDetailResponse {
	result := &v1.FollowupDetailResponse{
		WorkItem:   ChdbWkToApiWk(&preload.TbWorkItem),
		PlanRelate: ChdbRlToApiRl(&preload.TbRelate),
		Children:   make([]*v1.FollowupDetailResponse_Children, 0),
	}

	for _, child := range preload.FollowupWiChildren {
		respChild := v1.FollowupDetailResponse_Children{
			WorkItem: ChdbWkToApiWk(&child.TbWorkItem),
			Relates:  make([]*v1.FollowupDetailResponse_RelateS, 0),
		}
		for _, wiRelatePreload := range child.WiRelatePreloads {
			childRelate := v1.FollowupDetailResponse_RelateS{
				Relate:        ChdbRlToApiRl(&wiRelatePreload.TbRelate),
				RelateFormCss: nil,
				WarningNum:    fmt.Sprintf("%d", len(wiRelatePreload.TbFormWarnings)), // 预警信息
			}
			if nil != wiRelatePreload.TbFormCSS {
				fCss := wiRelatePreload.TbFormCSS
				formCss := models.FormCss{
					Id:          fCss.ID,
					Name:        &fCss.Name,
					Status:      &fCss.Status,
					CssCode:     &fCss.CSSCode,
					CssUrl:      &fCss.CSSURL,
					OssUrl:      &fCss.OssURL,
					Description: &fCss.Description,
					VersionName: &fCss.VersionName,
					VersionCode: &fCss.VersionCode,
					CreateTime:  TimeToString(fCss.CreateTime),
					UpdateTime:  TimeToString(fCss.UpdateTime),
					TbFormId:    fCss.TbFormID,
				}
				childRelate.RelateFormCss = &formCss
			}

			respChild.Relates = append(respChild.Relates, &childRelate)
		}

		result.Children = append(result.Children, &respChild)
	}

	return result
}

func ChdbWkToApiWk(req *chdb.TbWorkItem) *v1.WorkItems {
	return &v1.WorkItems{
		Id:                req.ID,
		Title:             req.Title,
		Type:              req.Type,
		Status:            req.Status,
		PrincipalType:     req.PrincipalType,
		PrincipalId:       req.PrincipalID,
		PrincipalName:     req.PrincipalName,
		Participant:       req.Participant,
		Cc:                req.Cc,
		Tag:               req.Tag,
		Pid:               req.Pid,
		AssignedType:      req.AssignedType,
		AssignedTo:        req.AssignedTo,
		AssignedToName:    req.AssignedToName,
		PlanStartTime:     *TimeToString(req.PlanStartTime),
		PlanEndTime:       *TimeToString(req.PlanEndTime),
		ActualStartTime:   *TimeToString(req.ActualStartTime),
		ActualEndTime:     *TimeToString(req.ActualEndTime),
		CreatBy:           req.CreatBy,
		CreatByName:       req.CreatByName,
		UpdateBy:          req.UpdateBy,
		BelongType:        req.BelongType,
		BelongTo:          req.BelongTo,
		SortNum:           req.SortNum,
		Event:             req.Event,
		FrequencyInterval: req.FrequencyInterval,
		FrequencyUnit:     req.FrequencyUnit,
		NotifyLeftOffset:  req.NotifyLeftOffset,
		NotifyRightOffset: req.NotifyRightOffset,
		NotifyOffsetUnit:  req.NotifyOffsetUnit,
		NotifyLeftDate:    *TimeToString(req.NotifyLeftDate),
		NotifyRightDate:   *TimeToString(req.NotifyRightDate),
		NotifyNode:        req.NotifyNode,
		ExecArea:          req.ExecArea,
		Description:       req.Description,
		CreateTime:        *TimeToString(req.CreateTime),
		UpdateTime:        *TimeToString(req.UpdateTime),
		DeletedAt:         *TimeToString(req.DeletedAt),
	}
}

func ChdbRlToApiRl(req *chdb.TbRelate) *v1.Relates {
	return &v1.Relates{
		Id:           req.ID,
		Title:        req.Title,
		Status:       req.Status,
		Conclusion:   req.Conclusion,
		ResourceId:   req.ResourceID,
		ResourceType: req.ResourceType,
		Suggestion:   req.Suggestion,
		TbWorkItemId: req.TbWorkItemID,
		CreateTime:   *TimeToString(&req.CreateTime),
		UpdateTime:   *TimeToString(&req.UpdateTime),
	}
}

func ChdbPtToApiPt(req *chdb.TbPatient) *v1.Patients {
	return &v1.Patients{
		AdtaTime:        *TimeToString(req.AdtaTime),
		AdtdTime:        *TimeToString(req.AdtdTime),
		BirthdayDate:    *TimeToString(req.BirthdayDate),
		Certified:       req.Certified,
		CreateTime:      *TimeToString(req.CreateTime),
		DeptId:          req.DeptID,
		DeptName:        req.DeptName,
		Email:           req.Email,
		Gender:          req.Gender,
		Id:              req.ID,
		IdNumber:        req.IDNumber,
		IndeptTime:      *TimeToString(req.IndeptTime),
		InpBedNo:        req.InpBedNo,
		InpPnurs:        req.InpPnurs,
		InpPnursId:      req.InpPnursID,
		InpWardId:       req.InpWardID,
		InpWardName:     req.InpWardName,
		Inpno:           req.Inpno,
		InHospital:      req.InHospital,
		Mobile:          req.Mobile,
		Name:            req.Name,
		NamePy:          req.NamePy,
		Nation:          req.Nation,
		OuterPatientId:  req.OuterPatientID,
		OuterPatientOrg: req.OuterPatientOrg,
		Outpno:          req.Outpno,
		OuterPlatform:   req.OuterPlatform,
		PatVisitType:    req.PatVisitType,
		PatAtdpscn:      req.PatAtdpscn,
		PatAtdpscnName:  req.PatAtdpscnName,
		PatInpTimes:     req.PatInpTimes,
		SurgeryTime:     *TimeToString(req.SurgeryTime),
		UpdateTime:      *TimeToString(req.UpdateTime),
		VisitCardNo:     req.VisitCardNo,
	}
}

func ConvertModelToApi(req *model.WorkItemDetailBean) *v1.WorkItemDetailBean {
	result := &v1.WorkItemDetailBean{
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

func (s *FollowupService) FollowupCreate(ctx context.Context, apiReq *v1.FollowupCreateRequest) (*v1.FollowupCreateResponse, error) {
	// 获取TOKEN信息
	contextHeader, err := NewSBasisService().ParseContextHeader(ctx)
	if err != nil {
		return nil, err
	}
	userId := contextHeader.TokenInfo.UserId
	userName := contextHeader.TokenInfo.UserName
	// TODO 校验请求参数
	if err := s.checkFollowupCreateRequest(apiReq); err != nil {
		return nil, err
	}
	// 查询方案详情
	planPreload, err := s.planUseCase.PlanPreload(ctx, apiReq.PlanId)
	if err != nil {
		return nil, err
	}
	// 初始化待入库数据
	tbWorkItems := make([]chdb.TbWorkItem, 0)     // 工作项【根任务及子任务】
	tbWorkItemRelates := make([]chdb.TbRelate, 0) // 工作项关联
	// 初始化本次随访计划的根任务
	parentWorkItem := chdb.TbWorkItem{
		ID:             GenNewWorkItemId(),
		Title:          fmt.Sprintf("'%s'的随访计划", apiReq.PatientName),
		Type:           protocol.WorkItemTypeFollowUp, // 随访工作项
		Status:         protocol.WorkItemStatusNew,    // 状态为新建-待处理
		PrincipalType:  apiReq.PrincipalType,
		PrincipalID:    apiReq.PrincipalId,
		PrincipalName:  apiReq.PrincipalName,
		AssignedType:   protocol.WorkItemAssignedTypePatient,
		AssignedTo:     apiReq.PatientId,
		AssignedToName: apiReq.PatientName,
		PlanStartTime:  ParseTimeIgnoreErr(apiReq.EventStartTime), // 计划开始时间
		CreatBy:        userId,
		CreatByName:    userName,
		BelongType:     protocol.WorkItemBelongTypeNone, // 工作项归属 - 目前没有项目
		Event:          planPreload.TbPlan.Event,
	}
	// 如果前端未传输负责人则默认当前操作者即为工作项负责人
	if "" == parentWorkItem.PrincipalID {
		parentWorkItem.PrincipalType = protocol.WorkItemPrincipalTypeUser
		parentWorkItem.PrincipalID = userId
		parentWorkItem.PrincipalName = userName
	}

	relateWorkItems := planPreload.RelateWorkItems
	for _, relateWorkItem := range relateWorkItems {
		tbWorkItem := relateWorkItem.PlanWorkItem.TbWorkItem
		tbRelates := relateWorkItem.PlanWorkItem.TbRelates

		tbWorkItem.ID = GenNewUUID() // 新的任务ID
		tbWorkItem.Status = parentWorkItem.Status
		tbWorkItem.Pid = parentWorkItem.ID
		// 负责人信息
		tbWorkItem.PrincipalType = parentWorkItem.PrincipalType
		tbWorkItem.PrincipalID = parentWorkItem.PrincipalID
		tbWorkItem.PrincipalName = parentWorkItem.PrincipalName
		// 指派人信息
		tbWorkItem.AssignedType = parentWorkItem.AssignedType
		tbWorkItem.AssignedTo = parentWorkItem.AssignedTo
		tbWorkItem.AssignedToName = parentWorkItem.AssignedToName
		// 计划执行时间
		tbWorkItem.PlanStartTime = parentWorkItem.PlanStartTime
		tbWorkItem.Event = parentWorkItem.Event

		tbWorkItems = append(tbWorkItems, tbWorkItem) // 添加到待添加子任务列表
		//	 更新工作项关联资源
		for _, tbRelate := range tbRelates {
			tbRelate.ID = GenNewUUID()
			tbRelate.Status = protocol.RelateStatusNew
			tbRelate.TbWorkItemID = tbWorkItem.ID
			tbWorkItemRelates = append(tbWorkItemRelates, tbRelate)
		}
	}
	// 记录随访工作项执行的是哪个方案
	workItemRelatePlan := chdb.TbRelate{
		ID:           GenNewUUID(),
		Title:        planPreload.TbPlan.Name,
		Status:       protocol.RelateStatusClosed, // 方案状态为已执行
		ResourceType: protocol.ResourceTypePlan,   // 工作项关联资源类型
		ResourceID:   planPreload.TbPlan.ID,
		Conclusion:   "",
		Suggestion:   "",
		Comments:     "",
		CreateTime:   time.Time{},
		UpdateTime:   time.Time{},
		TbWorkItemID: parentWorkItem.ID,
	}
	tbWorkItemRelates = append(tbWorkItemRelates, workItemRelatePlan)

	// 将团队信息写入数据库
	tbWorkItems = append(tbWorkItems, parentWorkItem) // 添加根任务
	// 随访计划入库
	if err := s.WorkItem.WorkItemCreateWithRelates(ctx, tbWorkItems, tbWorkItemRelates); err != nil {
		return nil, err
	}

	return &v1.FollowupCreateResponse{
		Id: parentWorkItem.ID,
	}, nil

}

func (s *FollowupService) FollowupDelete(ctx context.Context, req *v1.FollowupDeleteRequest) (*v1.FollowupDeleteResponse, error) {
	//token, err := NewSBasisService().ParseContextHeader(ctx)
	//if err!=nil{
	//	return nil, err
	//}

	// 递归删除随访工作项
	if _, err := s.Followup.FollowupWiDelByIdRecursion(ctx, req.Id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	followupDeleteResponse := v1.FollowupDeleteResponse{
		Id: req.Id,
	}
	return &followupDeleteResponse, nil
}

func (s *FollowupService) FollowupDetail(ctx context.Context, apiReq *v1.FollowupDetailRequest) (*v1.FollowupDetailResponse, error) {
	// 解析TOKEN
	//contextHeader, err := NewSBasisService().ParseContextHeader(ctx)
	//if err != nil {
	//	return nil, err
	//}
	// 根据工作项ID查询工作项详情及子任务
	wiDetailPreload, err := s.Followup.FollowupWiDetailQueryById(ctx, apiReq.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("followup workitem '%s' not exits", apiReq.Id)
		}
		return nil, err
	}
	log.Println(wiDetailPreload.TbWorkItem.ID)

	// 数据模型转换
	resp := ConvertFollowupWiDetailPreloadToApiResp(wiDetailPreload)

	return &v1.FollowupDetailResponse{
		WorkItem:   resp.WorkItem,
		PlanRelate: resp.PlanRelate,
		Children:   resp.Children,
	}, nil
}

func (s *FollowupService) FollowupFilter(ctx context.Context, apiReq *v1.FollowupFilterRequest) (*v1.FollowupFilterResponse, error) {
	// 解析TOKEN
	contextHeader, err := NewSBasisService().ParseContextHeader(ctx)
	if err != nil {
		return nil, err
	}
	log.Println(contextHeader.Token)

	wiFilterPreloads, count, err := s.Followup.FollowupWiFilterQuery(ctx, apiReq)
	if err != nil {
		return nil, err
	}
	filterPreloads := make([]*v1.FollowupWiFilterPreload, 0)

	//循环转换对象
	for _, fp := range wiFilterPreloads {
		filterPreloads = append(filterPreloads, &v1.FollowupWiFilterPreload{
			WorkItem: &v1.WorkItems{
				ActualEndTime:     *TimeToString(fp.TbWorkItem.ActualEndTime),
				ActualStartTime:   *TimeToString(fp.TbWorkItem.ActualStartTime),
				AssignedTo:        fp.TbWorkItem.AssignedTo,
				AssignedToName:    fp.TbWorkItem.AssignedToName,
				AssignedType:      fp.TbWorkItem.AssignedType,
				BelongTo:          fp.TbWorkItem.BelongTo,
				BelongType:        fp.TbWorkItem.BelongType,
				Cc:                fp.TbWorkItem.Cc,
				CreatBy:           fp.TbWorkItem.CreatBy,
				CreatByName:       fp.TbWorkItem.CreatByName,
				Description:       fp.TbWorkItem.Description,
				Event:             fp.TbWorkItem.Description,
				FrequencyInterval: fp.TbWorkItem.FrequencyInterval,
				FrequencyUnit:     fp.TbWorkItem.FrequencyUnit,
				Participant:       fp.TbWorkItem.Participant,
				Pid:               fp.TbWorkItem.Pid,
				PlanEndTime:       *TimeToString(fp.TbWorkItem.PlanEndTime),
				PlanStartTime:     *TimeToString(fp.TbWorkItem.PlanStartTime),
				PrincipalId:       fp.TbWorkItem.PrincipalID,
				PrincipalName:     fp.TbWorkItem.PrincipalName,
				PrincipalType:     fp.TbWorkItem.PrincipalType,
				SortNum:           fp.TbWorkItem.SortNum,
				Status:            fp.TbWorkItem.Status,
				Tag:               fp.TbWorkItem.Tag,
				Title:             fp.TbWorkItem.Title,
				Type:              fp.TbWorkItem.Type,
				UpdateBy:          fp.TbWorkItem.UpdateBy,
				Id:                fp.TbWorkItem.ID,
				CreateTime:        *TimeToString(fp.TbWorkItem.CreateTime),
				UpdateTime:        *TimeToString(GetNowTimeAddr()),
				DeletedAt:         *TimeToString(fp.TbWorkItem.DeletedAt),
				AppId:             fp.TbWorkItem.AppID,
				NotifyLeftOffset:  fp.TbWorkItem.NotifyLeftOffset,
				NotifyRightOffset: fp.TbWorkItem.NotifyRightOffset,
				NotifyOffsetUnit:  fp.TbWorkItem.NotifyOffsetUnit,
				NotifyLeftDate:    *TimeToString(fp.TbWorkItem.NotifyLeftDate),
				NotifyRightDate:   *TimeToString(fp.TbWorkItem.NotifyRightDate),
				NotifyNode:        fp.TbWorkItem.NotifyNode,
				ExecArea:          fp.TbWorkItem.ExecArea,
			},
			PlanRelate: &v1.Relates{
				Conclusion:   fp.TbRelate.Conclusion,
				CreateTime:   *TimeToString(&fp.TbRelate.CreateTime),
				Id:           fp.TbRelate.ID,
				ResourceId:   fp.TbRelate.ResourceID,
				ResourceType: fp.TbRelate.ResourceType,
				Status:       fp.TbRelate.Status,
				Suggestion:   fp.TbRelate.Suggestion,
				TbWorkItemId: fp.TbRelate.TbWorkItemID,
				Title:        fp.TbRelate.Title,
				UpdateTime:   *TimeToString(&fp.TbRelate.UpdateTime),
			},
		})
	}
	filterResponse := v1.FollowupFilterResponse{
		Page:    apiReq.Page,
		PerPage: apiReq.PerPage,
		Results: filterPreloads,
		Total:   count,
	}
	return &filterResponse, nil

}

// ParseFollowupPlan 解析随访方案，拆解获取具体的工作项
func (s *FollowupService) ParseFollowupPlan(ctx context.Context, tbWorkItem *chdb.TbWorkItem, planId string) ([]model.WorkItemDetailBean, *model.PlanDetailPreloadInfo, error) {
	// 查询方案详情&方案关联的工作项资源
	planDetailPreloadInfo, err := s.planUseCase.PlanDetailPreloadById(ctx, planId)
	if err != nil {
		return nil, nil, err
	}
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
	workItemDetailBeans, err := s.WorkItem.QueryPlanWorkItemByIds(ctx, workItemIds)
	if err != nil {
		return nil, nil, err
	}
	for i := 0; i < len(workItemDetailBeans); i++ {
		// 需要将原来的工作项模板改为具体的工作项数据
		workItemDetailBeans[i].TbWorkItem.ID = GenNewUUID()
		workItemDetailBeans[i].TbWorkItem.Status = protocol.WorkItemStatusNew
		workItemDetailBeans[i].TbWorkItem.Pid = tbWorkItem.ID
		workItemDetailBeans[i].TbWorkItem.AppID = tbWorkItem.AppID
		// 给子任务赋值负责人和被指派人信息
		workItemDetailBeans[i].TbWorkItem.AssignedType = tbWorkItem.AssignedType
		workItemDetailBeans[i].TbWorkItem.AssignedTo = tbWorkItem.AssignedTo
		workItemDetailBeans[i].TbWorkItem.AssignedToName = tbWorkItem.AssignedToName
		workItemDetailBeans[i].TbWorkItem.PrincipalType = tbWorkItem.PrincipalType
		workItemDetailBeans[i].TbWorkItem.PrincipalName = tbWorkItem.PrincipalName
		workItemDetailBeans[i].TbWorkItem.PrincipalID = tbWorkItem.PrincipalID
		// 将模板关联资源改为具体的关联数据
		for j := 0; j < len(workItemDetailBeans[i].TbRelates); j++ {
			workItemDetailBeans[i].TbRelates[j].ID = GenNewUUID()
			workItemDetailBeans[i].TbRelates[j].TbWorkItemID = workItemDetailBeans[i].TbWorkItem.ID
		}
	}

	return workItemDetailBeans, planDetailPreloadInfo, nil
}
