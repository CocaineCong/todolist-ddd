package entity

import (
	"time"

	"github.com/CocaineCong/todolist-ddd/consts"
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

func (u *User) SetPwd(pwd []byte) error {
	u.Password = string(pwd)
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) IsActive() bool {
	return u.ID > 0
}
