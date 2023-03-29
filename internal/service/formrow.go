package service

import (
	"context"
	"fmt"
	v1 "followup/api"
	"followup/gencode/chdb"
	"followup/model"
	"followup/msgcenter"
	"followup/utils"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ConvertFormRowCreateReqToTbFormRow(req *v1.FormRowCreateRequest) chdb.TbFormRow {
	return chdb.TbFormRow{
		ID:            GenNewUUID(),
		Title:         req.Title,
		Status:        req.Status,
		Data:          req.Data,
		Remark:        req.Remark,
		Submitter:     "",
		TbFormID:      req.TbFormId,
		TbRelateID:    req.TbRelateId,
		SignaturePath: req.SignaturePath,
		SignerID:      req.SignerId,
		SignerName:    req.SignerName,
		SignOnBehalf:  req.SignOnBehalf,
		CreateTime:    nil,
		UpdateTime:    nil,
	}
}

func ConvertFormRowUpdateReqToTbFormRowModel(req *v1.FormRowUpdateRequest) chdb.TbFormRow {
	return chdb.TbFormRow{
		ID:         "",
		Title:      req.Body.Title,
		Status:     req.Body.Status,
		Data:       req.Body.Data,
		Remark:     req.Body.Remark,
		TbFormID:   req.Body.TbFormId,
		TbRelateID: req.Body.TbRelateId,
	}
}

func ConvertTbFormRowToUpdateApiRespModel(req *chdb.TbFormRow) v1.FormRowUpdateResponse {
	return v1.FormRowUpdateResponse{
		CreateTime: *TimeToString(req.CreateTime),
		Data:       req.Data,
		Id:         req.ID,
		Remark:     req.Remark,
		Status:     req.Status,
		TbFormId:   req.TbFormID,
		TbRelateId: req.TbRelateID,
		Title:      req.Title,
		UpdateTime: *TimeToString(req.UpdateTime),
	}
}

func ConvertTbFormRowToDetailQueryApiRespModel(req *chdb.TbFormRow) v1.FormRowDetailResponse {
	return v1.FormRowDetailResponse{
		CreateTime:    *TimeToString(req.CreateTime),
		Data:          req.Data,
		Id:            req.ID,
		Remark:        req.Remark,
		SignOnBehalf:  req.SignOnBehalf,
		SignaturePath: req.SignaturePath,
		SignerId:      req.SignerID,
		SignerName:    req.SignerName,
		Status:        req.Status,
		TbFormId:      req.TbFormID,
		TbRelateId:    req.TbRelateID,
		Title:         req.Title,
		UpdateTime:    *TimeToString(req.UpdateTime),
	}
}

func GenNewRequestId() string {
	userId := strings.ReplaceAll(uuid.NewString(), "-", "")
	return strings.ToLower(userId)
}

func (s *FormRowService) FormRowCreate(ctx context.Context, apiReq *v1.FormRowCreateRequest) (*v1.FormRowCreateResponse, error) {
	contextHeader, err := NewSBasisService().ParseContextHeader(ctx)
	if err != nil {
		return nil, err
	}

	// 查询表单模板信息
	formId := apiReq.TbFormId
	formInfo, err := s.Form.FormInfoPreloadById(ctx, formId, "")
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("form '%s' not exists", apiReq.TbFormId)
		}
		return nil, err
	}
	log.Println(formInfo.TbForm.Name)
	// 查询表单行工作项信息
	//workItemPreloadInfo, err := s.WorkItem.WorkItemPreloadByRelateId(ctx, apiReq.TbRelateId)
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		return nil, fmt.Errorf("workitem preload by relate '%s' failed,resource not exits", apiReq.TbRelateId)
	//	}
	//	return nil, err
	//}
	// 校验表单工作项预计开始时间：当前时间与预计开始时间进行比对,当预计开始时间晚于当前时间，禁止提交
	//currTime := time.Now()
	//planStartTime := workItemPreloadInfo.TbWorkItem.PlanStartTime
	//if nil != planStartTime && !planStartTime.IsZero() && planStartTime.After(currTime) {
	//	return nil, fmt.Errorf("预计开始时间:'%s'还未到，无法提交", *TimeToString(planStartTime))
	//}

	// 组装FormRow写入参数
	tbFormRow := ConvertFormRowCreateReqToTbFormRow(apiReq)
	tbFormRow.Submitter = contextHeader.TokenInfo.UserId
	// 【表单数据字段阈值预警处理】
	//tbColumnWarning, err := s.ThresholdJudgment(apiReq, tbFormRow, &formInfo, workItemPreloadInfo)
	//if err != nil {
	//	log.Println(err)
	//	return nil, err
	//}

	// 将表单提交数据&预警信息写入数据库
	newFormRow, err := s.FormRow.FormRowCreate(ctx, &tbFormRow)
	if err != nil {
		return nil, err
	}

	// 如果有预警信息，给患者和负责人发送消息推送
	//if len(tbColumnWarning) > 0 {
	//	if err := s.SendFormWarningMsgToUsers(ctx, Token.ID, workItemPreloadInfo); err != nil {
	//		return nil, err
	//	}
	//}

	return &v1.FormRowCreateResponse{
		Id: newFormRow.ID,
	}, nil
}

