package util

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	PassWordCost = 12 // 密码加密难度
)

// EncryptPwd 设置密码
func EncryptPwd(password string) (pwd string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword 校验密码
func CheckPassword(pwdEncrypt, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdEncrypt), []byte(password))
	return err == nil
}
