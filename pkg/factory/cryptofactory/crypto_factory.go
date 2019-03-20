package cryptofactory

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/encryption/cryptoservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/encryption/aes"
)

func New(algorithm string)(cryptoservice.CryptoService){
	switch algorithm {
		case "aes":
			return &aes.AES{}
		default:
			return nil
	}
}