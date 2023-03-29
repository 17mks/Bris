package service

import (
	"archive/zip"
	"context"
	"encoding/json"
	"fmt"
	v1 "followup/api"
	"followup/gencode/chdb"
	"followup/model"
	"followup/utils"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"
	"gogs.buffalo-robot.com/gogs/module/lib"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Unzip(zipPath, dstDir string) error {
	// open zip file
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer func(reader *zip.ReadCloser) {
		_ = reader.Close()
	}(reader)
	for _, file := range reader.File {
		if err := unzipFile(file, dstDir); err != nil {
			return err
		}
	}
	return nil
}

func unzipFile(file *zip.File, dstDir string) error {
	// create the directory of file
	filePath := path.Join(dstDir, file.Name)
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// open the file
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer func(rc io.ReadCloser) {
		_ = rc.Close()
	}(rc)

	// create the file
	w, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(w *os.File) {
		_ = w.Close()
	}(w)

	// save the decompressed file content
	_, err = io.Copy(w, rc)
	return err
}

// ConvertFormImportApiReqToFormInfo 将表单模板导入接口参数转换为DBModel
func ConvertFormImportApiReqToFormInfo(apiReq *v1.FormImportRequest) model.FormInfo {
	// 将请求参数转换为DBModel
	var tbForm chdb.TbForm
	if nil != apiReq.Form {
		tbForm = chdb.TbForm{
			ID:          GenNewUUID(),
			Name:        apiReq.Form.Name,
			Type:        apiReq.Form.Type,
			Status:      apiReq.Form.Status,
			BranchLogic: apiReq.Form.BranchLogic,
			BelongType:  apiReq.Form.BelongType,
			BelongTo:    apiReq.Form.BelongTo,
			Description: apiReq.Form.Description,
			AppID:       "",
			VersionName: "",
			VersionCode: 0,
			CreateTime:  nil,
			UpdateTime:  nil,
			DeleteAt:    nil,
		}
	}
	var tbFormCSS chdb.TbFormCSS
	if nil != apiReq.FormCss {
		tbFormCSS = chdb.TbFormCSS{
			ID:          GenNewUUID(),
			Name:        apiReq.FormCss.Name,
			Status:      apiReq.FormCss.Status,
			CSSCode:     apiReq.FormCss.CssCode,
			Description: apiReq.FormCss.Description,
			CreateTime:  nil,
			UpdateTime:  nil,
			TbFormID:    tbForm.ID,
		}
	}

	formColumnInfos := make([]model.FormColumnInfo, 0)
	if nil != apiReq.FormColumns {
		for _, formColumn := range apiReq.FormColumns {
			formColumnInfo := model.FormColumnInfo{
				TbFormColumn: chdb.TbFormColumn{
					ID:             GenNewUUID(),
					Name:           formColumn.Column.Name,
					GroupName:      formColumn.Column.GroupName,
					DataType:       formColumn.Column.DataType,
					ViewType:       formColumn.Column.ViewType,
					AvailableValue: formColumn.Column.AvailableValue,
					DefaultValue:   formColumn.Column.DefaultValue,
					Regexp:         formColumn.Column.Regexp,
					SortIndex:      formColumn.Column.SortIndex,
					Description:    formColumn.Column.Description,
					CreateTime:     nil,
					UpdateTime:     nil,
					TbFormID:       tbForm.ID,
				},
				TbColumnOpts:       make([]chdb.TbColumnOpt, 0),
				TbColumnThresholds: make([]chdb.TbColumnThreshold, 0),
			}

			for _, columnOpt := range formColumn.ColumnOpts {
				tbColumnOpt := chdb.TbColumnOpt{
					ID:             GenNewUUID(),
					SortIndex:      columnOpt.SortIndex,
					Value:          columnOpt.Value,
					NextFiledID:    columnOpt.NextFiledId,
					CreateTime:     nil,
					UpdateTime:     nil,
					TbFormColumnID: formColumnInfo.TbFormColumn.ID,
				}
				formColumnInfo.TbColumnOpts = append(formColumnInfo.TbColumnOpts, tbColumnOpt)
			}
			for _, columnThreshold := range formColumn.ColumnThresholds {
				tbColumnThreshold := chdb.TbColumnThreshold{
					ID:             GenNewUUID(),
					Min:            columnThreshold.Min,
					Max:            columnThreshold.Max,
					Reverse:        columnThreshold.Reverse,
					Level:          columnThreshold.Level,
					WarningInfo:    columnThreshold.WarningInfo,
					CreateTime:     nil,
					UpdateTime:     nil,
					TbFormColumnID: formColumnInfo.TbFormColumn.ID,
				}
				formColumnInfo.TbColumnThresholds = append(formColumnInfo.TbColumnThresholds, tbColumnThreshold)
			}
			formColumnInfos = append(formColumnInfos, formColumnInfo)
		}
	}

	formInfo := model.FormInfo{}
	formInfo.TbForm = tbForm
	formInfo.TbFormCSS = tbFormCSS
	formInfo.FormColumnInfos = formColumnInfos
	return formInfo
}

// ConvertFormInfoToDetailQueryResp 将表单信息转换为接口响应数据
func ConvertFormInfoToDetailQueryResp(formInfo *model.FormInfo) *v1.FormDetailResponse {
	// 将请求参数转换为DBModel
	tbForm := v1.FormBean{
		BelongTo:    formInfo.TbForm.BelongTo,
		BelongType:  formInfo.TbForm.BelongType,
		BranchLogic: formInfo.TbForm.BranchLogic,
		Description: formInfo.TbForm.Description,
		Name:        formInfo.TbForm.Name,
		Status:      formInfo.TbForm.Status,
		Type:        formInfo.TbForm.Type,
		VersionCode: formInfo.TbForm.VersionCode,
		VersionName: formInfo.TbForm.VersionName,
		Id:          formInfo.TbForm.ID,
		CreateTime:  *TimeToString(formInfo.TbForm.CreateTime),
		UpdateTime:  *TimeToString(formInfo.TbForm.UpdateTime),
		DeleteAt:    "",
	}
	tbFormCSS := v1.FormCssBean{
		CreateTime:  *TimeToString(formInfo.TbFormCSS.CreateTime),
		CssCode:     formInfo.TbFormCSS.CSSCode,
		CssUrl:      formInfo.TbFormCSS.CSSURL,
		VersionCode: formInfo.TbFormCSS.VersionCode,
		VersionName: formInfo.TbFormCSS.VersionName,
		Description: formInfo.TbFormCSS.Description,
		Id:          formInfo.TbFormCSS.ID,
		Name:        formInfo.TbFormCSS.Name,
		Status:      formInfo.TbFormCSS.Status,
		TbFormId:    formInfo.TbFormCSS.TbFormID,
		UpdateTime:  *TimeToString(formInfo.TbFormCSS.UpdateTime),
	}
	formColumnInfos := make([]*v1.FormDetailResponse_Formcolumn, 0)
	for _, columnInfo := range formInfo.FormColumnInfos {
		formColumnInfo := v1.FormDetailResponse_Formcolumn{
			Column: &v1.Formcolumns{
				AvailableValue: columnInfo.TbFormColumn.AvailableValue,
				CreateTime:     *TimeToString(columnInfo.TbFormColumn.CreateTime),
				DataType:       columnInfo.TbFormColumn.DataType,
				DefaultValue:   columnInfo.TbFormColumn.DefaultValue,
				Description:    columnInfo.TbFormColumn.Description,
				GroupName:      columnInfo.TbFormColumn.GroupName,
				Id:             columnInfo.TbFormColumn.ID,
				Name:           columnInfo.TbFormColumn.Name,
				Regexp:         columnInfo.TbFormColumn.Regexp,
				SortIndex:      int32(columnInfo.TbFormColumn.SortIndex),
				TbFormId:       columnInfo.TbFormColumn.TbFormID,
				UpdateTime:     *TimeToString(columnInfo.TbFormColumn.UpdateTime),
				ViewType:       columnInfo.TbFormColumn.ViewType,
				Required:       columnInfo.TbFormColumn.Required,
				Visible:        columnInfo.TbFormColumn.Visible,
				Editable:       columnInfo.TbFormColumn.Editable,
				DeletedAt:      "",
			},
			ColumnOpts:       nil,
			ColumnThresholds: nil,
		}

		for _, columnOpt := range columnInfo.TbColumnOpts {
			tbColumnOpt := v1.Columnopts{
				CreateTime:     *TimeToString(columnOpt.CreateTime),
				Id:             columnOpt.ID,
				NextFiledId:    columnOpt.NextFiledID,
				Score:          int32(columnOpt.Score),
				SortIndex:      int32(columnOpt.SortIndex),
				TbFormColumnId: columnOpt.TbFormColumnID,
				UpdateTime:     *TimeToString(columnOpt.UpdateTime),
				Value:          columnOpt.Value,
			}
			formColumnInfo.ColumnOpts = append(formColumnInfo.ColumnOpts, &tbColumnOpt)
		}
		for _, columnThreshold := range columnInfo.TbColumnThresholds {
			threshold := v1.Columnthresholds{
				CreateTime:     *TimeToString(columnThreshold.CreateTime),
				Id:             columnThreshold.ID,
				Level:          int32(columnThreshold.Level),
				Max:            columnThreshold.Max,
				Min:            columnThreshold.Min,
				Reverse:        columnThreshold.Reverse,
				TbFormColumnId: columnThreshold.TbFormColumnID,
				UpdateTime:     *TimeToString(columnThreshold.UpdateTime),
				WarningInfo:    columnThreshold.WarningInfo,
				WarningRegex:   "",
			}
			formColumnInfo.ColumnThresholds = append(formColumnInfo.ColumnThresholds, &threshold)
		}
		formColumnInfos = append(formColumnInfos, &formColumnInfo)
	}
	formExportResp := v1.FormDetailResponse{}
	formExportResp.Form = &tbForm
	formExportResp.FormCss = &tbFormCSS
	formExportResp.FormColumns = formColumnInfos
	return &formExportResp
}

//
//func ConvertTbFormToGenModel(req *v1.FormDetailPreloadInfo) *v1.FormFilterResponse_Results {
//	return &v1.FormFilterResponse_Results{
//		Form: &v1.Forms{
//			BelongTo:    req.Form.BelongTo,
//			BelongType:  req.Form.BelongType,
//			BranchLogic: req.Form.BranchLogic,
//			CreateTime:  req.Form.CreateTime,
//			Description: req.Form.Description,
//			Id:          req.Form.Id,
//			Name:        req.Form.Name,
//			Status:      req.Form.Status,
//			Type:        req.Form.Type,
//			UpdateTime:  req.Form.UpdateTime,
//			VersionCode: int32(req.Form.VersionCode),
//			VersionName: req.Form.VersionName,
//		},
//		FormCss: &v1.Formcss{
//			CreateTime:  req.FormCss.CreateTime,
//			CssCode:     req.FormCss.CssCode,
//			CssUrl:      req.FormCss.CssUrl,
//			VersionCode: int32(req.FormCss.VersionCode),
//			VersionName: req.FormCss.VersionName,
//			Description: req.FormCss.Description,
//			Id:          req.FormCss.Id,
//			Name:        req.FormCss.Name,
//			Status:      req.FormCss.Status,
//			TbFormId:    req.FormCss.TbFormId,
//			UpdateTime:  req.FormCss.UpdateTime,
//		},
//	}
//}

func (service *FormService) FormCreate(ctx context.Context, req *v1.FormCreateRequest) (*v1.FormCreateResponse, error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	at, err := service.Form.CreateForm(ctx, &chdb.TbForm{})
	if err != nil {
		return nil, err
	}
	return &v1.FormCreateResponse{
		Id: at.ID,
	}, nil
}

func (service *FormService) FormDelete(ctx context.Context, req *v1.FormDeleteRequest) (*v1.FormDeleteResponse, error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	deletedId, err := service.Form.DeleteForm(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.FormDeleteResponse{
		Id: deletedId,
	}, nil
}

func (service *FormService) FormUpdate(ctx context.Context, req *v1.FormUpdateRequest) (res *v1.FormUpdateResponse, errs error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	form, err := service.Form.UpdateForm(ctx, &chdb.TbForm{
		ID:          req.Id,
		Name:        req.Body.Name,
		Description: req.Body.Description,
		Status:      req.Body.Status,
		Type:        req.Body.Type,
		BelongTo:    req.Body.BelongTo,
		BelongType:  req.Body.BelongType,
		BranchLogic: req.Body.BranchLogic,
		VersionCode: req.Body.VersionCode,
		VersionName: req.Body.VersionName,
	})
	if err != nil {
		return nil, err
	}
	return &v1.FormUpdateResponse{
		Id:          form.ID,
		Name:        form.Name,
		Status:      form.Status,
		Type:        form.Type,
		BelongTo:    form.BelongTo,
		BelongType:  form.BelongType,
		BranchLogic: form.BranchLogic,
		Description: form.Description,
		CreateTime:  *TimeToString(GetNowTimeAddr()),
		UpdateTime:  *TimeToString(GetNowTimeAddr()),
	}, nil
}

func (service *FormService) FormDetail(ctx context.Context, apiReq *v1.FormDetailRequest) (*v1.FormDetailResponse, error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	preloadInfo, err := service.Form.FormInfoPreloadById(ctx, apiReq.Id, "")
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound("Form", "not found")
		}
		return nil, err
	}
	resp := ConvertFormInfoToDetailQueryResp(preloadInfo)

	return resp, nil
}

