package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"wuzigoweb/internal/biz"
	"wuzigoweb/internal/data/model"
)

type classRepo struct {
	data *Data
	log  *log.Helper
}

func NewClassRepo(data *Data, logger log.Logger) biz.ClassRepo {
	return &classRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *classRepo) Save(ctx context.Context, class *model.Class) (int64, error) {
	u := r.data.query.Class
	// 查看班级名是否被占用
	classNames, _ := u.WithContext(ctx).Where(u.ClassName.Eq(class.ClassName)).First()
	if classNames != nil {
		return -2, fmt.Errorf("班级名称已被占用")
	}
	// 存入数据库
	err := u.WithContext(ctx).Save(class)
	if err != nil {
		return -3, fmt.Errorf("班级信息存入数据库失败")
	}
	return class.ID, nil
}

func (r *classRepo) GetByName(ctx context.Context, class *model.Class) (*model.Class, error) {
	u := r.data.query.Class
	classes, _ := u.WithContext(ctx).Where(u.ClassName.Eq(class.ClassName)).First()
	return classes, nil
}
