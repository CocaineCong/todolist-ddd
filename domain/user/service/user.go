package service

import (
	"context"
	"errors"

	"github.com/CocaineCong/todolist-ddd/domain/user/entity"
	"github.com/CocaineCong/todolist-ddd/domain/user/repository"
)

type UserDomain interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	FindUserByName(ctx context.Context, name string) (*entity.User, error)
	GetUserDetail(ctx context.Context, id uint) (*entity.User, error)
	CheckUserPwd(ctx context.Context, user *entity.User, src string) error
}

type UserDomainImpl struct {
	repo    repository.User
	encrypt repository.PwdEncrypt
}

func NewUserDomainImpl(repo repository.User, encrypt repository.PwdEncrypt) UserDomain {
	return &UserDomainImpl{repo: repo, encrypt: encrypt}
}

// CreateUser 创建用户
func (u *UserDomainImpl) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	// 加密
	encryptPwd, err := u.encrypt.Encrypt([]byte(user.Password))
	if err != nil {
		return nil, err
	}
	err = user.SetPwd(encryptPwd)
	if err != nil {
		return nil, err
	}
	return u.repo.CreateUser(ctx, user)
}

// FindUserByName 通过用户名找到用户
func (u *UserDomainImpl) FindUserByName(ctx context.Context, username string) (*entity.User, error) {
	return u.repo.GetUserByName(ctx, username)
}

// GetUserDetail 获取用户信息
func (u *UserDomainImpl) GetUserDetail(ctx context.Context, id uint) (*entity.User, error) {
	return u.repo.GetUserByID(ctx, id)
}

// CheckUserPwd 校验用户密码
func (u *UserDomainImpl) CheckUserPwd(ctx context.Context, user *entity.User, src string) error {
	if u.encrypt.Check([]byte(user.Password), []byte(src)) {
		return nil
	}
	return errors.New("wrong password")
}
