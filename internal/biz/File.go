package biz

import (
	"context"
	"followup/gencode/chdb"

	"github.com/go-kratos/kratos/v2/log"
	"time"
)

// TbFile 文件表
type TbFile struct {
	ID              string
	Type            string
	Name            string
	Pid             string
	StoragePlatform string
	FileURL         string
	Size            int
	ServiceModule   string
	Sha256          string
	Suffix          string
	Uploader        string
	CreateTime      *time.Time
	UpdateTime      *time.Time
}

// 定义 File 的操作接口
type FilesRepo interface {
	FileDetailQueryById(ctx context.Context, id string) (*chdb.TbFile, error)
}

type FilesUsecase struct {
	repo FilesRepo
	log  *log.Helper
}

func NewFilesUsecase(repo FilesRepo, logger log.Logger) *FilesUsecase {
	return &FilesUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *FilesUsecase) FileDetailQueryById(ctx context.Context, id string) (*chdb.TbFile, error) {
	uc.log.WithContext(ctx).Infof("FileDetailQueryById")
	return uc.repo.FileDetailQueryById(ctx, id)
}
