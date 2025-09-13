package service

import (
	"context"

	"github.com/CocaineCong/todolist-ddd/domain/user/entity"
	"github.com/CocaineCong/todolist-ddd/domain/user/repository"
)

type UserDomain interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
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

// FindUserByName 通过用户名找到用户
func (u *UserDomainImpl) FindUserByName(ctx context.Context, username string) (*entity.User, error) {
	return u.repo.GetUserByName(ctx, username)
}

// GetUserDetail 获取用户信息
func (u *UserDomainImpl) GetUserDetail(ctx context.Context, id uint) (*entity.User, error) {
	return u.repo.GetUserByID(ctx, id)
}
