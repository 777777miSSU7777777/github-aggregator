package webtokenservice

import (
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/constants"
	"github.com/777777miSSU7777777/github-aggregator/pkg/encoding/base64util"
	"github.com/777777miSSU7777777/github-aggregator/pkg/encryption/cryptoservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/cryptofactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/http/cookieutil"
)

var cryptoService cryptoservice.CryptoService

// SetCryptoService sets crypto service for web token service.
// Algorithm should be presented as string.
func SetCryptoService(algorithm string) {
	cryptoService = cryptofactory.New(algorithm)
}

// SetCryptoServiceKey sets secret for crypto servivce of web token service.
// Key should be presented as byte array.
func SetCryptoServiceKey(key []byte) {
	cryptoService.SetKey(key)
}

// SetCryptoServiceIV sets IV for crypto service of web token service.
// IV should be presented as byte array.
func SetCryptoServiceIV(IV []byte) {
	cryptoService.SetIV(IV)
}

// SaveToken encrypts and saves access token to cookie.
// Access token should be presented as string.
// If base64util.Decode or cryptoService.Encrypt or base64util.Encode
// occurs any error, this will be returned.
func SaveToken(rw http.ResponseWriter, tkn string) error {
	decodedTkn, err := base64util.Decode(tkn)
	if err != nil {
		return err
	}

	encryptedTkn, err := cryptoService.Encrypt(decodedTkn)
	if err != nil {
		return err
	}

	encodedTkn := base64util.Encode(encryptedTkn)

	cookieutil.SaveCookie(rw, constants.AccessToken, encodedTkn)

	return nil
}

// GetToken returns decrypted access token from cookie.
// Access token is presented as string.
// If cookieutil.GetCookieValue or base64util.Decode or cryptoService.Decrypt
// or base64util.Encode occurs any error, this will be returned.
func GetToken(req *http.Request) (string, error) {
	encodedTkn, err := cookieutil.GetCookieValue(req, constants.AccessToken)
	if err != nil {
		return "", err
	}

	decodedTkn, err := base64util.Decode(encodedTkn)
	if err != nil {
		return "", err
	}

	decryptedTkn, err := cryptoService.Decrypt(decodedTkn)
	if err != nil {
		return "", err
	}

	tkn := base64util.Encode(decryptedTkn)

	return tkn, nil
}

// DeleteToken deletes access token from cookie.
// If cookieutil.DeleteCookie occurs any error, this will be returned.
func DeleteToken(rw http.ResponseWriter, req *http.Request) error {
	err := cookieutil.DeleteCookie(rw, req, constants.AccessToken)
	if err != nil {
		return err
	}

	return nil
}
