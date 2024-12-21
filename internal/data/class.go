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

	s := r.data.query.User
	// 获取当前用户名字，创建班级的时候把老师的名字也存入数据库
	users, _ := s.WithContext(ctx).Where(s.ID.Eq(ctx.Value("user_id").(int64))).First()
	class.Teacher = users.Username

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

func (r *classRepo) GetAll(ctx context.Context) ([]*model.Class, error) {
	u := r.data.query.Class
	classes, _ := u.WithContext(ctx).Find()
	return classes, nil
}

func (r *classRepo) JoinClass(ctx context.Context, class *model.Class) (int64, error) {
	c := r.data.query.Class
	u := r.data.query.User
	// 获取当前用户名字
	users, _ := u.WithContext(ctx).Where(u.ID.Eq(ctx.Value("user_id").(int64))).First()

	// 将学生加入班级
	_, err := c.WithContext(ctx).Where(c.ClassName.Eq(class.ClassName)).Update(c.Student, users.Username)
	if err != nil {
		return -3, fmt.Errorf("班级信息存入数据库失败")
	}

	return 0, nil
}

func (r *classRepo) SetAttend(ctx context.Context, sign *model.Sign) (int64, error) {
	u := r.data.query.Sign
	err := u.WithContext(ctx).Save(sign)
	if err != nil {
		return -3, fmt.Errorf("班级信息存入数据库失败")
	}
	return 0, nil
}

func (r *classRepo) GetAttend(ctx context.Context, sign *model.Sign) (*model.Sign, error) {
	u := r.data.query.Sign
	signs, _ := u.WithContext(ctx).Where(u.AttendName.Eq(sign.AttendName)).First()
	return signs, nil
}

func (r *classRepo) NotAttend(ctx context.Context) error {
	u := r.data.query.User
	c := r.data.query.Class
	// 获取当前用户名字
	users, _ := u.WithContext(ctx).Where(u.ID.Eq(ctx.Value("user_id").(int64))).First()

	// 获取当前用户的班级
	classes, _ := c.WithContext(ctx).Where(c.Student.Eq(users.Username)).First()
	classes.StudentAbsence = users.Username
	classes.StudentAbsenceCount++
	_, err := c.WithContext(ctx).Where(c.ClassName.Eq(classes.ClassName)).Update(c.StudentAbsence, classes.StudentAbsence)
	_, err = c.WithContext(ctx).Where(c.ClassName.Eq(classes.ClassName)).Update(c.StudentAbsenceCount, classes.StudentAbsenceCount)
	if err != nil {
		return fmt.Errorf("记录缺勤学生信息失败")
	}
	return nil
}
