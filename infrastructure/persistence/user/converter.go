package user

import (
	"github.com/CocaineCong/todolist-ddd/domain/user/entity"
)

func Entity2PO(user *entity.User) *User {
	return &User{
		UserName:       user.Username,
		PasswordDigest: user.Password,
	}
}

func PO2Entity(user *User) *entity.User {
	return &entity.User{
		ID:        user.ID,
		Username:  user.UserName,
		Password:  user.PasswordDigest,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
