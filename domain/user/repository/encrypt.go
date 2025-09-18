package repository

type PwdEncrypt interface {
	Encrypt(data []byte) ([]byte, error)

	Decrypt(data []byte) ([]byte, error)

	Check(data, src []byte) bool
}
