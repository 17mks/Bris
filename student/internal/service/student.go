package service

import (
	"context"
	pb "student/api/student/v1"
	"student/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type StudentService struct {
	pb.UnimplementedStudentServer

	student *biz.StudentUsecase
	log     *log.Helper
}

func NewStudentService(stu *biz.StudentUsecase, logger log.Logger) *StudentService {
	return &StudentService{
		student: stu,
		log:     log.NewHelper(logger),
	}
}

func (s *StudentService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentReply, error) {
	stu, err := s.student.CreateStudent(ctx, &biz.Student{
		ID:   req.Id,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateStudentReply{
		Id:     stu.ID,
		Status: stu.Status,
		Name:   stu.Name,
	}, nil
}

func (s *StudentService) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentReply, error) {
	stu, err := s.student.Delete(ctx, req.Id)

	if err != nil {
		return nil, err
	}
	return &pb.DeleteStudentReply{
		Id:     stu.ID,
		Status: stu.Status,
		Name:   stu.Name,
	}, nil
}

func (s *StudentService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentReply, error) {
	stu, err := s.student.UpdateStudent(ctx, &biz.Student{
		ID:   req.Id,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateStudentReply{
		Id:     stu.ID,
		Status: stu.Status,
		Name:   stu.Name,
	}, nil
}

//	func (s *StudentService) Liststudent(ctx context.Context, req *pb.ListStudentRequest) (*pb.ListStudentReply, error) {
//		return &pb.ListStudentReply{}, nil
//	}
//
// 获取学生信息
func (s *StudentService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentReply, error) {
	stu, err := s.student.Get(ctx, req.Id)

	if err != nil {
		return nil, err
	}
	return &pb.GetStudentReply{
		Id:     stu.ID,
		Status: stu.Status,
		Name:   stu.Name,
	}, nil
}