func (service *FormService) FormFilter(ctx context.Context, apiReq *v1.FormFilterRequest) (resp *v1.FormFilterResponse, err error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	FormList, count, err := service.Form.FilterForm(ctx, apiReq)
	if err != nil {
		return nil, err
	}
	var FormListRes []*v1.FormDetailPreloadInfo
	for _, at := range FormList {
		FormListRes = append(FormListRes,
			&v1.FormDetailPreloadInfo{
				Form: &v1.FormBean{
					BelongTo:    at.TbForm.BelongTo,
					BelongType:  at.TbForm.BelongType,
					BranchLogic: at.TbForm.BranchLogic,
					Description: at.TbForm.Description,
					Name:        at.TbForm.Name,
					Status:      at.TbForm.Status,
					Type:        at.TbForm.Type,
					VersionCode: at.TbForm.VersionCode,
					VersionName: at.TbForm.VersionName,
					Id:          at.TbForm.ID,
					CreateTime:  *TimeToString(at.TbForm.CreateTime),
					UpdateTime:  *TimeToString(at.TbForm.UpdateTime),
					DeleteAt:    *TimeToString(at.TbForm.DeleteAt),
				},
				FormCss: &v1.FormCssBean{
					CreateTime:  *TimeToString(at.TbFormCSS.CreateTime),
					CssCode:     at.TbFormCSS.CSSCode,
					CssUrl:      at.TbFormCSS.CSSURL,
					VersionCode: at.TbFormCSS.VersionCode,
					VersionName: at.TbFormCSS.VersionName,
					Description: at.TbFormCSS.Description,
					Id:          at.TbFormCSS.ID,
					Name:        at.TbFormCSS.Name,
					Status:      at.TbFormCSS.Status,
					TbFormId:    at.TbFormCSS.TbFormID,
					UpdateTime:  *TimeToString(at.TbFormCSS.UpdateTime),
				},
			})
	}
	return &v1.FormFilterResponse{
		Page:    apiReq.Page,
		PerPage: apiReq.PerPage,
		Results: FormListRes,
		Total:   count,
	}, nil
}

