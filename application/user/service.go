package user

import (
	"context"
	"errors"
	"sync"

	"github.com/CocaineCong/todolist-ddd/domain/user/service"
	"github.com/CocaineCong/todolist-ddd/interfaces/types"
)

type Service interface {
	Register(ctx context.Context, user *types.UserReq) (any, error)
	Login(ctx context.Context, user *types.UserReq) (any, error)
	GetUserInfo(ctx context.Context) (any, error)
}

var ServiceImplIns *ServiceImpl
var ServiceImplOnce sync.Once

type ServiceImpl struct {
	ud service.UserDomain
}

func GetServiceImpl(srv service.UserDomain) *ServiceImpl {
	ServiceImplOnce.Do(func() {
		ServiceImplIns = &ServiceImpl{ud: srv}
	})
	return ServiceImplIns
}

func (s *ServiceImpl) Register(ctx context.Context, req *types.UserReq) (any, error) {
	entity := types.UserDTO2Entity(req)
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
	resp := &types.UserResp{
		ID:       user.ID,
		UserName: user.Username,
		CreateAt: user.CreatedAt.Unix(),
	}
	return resp, nil
}

func (s *ServiceImpl) Login(ctx context.Context, req *types.UserReq) (any, error) {
	entity := types.UserDTO2Entity(req)
	user, err := s.ud.FindUserByName(ctx, req.UserName)
	if err != nil {
		return nil, err
	}
	// 检查密码
	err = s.ud.CheckPwd(ctx, user, entity.Password)
	if err != nil {
		return nil, errors.New("invalid password")
	}
	// 生成token
	token, err := s.ud.GenerateToken(ctx, user)
	if err != nil {
		return nil, err
	}
	resp := &types.TokenData{
		User: types.UserResp{
			ID:       user.ID,
			UserName: user.Username,
			CreateAt: user.CreatedAt.Unix(),
		},
		Token: token,
	}
	return resp, nil
}

func (s *ServiceImpl) GetUserInfo(ctx context.Context) (any, error) {
	return nil, nil
}
