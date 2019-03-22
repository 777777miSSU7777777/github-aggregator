package cryptoservice


// CryptoService is an interface for block crypto systems with IV required.
type CryptoService interface {
	// SetKey set secret key for CryptoService instance.
	SetKey([]byte)

	// SetIV set initialization vector for CryptoService instance.
	SetIV([]byte)
 
	// Encrypt encrypts data using secret key and IV.
	Encrypt([]byte)([]byte, error)

	// Decrypt decrypts data using secret key and IV.
	Decrypt([]byte)([]byte, error)
}