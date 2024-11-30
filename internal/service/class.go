package service

import (
	"context"
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
