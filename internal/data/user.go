package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"wuzigoweb/internal/biz"
	"wuzigoweb/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, user *model.User) (int64, error) {
	u := r.data.query.User
	// 查看用户名是否被占用
	users, _ := u.WithContext(ctx).Where(u.Account.Eq(user.Account)).First()
	if users != nil {
		return -2, nil
	}
	// 存入数据库
	err := u.WithContext(ctx).Save(user)
	if err != nil {
		return -3, fmt.Errorf("用户注册存入数据库失败")
	}
	return user.ID, err
}

func (r *userRepo) Login(ctx context.Context, user *model.User) (*model.User, error) {
	u := r.data.query.User
	users, err := u.WithContext(ctx).Where(u.Account.Eq(user.Account)).First()
	if err != nil {
		r.log.WithContext(ctx).Debug("查询用户信息失败：", err)
		return nil, err
	}
	if users == nil {
		r.log.WithContext(ctx).Debugf("users 是空")
		return nil, err
	}
	if users.Password != user.Password {
		return nil, err
	}

	return users, err
}

func (r *userRepo) ListUser(ctx context.Context, id int64) (*model.User, error) {
	u := r.data.query.User
	users, err := u.WithContext(ctx).Where(u.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) GetUser(ctx context.Context) (*model.User, error) {
	u := r.data.query.User
	users, err := u.WithContext(ctx).Where(u.ID.Eq(ctx.Value("user_id").(int64))).First()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) ListUserPage(ctx context.Context, current, pageSize int32) ([]*model.User, int, error) {
	u := r.data.query.User
	users, err := u.WithContext(ctx).Offset(int(current - 1)).Limit(int(pageSize)).Find()
	count, err := u.WithContext(ctx).Count()
	if err != nil {
		return nil, 0, err
	}
	return users, int(count), nil
}

func (r *userRepo) GetSignByID(ctx context.Context, id int64) (*model.Sign, error) {
	signQuery := r.data.query.Sign // 获取 Sign 表查询器
	// 使用 Where 和 First 查询记录
	sign, err := signQuery.WithContext(ctx).Where(signQuery.ID.Eq(id)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) { // 使用 gorm.ErrRecordNotFound 检查记录是否存在
		return nil, errors.New("签到记录不存在")
	}
	return sign, err
}

func (r *userRepo) UpdateSign(ctx context.Context, sign *model.Sign) error {
	signQuery := r.data.query.Sign // 获取 Sign 表查询器
	sign.UpdateTime = time.Now()   // 更新最后修改时间
	// 使用 Save 保存记录
	err := signQuery.WithContext(ctx).Save(sign)
	return err
}

func (r *userRepo) GetStudentList(ctx context.Context, signID int64) ([]string, error) {
	signQuery := r.data.query.Sign // 获取 Sign 表查询器

	// 查询指定签到记录
	sign, err := signQuery.WithContext(ctx).Where(signQuery.ID.Eq(signID)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("签到记录不存在")
	}
	if err != nil {
		return nil, err // 其他查询错误
	}

	// 解析 JSON 格式的学生名单
	var studentList []string
	err = json.Unmarshal([]byte(sign.Student), &studentList)
	if err != nil {
		return nil, errors.New("解析学生名单失败")
	}

	return studentList, nil
}
