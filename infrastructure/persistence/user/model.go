package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
}
