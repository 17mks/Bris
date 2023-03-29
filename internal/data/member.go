package data

import (
	"context"
	"followup/gencode/chdb"
	"followup/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type MemberRepo struct {
	data *Data
	log  *log.Helper
}

func NewMemberRepo(data *Data, logger log.Logger) biz.MemberRepo {
	return &MemberRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *MemberRepo) MemberDetailQueryById(ctx context.Context, id string) (*chdb.TbMember, error) {
	tbMember := chdb.TbMember{}
	if err := r.data.gormDB.Where("id = ?", id).First(&tbMember).Error; err != nil {
		return nil, err
	}
	return &tbMember, nil
}
