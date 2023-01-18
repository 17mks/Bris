package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"student/internal/biz"
	"time"
)

type Student struct {
	ID        int32
	Name      string
	Info      string
	Status    int32
	UpdatedAt time.Time
	CreatedAt time.Time
}

type studentRepo struct {
	data *Data
	log  *log.Helper
}

func convertStudent(x Student) *biz.Student {
	return &biz.Student{
		ID:        x.ID,
		Name:      x.Name,
		Info:      x.Info,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}

// 初始化 studentRepo
func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *studentRepo) Save(ctx context.Context, stu *biz.Student) (*biz.Student, error) {
	var student Student
	//验证是否已经创建
	result := r.data.gormDB.Where(&biz.Student{ID: stu.ID}).First(&student)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "该学生已存在")
	}
	student.ID = stu.ID
	student.Name = stu.Name
	res := r.data.gormDB.Create(&student)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	return &biz.Student{
		ID:     stu.ID,
		Name:   stu.Name,
		Info:   stu.Info,
		Status: stu.Status,
	}, nil
}

func (r *studentRepo) GetStudent(ctx context.Context, id int32) (*biz.Student, error) {
	var stu biz.Student
	r.data.gormDB.Where("id = ?", id).First(&stu)
	r.log.WithContext(ctx).Info("gormDB: GetStudent, id: ", id)
	return &biz.Student{
		ID:        stu.ID,
		Name:      stu.Name,
		Status:    stu.Status,
		Info:      stu.Info,
		UpdatedAt: stu.UpdatedAt,
		CreatedAt: stu.CreatedAt,
	}, nil
}

func (r *studentRepo) Delete(ctx context.Context, id int32) (*biz.Student, error) {
	var stu biz.Student
	r.data.gormDB.Where("id = ?", id).Delete(&stu)
	r.log.WithContext(ctx).Infof("gormDB: DeleteStudent, id: ", id)
	return &biz.Student{
		ID:        stu.ID,
		Name:      stu.Name,
		Status:    stu.Status,
		Info:      stu.Info,
		UpdatedAt: stu.UpdatedAt,
		CreatedAt: stu.CreatedAt,
	}, nil
}

func (r *studentRepo) UpdateStudent(ctx context.Context, stu *biz.Student) (rv *biz.Student, err error) {
	u := new(Student)
	err = r.data.gormDB.Where("ID", stu.ID).First(u).Error
	if err != nil {
		return nil, err
	}
	err = r.data.gormDB.Model(&u).Updates(&Student{
		ID:   stu.ID,
		Name: stu.Name,
		Info: stu.Info,
	}).Error
	return &biz.Student{
		ID:   u.ID,
		Name: u.Name,
		Info: u.Info,
	}, nil
}
