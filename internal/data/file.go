package data

import (
	"context"
	"followup/gencode/chdb"
	"followup/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type FilesRepo struct {
	data *Data
	log  *log.Helper
}

func NewFilesRepo(data *Data, logger log.Logger) biz.FilesRepo {
	return &FilesRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *FilesRepo) FileDetailQueryById(ctx context.Context, id string) (*chdb.TbFile, error) {
	task := chdb.TbFile{ID: id}
	if err := r.data.gormDB.First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
