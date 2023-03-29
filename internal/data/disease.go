package data

import (
	"context"
	"database/sql"
	"fmt"
	v1 "followup/api"
	"followup/gencode/chdb"
	"followup/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

// TbDisease 病种表
type TbDisease struct {
	ID          string       `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'病种主键编码'"`                                      // 病种主键编码
	Code        string       `gorm:"unique;column:code;type:varchar(45);default:null;comment:'疾病编码(e.g. A00-B99 C00.1)'"`                  // 疾病编码(e.g. A00-B99 C00.1)
	Name        string       `gorm:"column:name;type:varchar(512);not null;comment:'疾病名称'"`                                                // 疾病名称
	NameJp      string       `gorm:"column:name_jp;type:varchar(512);not null;comment:'疾病名称首字母拼音'"`                                        // 疾病名称首字母拼音
	NameQp      string       `gorm:"column:name_qp;type:varchar(512);not null;comment:'疾病名称拼音全拼'"`                                         // 疾病名称拼音全拼
	Version     string       `gorm:"column:version;type:varchar(45);not null;comment:'版本(e.g. 'ICD10')'"`                                  // 版本(e.g. 'ICD10')
	Status      string       `gorm:"column:status;type:enum('ENABLED','DISABLED');default:null;comment:'状态('ENABLED 启用', 'DISABLED 禁用')'"` // 状态('ENABLED 启用', 'DISABLED 禁用')
	Tag         string       `gorm:"column:tag;type:varchar(512);default:null;comment:'备用标记'"`                                             // 备用标记
	Description string       `gorm:"column:description;type:varchar(512);default:null;comment:'疾病描述'"`                                     // 疾病描述
	Pid         string       `gorm:"column:pid;type:varchar(45);default:null;comment:'父级编码'"`                                              // 父级编码
	CreateTime  time.Time    `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`               // 创建时间
	UpdateTime  time.Time    `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`               // 更新时间
	DeleteAt    sql.NullTime `gorm:"column:delete_at;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'删除时间'"`                 // 删除时间
}

type DiseaseRepo struct {
	data *Data
	log  *log.Helper
}

func NewDiseaseRepo(data *Data, logger log.Logger) biz.DiseaseRepo {
	return &DiseaseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func AddLikeCharToStr(str string) string {
	likeStr := "%"
	for _, val := range str {
		likeStr = fmt.Sprintf("%s%c%c", likeStr, val, '%')
	}
	return likeStr
}

func GetNowTimeAddr() *time.Time {
	nTime := time.Now()
	return &nTime
}

// 支持什么类型都可以自己加
var timeTemplates = []string{
	"2006-01-02 15:04:05", //常规类型
	//"2006/01/02 15:04:05",
	//"2006-01-02",
	//"2006/01/02",
}

func TimeStringToGoTime(tm string) time.Time {

	for i := range timeTemplates {
		t, err := time.ParseInLocation(timeTemplates[i], tm, time.Local)
		if nil == err && !t.IsZero() {
			return t
		}
	}
	return time.Time{}
}

func (r *DiseaseRepo) Save(ctx context.Context, ds *chdb.TbDisease) error {
	disease := chdb.TbDisease{
		ID:          ds.ID,
		Code:        ds.Code,
		Name:        ds.Name,
		NameJp:      ds.NameJp,
		NameQp:      ds.NameQp,
		Version:     ds.Version,
		Status:      ds.Status,
		Tag:         ds.Tag,
		Description: ds.Description,
		Pid:         ds.Pid,
	}
	//验证是否已创建
	result := r.data.gormDB.Where(&v1.Diseases{Id: ds.ID}).First(&disease)
	if result.RowsAffected == 1 {
		return status.Errorf(codes.AlreadyExists, "该病种已存在")
	}
	disease.ID = ds.ID
	disease.Code = ds.Code
	disease.Name = ds.Name
	disease.NameJp = ds.NameJp
	disease.NameQp = ds.NameQp
	disease.Version = ds.Version
	disease.Status = ds.Status
	disease.Tag = ds.Tag
	disease.Description = ds.Description
	disease.Pid = ds.Pid
	res := r.data.gormDB.Create(&disease)
	return res.Error
}

func (r *DiseaseRepo) DeleteByID(ctx context.Context, id string) (string, error) {
	Disease := chdb.TbDisease{
		ID: id,
	}
	if err := r.data.gormDB.Delete(&Disease).Error; err != nil {
		return "", err
	}
	return Disease.ID, nil
}

func (r *DiseaseRepo) Update(ctx context.Context, req *chdb.TbDisease) (res *chdb.TbDisease, err error) {
	disease := new(chdb.TbDisease)
	err = r.data.gormDB.Where("ID", req.ID).First(disease).Error

	if err != nil {
		return nil, err
	}
	err = r.data.gormDB.Model(&disease).Updates(&chdb.TbDisease{
		ID:          req.ID,
		Code:        req.Code,
		Name:        req.Name,
		NameJp:      req.NameJp,
		NameQp:      req.NameQp,
		Version:     req.Version,
		Status:      req.Status,
		Tag:         req.Tag,
		Description: req.Description,
		Pid:         req.Pid,
		CreateTime:  req.CreateTime,
		UpdateTime:  req.UpdateTime,
	}).Error

	return &chdb.TbDisease{
		ID:          disease.ID,
		Code:        disease.Code,
		Name:        disease.Name,
		NameJp:      disease.NameJp,
		NameQp:      disease.NameQp,
		Version:     disease.Version,
		Status:      disease.Status,
		Tag:         disease.Tag,
		Description: disease.Description,
		Pid:         disease.Pid,
		CreateTime:  disease.CreateTime,
		UpdateTime:  GetNowTimeAddr(),
	}, nil
}

func (r *DiseaseRepo) GetByID(ctx context.Context, id string) (rv *chdb.TbDisease, err error) {
	disease := new(chdb.TbDisease)
	result := r.data.gormDB.Where("id = ?", id).First(disease)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("diseaseUseCase", "not found by id")
	}
	if result.Error != nil {
		return nil, err
	}
	return &chdb.TbDisease{
		ID:          disease.ID,
		Code:        disease.Code,
		Name:        disease.Name,
		NameJp:      disease.NameJp,
		NameQp:      disease.NameQp,
		Version:     disease.Version,
		Status:      disease.Status,
		Tag:         disease.Tag,
		Description: disease.Description,
		Pid:         disease.Pid,
		CreateTime:  disease.CreateTime,
		UpdateTime:  GetNowTimeAddr(),
	}, nil
}

func (r *DiseaseRepo) Filter(ctx context.Context, apiReq *v1.DiseaseFilterRequest) ([]*chdb.TbDisease, int64, error) {

	diseaselist := make([]*chdb.TbDisease, 0)
	var count int64 = 0

	filter := apiReq.Filter
	tx := r.data.gormDB.Model(&chdb.TbDisease{})

	if filter != nil {
		if "" != filter.Key {
			keyLike := AddLikeCharToStr(filter.Key)
			tx.Where("name like ? or name_qp like ?", keyLike, keyLike)
		}
		if len(filter.Ids) > 0 {
			tx.Where("id in ?", filter.Ids)
		}

		if "" != filter.Status {
			tx.Where("status = ?", filter.Status)
		}
	}

	tx.Order("update_time DESC")
	tx.Count(&count)

	result := tx.Find(&diseaselist)
	return diseaselist, count, result.Error
}

func (r *DiseaseRepo) QueryDiseasesByIds(ctx context.Context, ids []string) ([]chdb.TbDisease, error) {
	tbDiseases := make([]chdb.TbDisease, 0)
	if err := r.data.gormDB.Where("id in ?", ids).Find(&tbDiseases).Error; err != nil {
		return nil, err
	}
	return tbDiseases, nil
}

func (r *DiseaseRepo) QueryDisFunByIds(ctx context.Context, ids []string) ([]chdb.TbDisFunc, error) {
	tbDisFuncs := make([]chdb.TbDisFunc, 0)
	if err := r.data.gormDB.Where("id in ?", ids).Find(&tbDisFuncs).Error; err != nil {
		return nil, err
	}
	return tbDisFuncs, nil
}
