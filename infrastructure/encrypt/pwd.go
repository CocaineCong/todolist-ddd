package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type PwdEncryptService struct {
	Base
}

const (
	PassWordCost = 12 // 密码加密难度
)

func NewPwdEncryptService() PwdEncryptService {
	return PwdEncryptService{}
}

func (p *PwdEncryptService) Encrypt(pwd []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(pwd, PassWordCost)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func (p *PwdEncryptService) Decrypt(data []byte) ([]byte, error) {
	return []byte{}, nil
}

func (p *PwdEncryptService) Check(pwd, src []byte) bool {
	return bcrypt.CompareHashAndPassword(pwd, src) == nil
}
