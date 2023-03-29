package biz

import (
	"context"
	"followup/gencode/chdb"
	"github.com/go-kratos/kratos/v2/log"
)

// 定义Member 的操作接口
type MemberRepo interface {
	MemberDetailQueryById(ctx context.Context, id string) (*chdb.TbMember, error)
}

type MemberUseCase struct {
	repo MemberRepo
	log  *log.Helper
}

func NewMemberUseCase(repo MemberRepo, logger log.Logger) *MemberUseCase {
	return &MemberUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *MemberUseCase) MemberDetailQueryById(ctx context.Context, id string) (*chdb.TbMember, error) {
	uc.log.WithContext(ctx).Infof("MemberDetailQueryById")
	return uc.repo.MemberDetailQueryById(ctx, id)
}
