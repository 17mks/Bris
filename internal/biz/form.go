package biz

import (
	"context"
	"errors"
	v1 "followup/api"
	"followup/gencode/chdb"
	"followup/model"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v6"
	"gogs.buffalo-robot.com/gogs/module/lib"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// TbForm 病种表
type TbForm struct {
	ID          string
	Name        string
	Type        string
	Status      string
	BranchLogic bool
	BelongType  string
	BelongTo    string
	Description string
	AppID       string
	VersionName string
	VersionCode int32
	CreateTime  time.Time
	UpdateTime  time.Time
	DeleteAt    time.Time
}

// 定义Form 的操作接口
type FormRepo interface {
	Save(ctx context.Context, form *chdb.TbForm) error
	Delete(context.Context, string) (string, error)
	Update(context.Context, *chdb.TbForm) (*chdb.TbForm, error)
	GetByID(context.Context, string) (*chdb.TbForm, error)
	Filter(ctx context.Context, req *v1.FormFilterRequest) ([]model.FormDetailPreloadInfo, int64, error)

	Import(ctx context.Context, formInfo *model.FormInfo) (*model.FormInfo, error)
	FormInfoPreloadByIds(ctx context.Context, ids []string, projectId string) ([]model.FormInfo, error)
	FormPseudoDelById(ctx context.Context, id string) (string, error) // 伪删除
	FormDetailPreloadById(ctx context.Context, id string) (*model.FormDetailPreloadInfo, error)
	FormInfoPreloadById(ctx context.Context, id string, projectId string) (*model.FormInfo, error)

	QueryFormByType(ctx context.Context, fType string) ([]chdb.TbForm, error)     // 根据表单类型查询表单
	QueryLatestFormByName(ctx context.Context, name string) (*chdb.TbForm, error) // 根据表名称查询最新的表单

	//FormModelUpload(ctx context.Context, apiReq *v1.FormModelUploadApiReq) (protocol.StatusCode, error)
	FormQueryByIds(ctx context.Context, ids []string) ([]chdb.TbForm, error)
}

type FormUseCase struct {
	repo FormRepo
	//path     conf.SerachPath
	dataRepo DataRepo
	log      *log.Helper
}

func NewFormUseCase(repo FormRepo, dataRepo DataRepo, logger log.Logger) *FormUseCase {
	return &FormUseCase{
		repo: repo,
		//path:     path,
		dataRepo: dataRepo,
		log:      log.NewHelper(logger),
	}
}
func GenNewUUID() string {
	userId := strings.ReplaceAll(uuid.NewString(), "-", "")
	return strings.ToLower(userId)
}

func (uc *FormUseCase) CreateForm(ctx context.Context, Form *chdb.TbForm) (*chdb.TbForm, error) {
	df := &chdb.TbForm{
		ID:          Form.ID,
		Name:        Form.Name,
		Status:      Form.Status,
		Type:        Form.Type,
		BelongTo:    Form.BelongTo,
		BelongType:  Form.BelongType,
		BranchLogic: Form.BranchLogic,
		VersionCode: Form.VersionCode,
		VersionName: Form.VersionName,
		Description: Form.Description,
		CreateTime:  Form.CreateTime,
		UpdateTime:  Form.UpdateTime,
		DeleteAt:    Form.DeleteAt,
	}
	if err := uc.repo.Save(ctx, df); err != nil {
		return nil, err
	}
	return &chdb.TbForm{
		ID: Form.ID,
	}, nil
}

func (uc *FormUseCase) DeleteForm(ctx context.Context, id string) (string, error) {
	uc.log.WithContext(ctx).Infof("Delete: %v", id)
	return uc.repo.Delete(ctx, id)
}

func (uc *FormUseCase) UpdateForm(ctx context.Context, Form *chdb.TbForm) (*chdb.TbForm, error) {
	uc.log.WithContext(ctx).Infof("Update: %v", Form.ID)
	return uc.repo.Update(ctx, Form)
}

func (uc *FormUseCase) DetailForm(ctx context.Context, id string) (*chdb.TbForm, error) {
	uc.log.WithContext(ctx).Infof("Get: %v", id)
	return uc.repo.GetByID(ctx, id)
}

func (uc *FormUseCase) FilterForm(ctx context.Context, req *v1.FormFilterRequest) ([]model.FormDetailPreloadInfo, int64, error) {
	uc.log.WithContext(ctx).Infof("filter")
	return uc.repo.Filter(ctx, req)
}

func (uc *FormUseCase) ImportForm(ctx context.Context, formInfo *model.FormInfo) (*model.FormInfo, error) {
	uc.log.WithContext(ctx).Infof("Import")
	return uc.repo.Import(ctx, formInfo)
}

func (uc *FormUseCase) FormInfoPreloadByIds(ctx context.Context, ids []string, projectId string) ([]model.FormInfo, error) {
	uc.log.WithContext(ctx).Infof("Export")
	return uc.repo.FormInfoPreloadByIds(ctx, ids, projectId)
}

func (uc *FormUseCase) FormPseudoDelById(ctx context.Context, id string) (string, error) {
	uc.log.WithContext(ctx).Infof("FormPseudoDelById")
	return uc.repo.FormPseudoDelById(ctx, id)
}

func (uc *FormUseCase) FormDetailPreloadById(ctx context.Context, id string) (*model.FormDetailPreloadInfo, error) {
	uc.log.WithContext(ctx).Infof("FormDetailPreloadById")
	return uc.repo.FormDetailPreloadById(ctx, id)
}

func (uc *FormUseCase) FormInfoPreloadById(ctx context.Context, id string, projectId string) (*model.FormInfo, error) {
	uc.log.WithContext(ctx).Infof("FormInfoPreloadById")
	return uc.repo.FormInfoPreloadById(ctx, id, projectId)
}

func (uc *FormUseCase) QueryFormByType(ctx context.Context, fType string) ([]chdb.TbForm, error) {
	uc.log.WithContext(ctx).Infof("QueryFormByType")
	return uc.repo.QueryFormByType(ctx, fType)
}

func (uc *FormUseCase) QueryLatestFormByName(ctx context.Context, name string) (*chdb.TbForm, error) {
	uc.log.WithContext(ctx).Infof("QueryLatestFormByName")
	return uc.repo.QueryLatestFormByName(ctx, name)
}

func (uc *FormUseCase) FormQueryByIds(ctx context.Context, ids []string) ([]chdb.TbForm, error) {
	uc.log.WithContext(ctx).Infof("FormQueryByIds")
	return uc.repo.FormQueryByIds(ctx, ids)
}

func (uc *FormUseCase) OSSFileDownLoader(ctx context.Context, fileName string) (string, error) {

	file, err := uc.dataRepo.GetMinion().GetObject("bfr", fileName, minio.GetObjectOptions{})

	if err != nil {
		uc.log.WithContext(ctx).Error(err)
		return "", err
	}
	suffix := path.Ext(fileName)
	Name := GenNewUUID() + "." + suffix
	localFile, err := os.Create(filepath.Join("F:\\Kratos\\followup\\files", Name))

	if err != nil {
		uc.log.WithContext(ctx).Error(err)
		return "", errors.New("文件创建错误")
	}
	defer localFile.Close()

	stat, err := file.Stat()
	if err != nil {
		uc.log.WithContext(ctx).Error(err)
		return "", err
	}

	if _, err := io.CopyN(localFile, file, stat.Size); err != nil {
		//uc.log
		return "", err
	}
	return "", nil

}

func (uc *FormUseCase) ParseFile(ctx context.Context, fileName string) error {
	if fileName == "" {
		return errors.New("文件路径不存在")
	}
	//解压文件
	filePath := filepath.Join("F:\\Kratos\\followup\\files", fileName)
	suffix := path.Ext(fileName)

	fileprefix := fileName[0 : len(fileName)-len(suffix)]
	err := lib.Unzip(filePath, filepath.Join("F:\\Kratos\\followup\\files", fileprefix))

	if err != nil {
		uc.log.WithContext(ctx).Error(err)
		return errors.New("文件解压错误")
	}

	//读取setting
	//settingPath := filePath
	return nil
}
