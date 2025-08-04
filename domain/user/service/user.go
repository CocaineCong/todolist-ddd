package service

import (
	"context"
	"errors"

	"github.com/CocaineCong/todolist-ddd/domain/user/entity"
	"github.com/CocaineCong/todolist-ddd/domain/user/repository"
	"github.com/CocaineCong/todolist-ddd/infrastructure/common/jwt"
	"github.com/CocaineCong/todolist-ddd/infrastructure/common/util"
)

type UserDomain interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	EncryptPwd(ctx context.Context, user *entity.User) (*entity.User, error)
	CheckPwd(ctx context.Context, user *entity.User, pwd string) error
	GenerateToken(ctx context.Context, user *entity.User) (string, error)
	ParseToken(ctx context.Context, token string) (*entity.User, error)
	FindUserByName(ctx context.Context, name string) (*entity.User, error)
	GetUserDetail(ctx context.Context, id uint) (*entity.User, error)
}

type UserDomainImpl struct {
	repo repository.User
}

func NewUserDomainImpl(repo repository.User) UserDomain {
	return &UserDomainImpl{repo: repo}
}

// CreateUser 创建用户
func (u *UserDomainImpl) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return u.repo.CreateUser(ctx, user)
}

// EncryptPwd 加密密码
func (u *UserDomainImpl) EncryptPwd(_ context.Context, user *entity.User) (*entity.User, error) {
	pwd, err := util.EncryptPwd(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = pwd
	return user, nil
}

// CheckPwd 检查密码
func (u *UserDomainImpl) CheckPwd(_ context.Context, user *entity.User, pwd string) error {
	if ok := util.CheckPassword(user.Password, pwd); !ok {
		return errors.New("invalid password")
	}
	return nil
}

// FindUserByName 通过用户名找到用户
func (u *UserDomainImpl) FindUserByName(ctx context.Context, username string) (*entity.User, error) {
	return u.repo.GetUserByName(ctx, username)
}

// GetUserDetail 获取用户信息
func (u *UserDomainImpl) GetUserDetail(ctx context.Context, id uint) (*entity.User, error) {
	return u.repo.GetUserByID(ctx, id)
}

// GenerateToken 生成token
func (u *UserDomainImpl) GenerateToken(_ context.Context, user *entity.User) (string, error) {
	token, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

// ParseToken 解密token
func (u *UserDomainImpl) ParseToken(_ context.Context, token string) (*entity.User, error) {
	claim, err := jwt.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:       claim.Id,
		Username: claim.Username,
	}, nil
}
