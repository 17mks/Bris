package service

import (
	v1 "followup/api"
	"followup/internal/biz"
	"followup/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/minio/minio-go/v6"
	"io"
	"os"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPlanService, NewAuthTokenService,
	NewDiseaseService, NewArticleService,
	NewDisFunctionService, NewFormService,
	NewFollowupService, NewFormRowService,
	NewFileService,
)

// AuthTokenService
type AuthTokenService struct {
	v1.UnimplementedAuthTokenServer
	Auth *biz.AuthUsecase
	log  *log.Helper
}

func NewAuthTokenService(auth *biz.AuthUsecase, logger log.Logger) *AuthTokenService {
	return &AuthTokenService{
		Auth: auth,
		log:  log.NewHelper(logger),
	}
}

// DiseaseService
type DiseaseService struct {
	v1.UnimplementedDiseaseServer
	Disease *biz.DiseaseUseCase
	log     *log.Helper
}

func NewDiseaseService(disease *biz.DiseaseUseCase, logger log.Logger) *DiseaseService {
	return &DiseaseService{
		Disease: disease,
		log:     log.NewHelper(logger),
	}
}

// ArticleService
type ArticleService struct {
	v1.UnimplementedArticleServer

	Article *biz.ArticleUsecase
	log     *log.Helper
}

func NewArticleService(article *biz.ArticleUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{
		Article: article,
		log:     log.NewHelper(logger),
	}
}

// DisFunctionService
type DisFunctionService struct {
	v1.UnimplementedDisFunctionServer
	DisFunction *biz.DisFunctionUseCase
	Log         *log.Helper
}

func NewDisFunctionService(disFunction *biz.DisFunctionUseCase, logger log.Logger) *DisFunctionService {
	return &DisFunctionService{
		DisFunction: disFunction,
		Log:         log.NewHelper(logger),
	}
}

type FileService struct {
	saver FileSaver
	log   log.Logger
}

func NewFileService(data biz.DataRepo, log log.Logger) *FileService {
	service := FileService{}
	service.saver = NewFileSaver(data)
	service.log = log
	return &service
}

// FormService 表单服务
type FormService struct {
	v1.UnimplementedFormServer
	Form     *biz.FormUseCase
	FormCss  *biz.FormCssUsecase
	File     *biz.FilesUsecase
	log      *log.Helper
	confData *conf.Data
}

// DownLoadFile 将文件下载到target路径
func (service *FormService) DownLoadFile(bucketName string, fileObjName string, target string) error {

	// 从配置文件读取minio配置信息
	endPoint := service.confData.Minion.Addr
	user := service.confData.Minion.AccessKeyID
	password := service.confData.Minion.SecretAccessKey

	client, err := minio.New(endPoint, user, password, false)
	if err != nil {
		panic(err)
	}
	object, err := client.GetObject(bucketName, fileObjName, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	localFile, err := os.Create(target)

	if err != nil {
		return err
	}
	defer func(localFile *os.File) {
		_ = localFile.Close()
	}(localFile)

	stat, err := object.Stat()
	if err != nil {
		return err
	}

	if _, err := io.CopyN(localFile, object, stat.Size); err != nil {
		return err
	}
	return nil
}

func NewFormService(form *biz.FormUseCase, file *biz.FilesUsecase, formCss *biz.FormCssUsecase, logger log.Logger, data *conf.Data) *FormService {
	return &FormService{
		Form:     form,
		FormCss:  formCss,
		File:     file,
		log:      log.NewHelper(logger),
		confData: data,
	}
}

// FollowupService
type FollowupService struct {
	v1.UnimplementedFollowupServer
	Followup    *biz.FollowupUseCase
	User        *biz.UserUsecase
	planUseCase *biz.PlanUsecase
	WorkItem    *biz.WorkItemUsecase
	log         *log.Helper
}

func (s *FollowupService) checkFollowupCreateRequest(request *v1.FollowupCreateRequest) error {
	return nil
}

func NewFollowupService(followup *biz.FollowupUseCase, user *biz.UserUsecase, plan *biz.PlanUsecase, workItem *biz.WorkItemUsecase, logger log.Logger) *FollowupService {
	return &FollowupService{
		Followup:    followup,
		User:        user,
		planUseCase: plan,
		WorkItem:    workItem,
		log:         log.NewHelper(logger),
	}
}

type FormRowService struct {
	v1.UnimplementedFormRowServiceServer
	Form     *biz.FormUseCase
	FormRow  *biz.FormRowUseCase
	WorkItem *biz.WorkItemUsecase
	Member   *biz.MemberUseCase
	Log      *log.Helper
}

func NewFormRowService(formRow *biz.FormRowUseCase, form *biz.FormUseCase, workItem *biz.WorkItemUsecase, member *biz.MemberUseCase, logger log.Logger) *FormRowService {
	return &FormRowService{
		FormRow:  formRow,
		Form:     form,
		WorkItem: workItem,
		Member:   member,
		Log:      log.NewHelper(logger),
	}
}
