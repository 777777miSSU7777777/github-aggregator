package cryptofactory

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/encryption/aes"
	"github.com/777777miSSU7777777/github-aggregator/pkg/encryption/cryptoservice"
)

// New returns an instance of CryptoService.
// String param "algorithm" is responsible for switch-case which returns neccessary crypto service.
func New(algorithm string) cryptoservice.CryptoService {
	switch algorithm {
	case "aes":
		return &aes.AES{}
	default:
		return &aes.AES{}
	}
}
