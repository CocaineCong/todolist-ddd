package entity

import (
	"errors"
	"time"

	"github.com/CocaineCong/todolist-ddd/infrastructure/consts"
	"github.com/CocaineCong/todolist-ddd/infrastructure/encrypt"
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

func (u *User) EncryptPwd(pwd string) error {
	ps := encrypt.NewPwdEncryptService()
	password, err := ps.Encrypt([]byte(pwd))
	if err != nil {
		return err
	}
	u.Password = string(password)
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) CheckPwd(src string) error {
	ps := encrypt.NewPwdEncryptService()
	check := ps.Check([]byte(u.Password), []byte(src))
	if !check {
		return errors.New("wrong password")
	}
	return nil
}

func (u *User) IsActive() bool {
	return u.ID > 0
}
