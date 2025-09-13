package entity

import (
	"errors"
	"time"

	"github.com/CocaineCong/todolist-ddd/infrastructure/consts"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) IsValidUserName() bool {
	return len(u.Username) >= consts.UserNameLengthMin &&
		len(u.Username) <= consts.UserNameLengthMax
}

func (u *User) ChangePassword(newPassword string) error {
	if len(newPassword) < 6 {
		return errors.New(consts.UserPasswdMustMoreSix)
	}
	u.Password = newPassword
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) IsActive() bool {
	return u.ID > 0
}
