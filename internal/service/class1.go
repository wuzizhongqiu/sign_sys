package service

import (
	"context"
	"time"
	"wuzigoweb/internal/biz"
	"wuzigoweb/internal/data/model"

	pb "wuzigoweb/api/http/class"
)

type ClassService struct {
	uc *biz.ClassUsecase
}

func NewClassService(uc *biz.ClassUsecase) *ClassService {
	return &ClassService{uc: uc}
}

// CreateClass 创建班级
func (s *ClassService) CreateClass(ctx context.Context, req *pb.CreateClassRequest) (*pb.CreateClassReply, error) {
	id, err := s.uc.CreateClass(ctx, &model.Class{
		ClassName: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateClassReply{Id: id}, nil
}

// GetClassByName 根据班级名称获取班级
func (s *ClassService) GetClassByName(ctx context.Context, req *pb.GetClassByNameRequest) (*pb.GetClassByNameReply, error) {
	class, err := s.uc.GetClassByName(ctx, &model.Class{
		ClassName: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetClassByNameReply{Name: class.ClassName, Student: class.Student, Teacher: class.Teacher, StudentAbsence: class.StudentAbsence, StudentAbsenceCount: int64(class.StudentAbsenceCount)}, nil
}

// GetAllClass 获取所有班级
func (s *ClassService) GetAllClass(ctx context.Context, req *pb.GetAllClassRequest) (*pb.GetAllClassReply, error) {
	classes, err := s.uc.GetAllClass(ctx)
	if err != nil {
		return nil, err
	}
	var classList []string
	for _, v := range classes {
		classList = append(classList, v.ClassName)
	}
	return &pb.GetAllClassReply{Name: classList}, nil
}

// JoinClass 加入班级
func (s *ClassService) JoinClass(ctx context.Context, req *pb.JoinClassRequest) (*pb.JoinClassReply, error) {
	id, err := s.uc.JoinClass(ctx, &model.Class{
		ClassName: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.JoinClassReply{Id: id}, nil
}

// SetAttendance 创建考勤
func (s *ClassService) SetAttendance(ctx context.Context, req *pb.SetAttendanceRequest) (*pb.SetAttendanceReply, error) {
	id, err := s.uc.SetAttendance(ctx, &model.Sign{
		AttendName: req.AttendName,
		SignTime:   req.Time,
	})
	if err != nil {
		return nil, err
	}
	return &pb.SetAttendanceReply{Id: id}, nil
}

// StudentSignIn 学生签到
func (s *ClassService) StudentSignIn(ctx context.Context, req *pb.StudentSignInRequest) (*pb.StudentSignInReply, error) {
	// 获取这个签到的信息
	sign, err := s.uc.StudentSignIn(ctx, &model.Sign{
		AttendName: req.AttendName,
	})
	if err != nil {
		return nil, err
	}

	// 根据这次签到的信息判断能不能签到
	var message string
	Deadline := sign.CreateTime.Add(time.Duration(sign.SignTime) * time.Second)
	currentTime := time.Now()
	if currentTime.After(Deadline) {
		message = "签到已过期"
		// 将当前学生归入缺勤名单
		_ = s.uc.StudentNotAttend(ctx)
	} else {
		message = "签到成功"
	}

	return &pb.StudentSignInReply{Message: message}, nil
}
