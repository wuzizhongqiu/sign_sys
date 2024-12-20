package biz

import (
	"context"
	"encoding/json"
	"errors"
	"wuzigoweb/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	// 添加与签到相关的接口方法
	GetSignByID(ctx context.Context, id int64) (*model.Sign, error)
	UpdateSign(ctx context.Context, sign *model.Sign) error
	GetStudentList(ctx context.Context, signID int64) ([]string, error)
	Save(context.Context, *model.User) (int64, error)
	Login(context.Context, *model.User) (*model.User, error)
	ListUser(ctx context.Context, id int64) (*model.User, error)
	GetUser(ctx context.Context) (*model.User, error)
	ListUserPage(ctx context.Context, current, pageSize int32) ([]*model.User, int, error)
	//Update(context.Context, *Greeter) (*Greeter, errcode)
	//FindByID(context.Context, int64) (*Greeter, errcode)
	//ListByHello(context.Context, string) ([]*Greeter, errcode)
	//ListAll(context.Context) ([]*Greeter, errcode)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 CreateUser 创建用户操作")
	return uc.repo.Save(ctx, user)
}

func (uc *UserUsecase) UserLogin(ctx context.Context, user *model.User) (*model.User, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 UserLogin 用户登录逻辑")
	return uc.repo.Login(ctx, user)
}

func (uc *UserUsecase) GetUserList(ctx context.Context, id int64) (*model.User, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 GetUserList 获取用户列表")
	return uc.repo.ListUser(ctx, id)
}

func (uc *UserUsecase) GetCurrentUser(ctx context.Context) (*model.User, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 GetCurrentUser 获取当前用户信息")
	return uc.repo.GetUser(ctx)
}

func (uc *UserUsecase) ListUserPage(ctx context.Context, current, pageSize int32) ([]*model.User, int, error) {
	uc.log.WithContext(ctx).Debugf("biz: 正在执行 ListUserPage 获取用户分页逻辑")
	return uc.repo.ListUserPage(ctx, current, pageSize)
}

// 实现 StudentSign 方法
func (uc *UserUsecase) StudentSign(ctx context.Context, signID int64, studentName string) error {
	// 调用 repo 获取签到记录
	sign, err := uc.repo.GetSignByID(ctx, signID)
	if err != nil {
		return err // 记录不存在或查询失败，返回错误
	}

	// 解析已签到学生列表
	var signedStudents []string
	if err := json.Unmarshal([]byte(sign.StudentSign), &signedStudents); err != nil {
		return errors.New("已签到学生列表解析失败")
	}

	// 解析应签到学生名单
	var studentList []string
	if err := json.Unmarshal([]byte(sign.Student), &studentList); err != nil {
		return errors.New("应签到学生名单解析失败")
	}

	// 判断学生是否在应签到名单中
	needsSign := false
	for _, name := range studentList {
		if name == studentName {
			needsSign = true
			break
		}
	}
	if !needsSign {
		return errors.New("学生不在应签到名单中，无需签到")
	}

	// 检查学生是否已签到
	for _, name := range signedStudents {
		if name == studentName {
			return errors.New("学生已签到")
		}
	}

	// 将学生加入已签到列表
	signedStudents = append(signedStudents, studentName)

	// 更新签到记录
	updatedData, err := json.Marshal(signedStudents)
	if err != nil {
		return errors.New("更新签到数据失败")
	}
	sign.StudentSign = string(updatedData)

	// 调用 repo 更新数据库
	return uc.repo.UpdateSign(ctx, sign)
}
