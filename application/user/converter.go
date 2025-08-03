package user

import (
	"github.com/CocaineCong/todolist-ddd/domain/user/entity"
	"github.com/CocaineCong/todolist-ddd/interfaces/types"
)

func Dto2Entity(user *types.UserReq) *entity.User {
	return &entity.User{
		Username: user.UserName,
		Password: user.Password,
	}
}
