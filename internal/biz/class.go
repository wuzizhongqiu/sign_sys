package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wuzigoweb/internal/data/model"
)

type ClassRepo interface {
	Save(context.Context, *model.Class) (int64, error)

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