func (service *FormService) FormImport(ctx context.Context, apiReq *v1.FormImportRequest) (res *v1.FormImportResponse, errs error) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		return nil, fmt.Errorf("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	log.Println(Token)

	//转换请求参数
	formInfo := ConvertFormImportApiReqToFormInfo(apiReq)
	//执行表单模板导入
	formImport, err := service.Form.ImportForm(ctx, &formInfo)
	if err != nil {
		return nil, err
	}
	//结果赋值
	res.Id = formImport.TbForm.ID
	return nil, err
}

func (service *FormService) FormModelUpload(ctx context.Context, apiReq *v1.FormModelUploadRequest) (resp *v1.FormModelUploadResponse, errs error) {
	// TODO 下载文件 & 解压文件 & 读取setting.json & 数据转换&写入数据库 & 删除下载的文件和解压后的文件夹 & 组装API返回参数
	if "" == apiReq.FileId {
		return nil, fmt.Errorf("FileId is empty")
	}
	split := strings.SplitN(apiReq.FileId, "/", 2)

	if len(split) != 2 {
		return nil, fmt.Errorf("bad file id")
	}

	// 下载文件

	target := fmt.Sprintf("%s%c%s", service.confData.FileService.UnZipPath, filepath.Separator, split[1])
	if err := service.DownLoadFile(split[0], split[1], target); err != nil {
		return nil, err
	}
	defer func(t string) {
		//	删除已下载和解压的文件
		_ = os.Remove(t)
	}(target)
	// 解压文件
	unzipFileName := strings.TrimSuffix(split[1], ".zip")
	unzipTarget := fmt.Sprintf("%s%c%s", service.confData.FileService.UnZipPath, filepath.Separator, unzipFileName)
	//unzipTarget := "D:\\tmp\\download\\" + strings.TrimSuffix(split[1], ".zip")
	if err := lib.Unzip(target, unzipTarget); err != nil {
		return nil, err
	}
	defer func(ut string) {
		//	删除已下载和解压的文件
		//_ = os.RemoveAll(ut)
	}(unzipTarget)

	// 解析json配置文件并入库
	//查询表单样式
	var formCss chdb.TbFormCSS
	var insertedFormInfo model.FormInfo
	if "" == apiReq.FormId {
		// 读取取setting.json文件内容
		settingFile := fmt.Sprintf("%s%c%s", unzipTarget, filepath.Separator, "setting.json")
		settingJsonBytes, err := os.ReadFile(settingFile)
		if err != nil {
			return nil, err
		}
		formImportReq := v1.FormImportRequest{}

		if err := json.Unmarshal(settingJsonBytes, &formImportReq); err != nil {
			return nil, err
		}
		toImportFormInfo := ConvertFormImportApiReqToFormInfo(&formImportReq)

		// TODO 校验

		toImportFormInfo.VersionName = apiReq.VersionName
		toImportFormInfo.VersionCode = int32(apiReq.VersionCode)

		formInfo, err := service.Form.ImportForm(ctx, &toImportFormInfo)
		if err != nil {
			return nil, err
		}
		formCss = formInfo.TbFormCSS
		insertedFormInfo = *formInfo
	}
	//更新表单信息
	if "" == formCss.ID {
		formCss = chdb.TbFormCSS{
			ID:          GenNewUUID(),
			Name:        insertedFormInfo.TbForm.Name,
			Status:      "ENABLED",
			CSSCode:     "",
			CSSURL:      unzipFileName,
			OssURL:      apiReq.FileId,
			Description: "",
			VersionName: insertedFormInfo.TbForm.VersionName,
			VersionCode: insertedFormInfo.TbForm.VersionCode,
			CreateTime:  nil,
			UpdateTime:  nil,
			TbFormID:    insertedFormInfo.TbForm.ID,
		}
		if _, err := service.FormCss.FormCssCreate(ctx, &formCss); err != nil {
			return nil, err
		}
	} else {
		formCss.CSSURL = unzipFileName
		formCss.OssURL = apiReq.FileId
		formCss.Name = insertedFormInfo.TbForm.Name
		formCss.VersionName = insertedFormInfo.TbForm.VersionName
		formCss.VersionCode = insertedFormInfo.TbForm.VersionCode

		if _, err := service.FormCss.UpdateFormCssById(ctx, &formCss); err != nil {
			return nil, err
		}
	}
	return nil, nil
}
