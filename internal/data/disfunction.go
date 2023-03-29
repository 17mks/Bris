package data

import (
	"context"
	v1 "followup/api"
	"followup/gencode/chdb"
	"followup/internal/biz"
	"followup/utils"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	Log "log"
	"regexp"
	"strings"
	"time"
)

// TbDisFunction 病种表
type TbDisFunction struct {
	ID          string    `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'功能障碍编码'"`
	Name        string    `gorm:"column:name;type:varchar(512);not null;comment:'功能障碍名称'"`
	Description string    `gorm:"column:description;type:varchar(512);default:null;comment:'描述'"`
	Py          string    `gorm:"column:py;type:datetime;default:null;CURRENT_TIMESTAMP;comment:'功能障碍名称拼音'"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	DeleteAt    time.Time `gorm:"column:delete_at;type:datetime;default:not null;default:CURRENT_TIMESTAMP;comment:'删除时间'"`
}

func TimeToString(time *time.Time) string {
	if nil == time || time.IsZero() {
		return ""
	}
	return time.Format("2006-01-02 15:04:05")
}

// ParseTime 将时间字符串转换为time.Time类型
func ParseTime(value string) (*time.Time, error) {
	if "" == value {
		return nil, nil
	}
	// 处理"1971-04-07 00:00:00.0"格式
	if strings.HasSuffix(value, ".0") {
		value = value[0 : len(value)-2]
	}

	layout := "2006-01-02 15:04:05"
	// 时间格式正则判断
	dateTimePatter := `^[1-9]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])\s+(20|21|22|23|[0-1]\d):[0-5]\d:[0-5]\d$`
	datePatter := `^[1-9]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$`

	if match, err := regexp.Match(dateTimePatter, []byte(value)); err != nil {
		return nil, err
	} else if match {
		layout = "2006-01-02 15:04:05"
	}
	if match, err := regexp.Match(datePatter, []byte(value)); err != nil {
		return nil, err
	} else if match {
		layout = "2006-01-02"
	}
	parse, err := time.Parse(layout, value)
	if err != nil {
		return nil, err
	}

	return &parse, nil
}

func ParseTimeIgnoreErr(value string) *time.Time {
	timeStruct, err := ParseTime(value)
	if err != nil {
		Log.Println(err)
		return nil
	}

	return timeStruct
}

type DisFunctionRepo struct {
	data *Data
	log  *log.Helper
}

func NewDisFunctionRepo(data *Data, logger log.Logger) biz.DisFunctionRepo {
	return &DisFunctionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *DisFunctionRepo) Save(ctx context.Context, df *chdb.TbDisFunc) error {
	DisFunction := chdb.TbDisFunc{
		ID:          df.ID,
		Name:        df.Name,
		Description: df.Description,
		Py:          df.Py,
	}
	//验证是否已创建
	result := r.data.gormDB.Where(&chdb.TbDisFunc{ID: df.ID}).First(&DisFunction)
	if result.RowsAffected == 1 {
		return status.Errorf(codes.AlreadyExists, "该功能障碍已存在")
	}
	DisFunction.ID = df.ID
	DisFunction.Name = df.Name
	DisFunction.Description = df.Description
	DisFunction.Py = df.Py
	res := r.data.gormDB.Create(&DisFunction)
	return res.Error
}

func (r *DisFunctionRepo) Delete(ctx context.Context, id string) (string, error) {
	DisFunction := chdb.TbDisFunc{ID: id}
	if err := r.data.gormDB.Delete(&DisFunction).Error; err != nil {
		return "", nil
	}
	return DisFunction.ID, nil
}

func (r *DisFunctionRepo) Update(ctx context.Context, req *chdb.TbDisFunc) (res *chdb.TbDisFunc, err error) {
	DisFunction := new(chdb.TbDisFunc)
	err = r.data.gormDB.Where("ID", req.ID).First(DisFunction).Error

	if err != nil {
		return nil, err
	}
	err = r.data.gormDB.Model(&DisFunction).Updates(&chdb.TbDisFunc{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Py:          req.Py,
	}).Error

	return &chdb.TbDisFunc{
		ID:          DisFunction.ID,
		Name:        DisFunction.Name,
		Description: DisFunction.Description,
		Py:          DisFunction.Py,
	}, nil
}

func (r *DisFunctionRepo) GetByID(ctx context.Context, id string) (rv *chdb.TbDisFunc, err error) {
	DisFunction := new(chdb.TbDisFunc)
	result := r.data.gormDB.Where("id = ?", id).First(DisFunction)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("DisFunction", "not found by id")
	}
	if result.Error != nil {
		return nil, err
	}
	return &chdb.TbDisFunc{
		ID:          DisFunction.ID,
		Name:        DisFunction.Name,
		Description: DisFunction.Description,
		CreateTime:  DisFunction.CreateTime,
		UpdateTime:  DisFunction.UpdateTime,
		DeleteAt:    DisFunction.DeleteAt,
	}, nil
}

func (r *DisFunctionRepo) Filter(ctx context.Context, apiReq *v1.DisFunctionFilterRequest) ([]chdb.TbDisFunc, int64, error) {
	serverContext, _ := transport.FromServerContext(ctx)

	rv := serverContext.RequestHeader().Get("Usertoken")
	Token, _ := utils.ParseToken(rv)
	Log.Println(Token)

	funclist := make([]chdb.TbDisFunc, 0)
	var count int64 = 0

	filter := apiReq.Filter
	tx := r.data.gormDB.Model(&chdb.TbDisFunc{})

	if filter != nil {
		if "" != filter.Key {
			keyLike := AddLikeCharToStr(filter.Key)
			tx.Where("name like ? or py like ?", keyLike, keyLike)
		}
		if len(filter.Ids) > 0 {
			tx.Where("id in ?", filter.Ids)
		}
	}

	tx.Order("update_time DESC")
	tx.Count(&count)

	result := tx.Find(&funclist)
	return funclist, count, result.Error
}

func (r *DisFunctionRepo) DisFuncPseudoDelById(ctx context.Context, id int64) (int64, error) {
	if err := r.data.gormDB.Model(&chdb.TbDisFunc{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (r *DisFunctionRepo) HasDisFunc(ctx context.Context, disFuncName string) (*chdb.TbDisFunc, bool, error) {
	tbDisFunc := chdb.TbDisFunc{}
	if err := r.data.gormDB.Model(&chdb.TbDisFunc{}).Where("name = ? and deleted_at is null", disFuncName).First(&tbDisFunc).Error; err != nil {
		return nil, false, nil
	}
	return nil, true, nil
}

func (r *DisFunctionRepo) QueryDisFunByIds(ctx context.Context, ids []string) ([]chdb.TbDisFunc, error) {
	tbDisFunc := make([]chdb.TbDisFunc, 0)
	if err := r.data.gormDB.Where("id in ?", ids).Find(&tbDisFunc).Error; err != nil {
		return nil, err
	}
	return tbDisFunc, nil
}
