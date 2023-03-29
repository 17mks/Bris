package modelconv

import (
	"followup/api/models"
	"followup/gencode/chdb"
	"followup/protocol"
)

func ConvTbPlanToModelsPlan(tbPlan chdb.TbPlan) *models.Plan {
	plan := models.Plan{
		Id:               tbPlan.ID,
		Name:             tbPlan.Name,
		Type:             tbPlan.Type,
		Status:           tbPlan.Status,
		BelongType:       tbPlan.BelongType,
		BelongTo:         &tbPlan.BelongTo,
		ApplyDisease:     &tbPlan.ApplyDisease,
		ApplyDysfunction: &tbPlan.ApplyDysfunction,
		ApplyAges:        &tbPlan.ApplyAges,
		Event:            &tbPlan.Event,
		CreatorId:        &tbPlan.CreatorID,
		CreatorName:      &tbPlan.CreatorName,
		AppId:            &tbPlan.AppID,
		NotifyNode:       nil,
		CreateTime:       protocol.TimeToString(tbPlan.CreateTime),
		UpdateTime:       protocol.TimeToString(tbPlan.UpdateTime),
		DeleteAt:         nil,
	}
	return &plan
}

func ConvTbRelateToModelsRelate(tbRelate chdb.TbRelate) *models.Relate {
	relate := models.Relate{
		Id:           tbRelate.ID,
		Title:        tbRelate.Title,
		Status:       tbRelate.Status,
		ResourceType: tbRelate.ResourceType,
		ResourceId:   tbRelate.ResourceID,
		Conclusion:   &tbRelate.Conclusion,
		Suggestion:   &tbRelate.Suggestion,
		Comments:     &tbRelate.Comments,
		CreateTime:   *protocol.TimeToString(&tbRelate.CreateTime),
		UpdateTime:   *protocol.TimeToString(&tbRelate.UpdateTime),
		TbWorkItemId: tbRelate.TbWorkItemID,
	}
	return &relate
}

func ConvTbWorkItemToModelsWorkItem(tbWorkItem chdb.TbWorkItem) *models.WorkItem {
	workItem := models.WorkItem{
		Id:                tbWorkItem.ID,
		Title:             &tbWorkItem.Title,
		Type:              tbWorkItem.Type,
		Status:            tbWorkItem.Status,
		PrincipalType:     &tbWorkItem.PrincipalType,
		PrincipalId:       &tbWorkItem.PrincipalID,
		PrincipalName:     &tbWorkItem.PrincipalName,
		Participant:       &tbWorkItem.Participant,
		Cc:                &tbWorkItem.Cc,
		Tag:               &tbWorkItem.Tag,
		Pid:               &tbWorkItem.Pid,
		AssignedType:      tbWorkItem.AssignedType,
		AssignedTo:        &tbWorkItem.AssignedTo,
		AssignedToName:    &tbWorkItem.AssignedToName,
		PlanStartTime:     protocol.TimeToString(tbWorkItem.PlanStartTime),
		PlanEndTime:       protocol.TimeToString(tbWorkItem.PlanEndTime),
		ActualStartTime:   protocol.TimeToString(tbWorkItem.ActualStartTime),
		ActualEndTime:     protocol.TimeToString(tbWorkItem.ActualEndTime),
		CreatBy:           &tbWorkItem.CreatBy,
		CreatByName:       &tbWorkItem.CreatByName,
		UpdateBy:          &tbWorkItem.UpdateBy,
		BelongType:        tbWorkItem.BelongType,
		BelongTo:          &tbWorkItem.BelongTo,
		SortNum:           &tbWorkItem.SortNum,
		Event:             tbWorkItem.Event,
		FrequencyInterval: &tbWorkItem.FrequencyInterval,
		FrequencyUnit:     &tbWorkItem.FrequencyUnit,
		AppId:             &tbWorkItem.AppID,
		NotifyLeftOffset:  &tbWorkItem.NotifyLeftOffset,
		NotifyRightOffset: &tbWorkItem.NotifyRightOffset,
		NotifyOffsetUnit:  &tbWorkItem.NotifyOffsetUnit,
		NotifyLeftDate:    protocol.TimeToString(tbWorkItem.NotifyLeftDate),
		NotifyRightDate:   protocol.TimeToString(tbWorkItem.NotifyRightDate),
		NotifyNode:        &tbWorkItem.NotifyNode,
		ExecArea:          &tbWorkItem.ExecArea,
		Description:       &tbWorkItem.Description,
		CreateTime:        protocol.TimeToString(tbWorkItem.CreateTime),
		UpdateTime:        protocol.TimeToString(tbWorkItem.UpdateTime),
		DeletedAt:         protocol.TimeToString(tbWorkItem.DeletedAt),
	}

	return &workItem
}

func ConvTbPlanRelateToModelsPlanRelate(tbPlanRelate chdb.TbPlanRelate) *models.PlanRelate {
	planRelate := models.PlanRelate{
		Id:                tbPlanRelate.ID,
		Title:             tbPlanRelate.Title,
		ResourceType:      tbPlanRelate.ResourceType,
		ResourceId:        tbPlanRelate.ResourceID,
		FrequencyInterval: &tbPlanRelate.FrequencyInterval,
		FrequencyOffset:   &tbPlanRelate.FrequencyOffset,
		Times:             &tbPlanRelate.Times,
		SortNum:           &tbPlanRelate.SortNum,
		CreateTime:        *protocol.TimeToString(tbPlanRelate.CreateTime),
		UpdateTime:        *protocol.TimeToString(tbPlanRelate.UpdateTime),
		TbPlanId:          tbPlanRelate.TbPlanID,
		DeleteAt:          protocol.TimeToString(tbPlanRelate.DeleteAt),
	}

	return &planRelate
}

func ConvTbDisFuncToModelsDisFunc(tbDisFunc *chdb.TbDisFunc) *models.DisFunc {
	if nil == tbDisFunc {
		return nil
	}
	disFunc := models.DisFunc{
		Id:          tbDisFunc.ID,
		Name:        tbDisFunc.Name,
		Description: &tbDisFunc.Description,
		Py:          protocol.TimeToString(tbDisFunc.Py),
		CreateTime:  *protocol.TimeToString(tbDisFunc.CreateTime),
		UpdateTime:  *protocol.TimeToString(tbDisFunc.UpdateTime),
		DeleteAt:    protocol.TimeToString(tbDisFunc.DeleteAt),
	}

	return &disFunc
}

func ConvTbDiseaseToModelsDisease(tbDisease *chdb.TbDisease) *models.Disease {
	if nil == tbDisease {
		return nil
	}
	disease := models.Disease{
		Id:          tbDisease.ID,
		Code:        &tbDisease.Code,
		Name:        tbDisease.Name,
		NameJp:      tbDisease.NameJp,
		NameQp:      tbDisease.NameQp,
		Version:     tbDisease.Version,
		Status:      &tbDisease.Status,
		Tag:         &tbDisease.Tag,
		Description: &tbDisease.Description,
		Pid:         &tbDisease.Pid,
		CreateTime:  protocol.TimeToString(tbDisease.CreateTime),
		UpdateTime:  protocol.TimeToString(tbDisease.UpdateTime),
		DeleteAt:    protocol.TimeToString(tbDisease.DeleteAt),
	}

	return &disease
}
