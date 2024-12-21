package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wuzigoweb/internal/data/model"
)

type ClassRepo interface {
	Save(context.Context, *model.Class) (int64, error)
	GetByName(context.Context, *model.Class) (*model.Class, error)
	GetAll(context.Context) ([]*model.Class, error)
	JoinClass(context.Context, *model.Class) (int64, error)
	SetAttend(context.Context, *model.Sign) (int64, error)
	GetAttend(context.Context, *model.Sign) (*model.Sign, error)
	NotAttend(context.Context) error

	//Update(context.Context, *Greeter) (*Greeter, errcode)
	//FindByID(context.Context, int64) (*Greeter, errcode)
	//ListByHello(context.Context, string) ([]*Greeter, errcode)
	//ListAll(context.Context) ([]*Greeter, errcode)
}

type ClassUsecase struct {
	repo ClassRepo
	log  *log.Helper
}

func NewClassUsecase(repo ClassRepo, logger log.Logger) *ClassUsecase {
	return &ClassUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ClassUsecase) CreateClass(ctx context.Context, class *model.Class) (int64, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 CreateClass 创建班级操作")
	return uc.repo.Save(ctx, class)
}

func (uc *ClassUsecase) GetClassByName(ctx context.Context, class *model.Class) (*model.Class, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 GetClassByName 操作")
	return uc.repo.GetByName(ctx, class)
}

func (uc *ClassUsecase) GetAllClass(ctx context.Context) ([]*model.Class, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 GetAllClass 获取所有班级操作")
	return uc.repo.GetAll(ctx)
}

func (uc *ClassUsecase) JoinClass(ctx context.Context, class *model.Class) (int64, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 JoinClass 加入班级操作")
	return uc.repo.JoinClass(ctx, class)
}

func (uc *ClassUsecase) SetAttendance(ctx context.Context, sign *model.Sign) (int64, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 SetAttendance 设置考勤操作")
	return uc.repo.SetAttend(ctx, sign)
}

func (uc *ClassUsecase) StudentSignIn(ctx context.Context, sign *model.Sign) (*model.Sign, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 StudentSignIn 学生签到操作")
	return uc.repo.GetAttend(ctx, sign)
}

func (uc *ClassUsecase) StudentNotAttend(ctx context.Context) error {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 StudentNotAttend 记录学生缺勤")
	return uc.repo.NotAttend(ctx)
}