func (s *FormRowService) FormRowDelete(ctx context.Context, apiReq *v1.FormRowDeleteRequest) (resp *v1.FormRowDeleteResponse, err error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	deleteId, err := s.FormRow.FormRowDelete(ctx, apiReq.Id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//资源不存在按照删除成功处理
			return nil, err
		}
		return nil, err
	}
	return &v1.FormRowDeleteResponse{
		Id: deleteId,
	}, err
}

// 表单行更新
func (s *FormRowService) FormRowUpdate(ctx context.Context, apiReq *v1.FormRowUpdateRequest) (resp *v1.FormRowUpdateResponse, err error) {
	// 根据ID查询文章详情
	tbFormRow, err := s.FormRow.FormRowDetail(ctx, apiReq.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	// 赋新值
	newTbFormRow := ConvertFormRowUpdateReqToTbFormRowModel(apiReq)
	newTbFormRow.ID = tbFormRow.ID
	newTbFormRow.UpdateTime = GetNowTimeAddr()
	// 执行更新
	workItem, err := s.FormRow.FormRowUpdate(ctx, &newTbFormRow)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	// 结果赋值
	//*resp = ConvertTbFormRowToUpdateApiRespModel(workItem)

	return &v1.FormRowUpdateResponse{
		Id:         workItem.ID,
		Title:      workItem.Title,
		Status:     workItem.Status,
		Data:       workItem.Data,
		Remark:     workItem.Remark,
		TbFormId:   workItem.TbFormID,
		TbRelateId: workItem.TbRelateID,
		CreateTime: *TimeToString(workItem.CreateTime),
		UpdateTime: *TimeToString(workItem.UpdateTime),
	}, err
}

//表单行详情查询

func (s *FormRowService) FormRowDetail(ctx context.Context, apiReq *v1.FormRowDetailRequest) (resp *v1.FormRowDetailResponse, err error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	tbFormRow, err := s.FormRow.FormRowDetail(ctx, apiReq.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	//*resp = ConvertTbFormRowToDetailQueryApiRespModel(tbFormRow)
	return &v1.FormRowDetailResponse{
		Id:            tbFormRow.ID,
		Title:         tbFormRow.Title,
		Status:        tbFormRow.Status,
		Data:          tbFormRow.Data,
		Remark:        tbFormRow.Remark,
		TbFormId:      tbFormRow.TbFormID,
		TbRelateId:    tbFormRow.TbRelateID,
		SignaturePath: tbFormRow.SignaturePath,
		SignerId:      tbFormRow.SignerID,
		SignerName:    tbFormRow.SignerName,
		SignOnBehalf:  tbFormRow.SignOnBehalf,
		CreateTime:    *TimeToString(tbFormRow.CreateTime),
		UpdateTime:    *TimeToString(GetNowTimeAddr()),
	}, err
}

//表单行过滤查询

func (s *FormRowService) FormRowFilter(ctx context.Context, apiReq *v1.FormRowFilterRequest) (resp *v1.FormRowFilterResponse, err error) {
	//serverContext, ok := transport.FromServerContext(ctx)
	//if !ok {
	//	return nil, fmt.Errorf("解析Context获取TOKEN失败")
	//}
	//rv := serverContext.RequestHeader().Get("Usertoken")
	//Token, _ := utils.ParseToken(rv)
	//log.Println(Token)

	FormRowList, count, err := s.FormRow.FormRowFilter(ctx, apiReq)
	if err != nil {
		return nil, err
	}
	var formRowlistRes []*v1.Formrow

	//循环转换对象
	for _, Fr := range FormRowList {
		formRowlistRes = append(formRowlistRes,
			&v1.Formrow{
				Id:            Fr.ID,
				Data:          Fr.Data,
				Remark:        Fr.Remark,
				Status:        Fr.Status,
				TbFormId:      Fr.TbFormID,
				TbRelateId:    Fr.TbRelateID,
				Title:         Fr.Title,
				SignerName:    Fr.SignerName,
				SignOnBehalf:  Fr.SignOnBehalf,
				SignaturePath: Fr.SignaturePath,
				SignerID:      Fr.SignerID,
				Submitter:     Fr.Submitter,
				CreateTime:    *TimeToString(Fr.CreateTime),
				UpdateTime:    *TimeToString(Fr.UpdateTime),
			})
	}
	filterResponse := v1.FormRowFilterResponse{
		Page:    apiReq.Page,
		PerPage: apiReq.PerPage,
		Results: formRowlistRes,
		Total:   count,
	}
	return &filterResponse, nil
}

func (s *FormRowService) ThresholdJudgment(req *v1.FormRowCreateRequest, tbFormRow chdb.TbFormRow, formInfo *model.FormInfo, workItemPreloadInfo *model.WorkItemPreloadInfo) ([]chdb.TbFormWarning, error) {
	// 初始化响应数据
	result := make([]chdb.TbFormWarning, 0)
	// 将表单提交的json数据转为map
	reqDataMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(req.Data), &reqDataMap); err != nil {
		return nil, err
	}
	// 遍历表单字段，对字段进行预警判断
	for _, columnInfo := range formInfo.FormColumnInfos {
		tbFormColumn := columnInfo.TbFormColumn
		// 判断当前表单字段是否包含在提交的数据里面
		cValue, ok := reqDataMap[tbFormColumn.Name]
		if !ok {
			if tbFormColumn.Required {
				// 字段不存再且字段为必传
				//return nil, fmt.Errorf("bad request, column '%s' is required", tbFormColumn.Name)
				continue
			} else {
				// 字段不存再单字段为非必传
				continue
			}
		}
		// 遍历字段阈值进行判断
		columnThresholds := columnInfo.TbColumnThresholds
		if nil == columnThresholds || len(columnThresholds) == 0 {
			continue
		}
		// 遍历字段预警规则
		for _, columnThreshold := range columnThresholds {
			// 根据字段数据类型进行类型断言
			switch cValue.(type) {
			case float64:
				cFloat := cValue.(float64)

				//columnOpts := columnInfo.TbColumnOpts
				//for _, columnOpt := range columnOpts {
				//
				//}

				minFloat, err := strconv.ParseFloat(columnThreshold.Min, 64)
				if err != nil {
					return nil, err
				}
				maxFloat, err := strconv.ParseFloat(columnThreshold.Max, 64)
				if err != nil {
					return nil, err
				}
				thresholdMatched := false
				if columnThreshold.Reverse {
					// 阈值反转,判断: x < min, x > max
					if cFloat < minFloat || cFloat > maxFloat {
						thresholdMatched = true
					}
				} else {
					// 阈值未反转,判断: min <= x <= max
					if cFloat >= minFloat && cFloat <= maxFloat {
						thresholdMatched = true
					}
				}
				warningInfo := ""
				if thresholdMatched {
					// 解析预警信息规则并将提交的字段值写入预警信息,正则提取预警规则字段
					pattern := `{{(.*)}}`
					compile, err := regexp.Compile(pattern)
					if err != nil {
						return nil, err
					}
					findAll := compile.FindAll([]byte(columnThreshold.WarningInfo), -1)
					for _, matchedStr := range findAll {
						warningInfo = strings.ReplaceAll(columnThreshold.WarningInfo, string(matchedStr), fmt.Sprintf("%v", cFloat))
					}
				}

				if thresholdMatched {
					tbColumnWarning := chdb.TbFormWarning{
						ID:                  GenNewUUID(),
						PrincipalType:       workItemPreloadInfo.TbWorkItem.PrincipalType,
						PrincipalID:         workItemPreloadInfo.TbWorkItem.PrincipalID,
						PrincipalName:       workItemPreloadInfo.TbWorkItem.PrincipalName,
						AssignedType:        workItemPreloadInfo.TbWorkItem.AssignedType,
						AssignedTo:          workItemPreloadInfo.TbWorkItem.AssignedTo,
						AssignedToName:      workItemPreloadInfo.TbWorkItem.AssignedToName,
						FieldValue:          cFloat,
						WarningInfo:         warningInfo,
						TbFormRowID:         tbFormRow.ID, // 此处组装的数据错误
						TbColumnThresholdID: columnThreshold.ID,
						TbFormColumnID:      tbFormColumn.ID, // 此处组装的数据错误
						CreateTime:          nil,
						UpdateTime:          nil,
					}
					result = append(result, tbColumnWarning)
				}

			default:

			}

		}

	}

	return result, nil
}

func (s *FormRowService) SendFormWarningMsgToUsers(ctx context.Context, commitUserId string, workItemPreloadInfo *model.WorkItemPreloadInfo) error {
	// 根据工作项负责成员ID查询用户ID
	tbMember, err := s.Member.MemberDetailQueryById(ctx, workItemPreloadInfo.TbWorkItem.PrincipalID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("member '%s' not exists", err.Error())
		}
		return err
	}

	toSendUserIds := make([]string, 0)
	toSendUserIds = append(toSendUserIds, commitUserId)      // 向表单提交者发送表单预警推送消息
	toSendUserIds = append(toSendUserIds, tbMember.TbUserID) // 向工作项负责人发送表单预警推送消息

	// 执行消息发送
	msgStructure := msgcenter.MsgStructure{
		Type:      msgcenter.MsgTypeBackgroundMessage,
		Payload:   nil,
		Time:      time.Now().Unix(),
		RequestId: GenNewRequestId(),
	}
	bytes, err := json.Marshal(&msgStructure)
	if err != nil {
		return err
	}
	msgcenter.GetEProMsgService().Emits(msgcenter.EventMsgFormWarning, toSendUserIds, string(bytes))

	return nil
}
