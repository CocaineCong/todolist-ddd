package user

import (
	"context"
	"errors"
	"sync"

	"github.com/CocaineCong/todolist-ddd/domain/user/entity"
	"github.com/CocaineCong/todolist-ddd/domain/user/service"
	"github.com/CocaineCong/todolist-ddd/infrastructure/auth"
	"github.com/CocaineCong/todolist-ddd/infrastructure/interfaces/types"
)

type Service interface {
	Register(ctx context.Context, user *types.UserReq) (any, error)
	Login(ctx context.Context, user *types.UserReq) (any, error)
	GetUserInfo(ctx context.Context) (any, error)
}

type ServiceImpl struct {
	ud           service.UserDomain
	tokenService auth.TokenService // 依赖抽象接口
}

var (
	ServiceImplIns  *ServiceImpl
	ServiceImplOnce sync.Once
)

func GetServiceImpl(srv service.UserDomain) *ServiceImpl {
	ServiceImplOnce.Do(func() {
		ServiceImplIns = &ServiceImpl{ud: srv}
	})
	return ServiceImplIns
}

// Register 用户注册
func (s *ServiceImpl) Register(ctx context.Context, entity *entity.User) (any, error) {
	userExist, err := s.ud.FindUserByName(ctx, entity.Username)
	if err != nil {
		return nil, err
	}
	if userExist.ID != 0 {
		return nil, errors.New("user exist")
	}
	// 加密
	entityEncrypt, err := s.ud.EncryptPwd(ctx, entity)
	if err != nil {
		return nil, err
	}
	// 创建用户
	user, err := s.ud.CreateUser(ctx, entityEncrypt)
	if err != nil {
		return nil, err
	}

	return RegisterResponse(user), nil
}

// Login 用户登陆
func (s *ServiceImpl) Login(ctx context.Context, entity *entity.User) (any, error) {
	user, err := s.ud.FindUserByName(ctx, entity.Username)
	if err != nil {
		return nil, err
	}

	// 检查密码
	err = s.ud.CheckPwd(ctx, user, entity.Password)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// 生成token
	token, err := s.tokenService.GenerateToken(ctx, user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	return LoginResponse(user, token), nil
}

func (s *ServiceImpl) GetUserInfo(ctx context.Context) (any, error) {
	return nil, nil
}
