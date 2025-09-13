package encrypt

type Service interface {
	Encrypt(data []byte) ([]byte, error)

	Decrypt(data []byte) ([]byte, error)
}
