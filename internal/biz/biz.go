package biz

import (
	"github.com/google/wire"
	"github.com/minio/minio-go/v6"
	"gorm.io/gorm"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewPlanUsecase, NewArticleUsecase,
	NewAuthUsecase, NewDiseaseUseCase,
	NewDisFunctionUseCase, NewFormUseCase,
	NewFilesUsecase, NewFormCssUsecase,
	NewWorkItemUsecase, NewFollowupUseCase,
	NewUserUsecase, NewFormRowUseCase, NewMemberUseCase)

type DataRepo interface {
	GetDB() *gorm.DB
	GetMinion() *minio.Client
}
