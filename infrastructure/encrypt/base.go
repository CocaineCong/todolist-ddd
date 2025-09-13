package encrypt

type Base interface {
	Encrypt(data []byte) ([]byte, error)

	Decrypt(data []byte) ([]byte, error)

	Check(data, src []byte) bool
}
