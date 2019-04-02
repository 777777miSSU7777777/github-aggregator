package webtokenservice

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/encryption/cryptoservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/cryptofactory"
	"github.com/stretchr/testify/assert"
)

func TestSetCryptoService__AES__Equal(t *testing.T) {
	SetCryptoService("aes")

	aesInstance := cryptofactory.New("aes")

	assert.Equal(t, aesInstance, cryptoService)
}

func TestSetCryptoService__RandomString__Equal(t *testing.T) {
	SetCryptoService("test123")

	aesInstance := cryptofactory.New("aes")

	assert.Equal(t, aesInstance, cryptoService)
}

func TestSetCryptoServiceKey__SameKey__Equal(t *testing.T) {
	SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceKey(key)

	serviceKey := reflect.ValueOf(cryptoService).Elem().FieldByName("key").Bytes()

	assert.Equal(t, key, serviceKey)
}

func TestSetCryptoServiceKey__DifferentKeys__(t *testing.T) {
	SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceKey(key)

	serviceKey := reflect.ValueOf(cryptoService).Elem().FieldByName("key").Bytes()

	key, _ = randutil.GenerateRandomBytes(16)

	assert.NotEqual(t, key, serviceKey)
}

func TestSetCryptoServiceIV__SameIV__Equal(t *testing.T) {
	SetCryptoService("aes")

	IV, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceIV(IV)

	serviceIV := reflect.ValueOf(cryptoService).Elem().FieldByName("iv").Bytes()

	assert.Equal(t, IV, serviceIV)
}

func TestSetCryptoServiceIV__DifferentIV__NotEqual(t *testing.T) {
	SetCryptoService("aes")

	IV, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceIV(IV)

	serviceIV := reflect.ValueOf(cryptoService).Elem().FieldByName("iv").Bytes()

	IV, _ = randutil.GenerateRandomBytes(16)

	assert.NotEqual(t, IV, serviceIV)
}

func TestSaveToken__SameToken__Equal(t *testing.T) {
	SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceIV(IV)

	rw := httptest.NewRecorder()

	randomToken, _ := randutil.GenerateRandomBytes(16)

	encodedToken := encode(randomToken)

	SaveToken(rw, encodedToken)

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	cookie, _ := req.Cookie("access_token")

	encryptedToken, _ := encrypt(cryptoService, randomToken)

	testToken := encode(encryptedToken)

	assert.Equal(t, testToken, cookie.Value)
}

func TestSaveToken__DifferentTokens__NotEqual(t *testing.T) {
	SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceIV(IV)

	rw := httptest.NewRecorder()

	randomToken, _ := randutil.GenerateRandomBytes(16)

	encodedToken := encode(randomToken)

	SaveToken(rw, encodedToken)

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	cookie, _ := req.Cookie("access_token")

	randomToken, _ = randutil.GenerateRandomBytes(16)

	encodedToken = encode(randomToken)

	encryptedToken, _ := encrypt(cryptoService, randomToken)

	testToken := encode(encryptedToken)

	assert.NotEqual(t, testToken, cookie.Value)
}

func TestSaveToken__DecodeError__Error(t *testing.T) {
	SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceIV(IV)

	original := decode

	defer func() {
		decode = original
	}()

	decode = func(s string) ([]byte, error) {
		return nil, errors.New("DECODE ERROR")
	}

	rw := httptest.NewRecorder()

	err := SaveToken(rw, "123")

	assert.Error(t, err)
}

func TestSaveToken__EncryptError__Error(t *testing.T) {
	SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceIV(IV)

	original := encrypt

	defer func() {
		encrypt = original
	}()

	encrypt = func(cs cryptoservice.CryptoService, data []byte) ([]byte, error) {
		return nil, errors.New("ENCRYPT ERROR")
	}

	rw := httptest.NewRecorder()

	randomToken, _ := randutil.GenerateRandomBytes(16)

	encodedToken := encode(randomToken)

	err := SaveToken(rw, encodedToken)

	assert.Error(t, err)
}

func TestGetToken__SameToken__Equal(t *testing.T) {
	SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceIV(IV)

	rw := httptest.NewRecorder()

	randomToken, _ := randutil.GenerateRandomBytes(16)

	encodedToken := encode(randomToken)

	SaveToken(rw, encodedToken)

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	cookieToken, _ := GetToken(req)

	assert.Equal(t, encodedToken, cookieToken)
}

func TestGetToken__DifferentTokens__NotEqual(t *testing.T) {
	SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceIV(IV)

	rw := httptest.NewRecorder()

	randomToken, _ := randutil.GenerateRandomBytes(16)

	encodedToken := encode(randomToken)

	SaveToken(rw, encodedToken)

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	randomToken, _ = randutil.GenerateRandomBytes(16)

	encodedToken = encode(randomToken)

	cookieToken, _ := GetToken(req)

	assert.NotEqual(t, encodedToken, cookieToken)
}

func TestGetToken__CookieNotFound__Error(t *testing.T) {
	SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceIV(IV)

	req := &http.Request{}

	_, err := GetToken(req)

	assert.Error(t, err)
}

func TestGetToken__DecodeError__Error(t *testing.T) {
	SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceIV(IV)

	rw := httptest.NewRecorder()

	randomToken, _ := randutil.GenerateRandomBytes(16)

	encodedToken := encode(randomToken)

	SaveToken(rw, encodedToken)

	original := decode

	defer func() {
		decode = original
	}()

	decode = func(s string) ([]byte, error) {
		return nil, errors.New("DECODE ERROR")
	}

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	_, err := GetToken(req)

	assert.Error(t, err)
}

func TestGetToken__DecryptError__Error(t *testing.T) {
	SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceIV(IV)

	original := decrypt

	defer func() {
		decrypt = original
	}()

	decrypt = func(cs cryptoservice.CryptoService, data []byte) ([]byte, error) {
		return nil, errors.New("DECRYPT ERROR")
	}

	rw := httptest.NewRecorder()

	randomToken, _ := randutil.GenerateRandomBytes(16)

	encodedToken := encode(randomToken)

	SaveToken(rw, encodedToken)

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	_, err := GetToken(req)

	assert.Error(t, err)
}

func TestDeleteToken__ExistingCookie__Deleted(t *testing.T) {
	SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceIV(IV)

	rw := httptest.NewRecorder()

	randomToken, _ := randutil.GenerateRandomBytes(16)

	encodedToken := encode(randomToken)

	SaveToken(rw, encodedToken)

	req := &http.Request{Header: http.Header{"Cookie": rw.HeaderMap["Set-Cookie"]}}

	err := DeleteToken(rw, req)

	assert.Nil(t, err)
}

func TestDeleteToken__CookieNotFound__Error(t *testing.T) {
	SetCryptoService("aes")

	key, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceKey(key)

	IV, _ := randutil.GenerateRandomBytes(16)

	SetCryptoServiceIV(IV)

	rw := httptest.NewRecorder()

	req := &http.Request{}

	err := DeleteToken(rw, req)

	assert.Error(t, err)
}
