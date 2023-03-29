package data

import (
	"followup/internal/biz"
	"followup/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/minio/minio-go/v6"
	"gogs.buffalo-robot.com/services/basis/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGormDB,
	NewFileData,
	NewPlanRepo, NewArticleRepo,
	NewAuthRepo, NewDiseaseRepo,
	NewDisFunctionRepo, NewFormRepo,
	NewFilesRepo, NewFormCssRepo,
	NewWorkItemRepo, NewFollowupRepo,
	NewUserRepo, NewFormRowRepo,
	NewMinio, NewMemberRepo,
)

// Data .
type Data struct {
	gormDB *gorm.DB
	logger log.Logger
}

type DataRepo struct {
	gormDB *gorm.DB
	minio  *minio.Client
	logger log.Logger
}

// NewData .
func NewData(logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{gormDB: db}, cleanup, nil

}

func NewFileData(logger log.Logger, gorm *gorm.DB, minioClient *minio.Client) biz.DataRepo {
	return &DataRepo{
		gormDB: gorm,
		minio:  minioClient,
		logger: logger,
	}
}

func (d *DataRepo) GetMinion() *minio.Client {
	return d.minio
}

func (d *DataRepo) GetDB() *gorm.DB {
	return d.gormDB
}

func NewMinio(c *conf.Data) *minio.Client {
	client, err := minio.New(c.Minion.Addr, c.Minion.AccessKeyID, c.Minion.SecretAccessKey, false)
	if err != nil {
		panic(err)
	}
	buckets := []string{models.ConsultFileBucket, models.BaseBucket}
	for _, bucket := range buckets {
		ifExist, err := client.BucketExists(bucket)
		if err != nil {
			panic(err)
		}
		if !ifExist {
			client.MakeBucket(bucket, "")
			// if err != nil {
			// 	panic(err)
			// }
		}
	}
	return client
}

func NewGormDB(c *conf.Data) (*gorm.DB, error) {
	dsn := c.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(150)
	sqlDB.SetConnMaxLifetime(time.Second * 25)
	InitDB(db)
	return db, err
}

func InitDB(db *gorm.DB) {
	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}
}
