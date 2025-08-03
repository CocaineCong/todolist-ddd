package repository

import (
	"context"

	"github.com/CocaineCong/todolist-ddd/domain/user/entity"
)

type User interface {
	UserBase
}

type UserBase interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByName(ctx context.Context, username string) (*entity.User, error)
	GetUserByID(ctx context.Context, id uint) (*entity.User, error)
}
