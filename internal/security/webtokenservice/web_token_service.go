package webtokenservice

import(
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/encryption/cryptoservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/cryptofactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/constants"
	"github.com/777777miSSU7777777/github-aggregator/pkg/http/cookieutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/encoding/base64util"
)

var cryptoService cryptoservice.CryptoService


func SetCryptoService(algorithm string){
	cryptoService = cryptofactory.New(algorithm)
}


func SetCryptoServiceKey(key []byte){
	cryptoService.SetKey(key)
}


func SetCryptoServiceIV(IV []byte){
	cryptoService.SetIV(IV)
}


func SaveToken(rw http.ResponseWriter, tkn string)(error){
	decodedTkn, err := base64util.Decode(tkn); if err != nil {
		return err
	}

	encryptedTkn, err := cryptoService.Encrypt(decodedTkn); if err != nil {
		return err
	}

	encodedTkn := base64util.Encode(encryptedTkn)

	cookieutil.SaveCookie(rw,constants.AccessToken, encodedTkn)

	return nil
}

func GetToken(req *http.Request)(string, error){
	encodedTkn, err := cookieutil.GetCookieValue(req,constants.AccessToken); if err != nil {
		return "", err
	}

	decodedTkn, err := base64util.Decode(encodedTkn); if err != nil {
		return "", err
	}

	decryptedTkn, err := cryptoService.Decrypt(decodedTkn); if err != nil {
		return "",err
	}

	tkn := base64util.Encode(decryptedTkn)

	return tkn, nil
}

func DeleteToken(rw http.ResponseWriter,req *http.Request)(error){
	err := cookieutil.DeleteCookie(rw, req, constants.AccessToken); if err != nil {
		return err
	}

	return nil
}