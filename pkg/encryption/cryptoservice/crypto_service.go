package cryptoservice

type CryptoService interface {
	SetKey([]byte)
	SetIV([]byte)
	Encrypt([]byte)([]byte, error)
	Decrypt([]byte)([]byte, error)
}