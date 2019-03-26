package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/api"
	"github.com/777777miSSU7777777/github-aggregator/internal/security/webtokenservice"
	"github.com/777777miSSU7777777/github-aggregator/internal/view/index"
	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/encoding/base64util"
	"github.com/777777miSSU7777777/github-aggregator/pkg/http/cookieutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"

	"github.com/gorilla/mux"
)

var host string
var port string
var duration string
var algorithm string
var key string
var iv string

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "Defines host ip")
	flag.StringVar(&host, "h", "127.0.0.1", "Defines host ip")
	flag.StringVar(&port, "port", "8080", "Defines host port")
	flag.StringVar(&port, "p", "8080", "Defines host port")
	flag.StringVar(&duration, "duration", "1h", "Defines cookie expiration duration")
	flag.StringVar(&duration, "d", "1h", "Defines cookie expiration duration")
	encryptionInitSetup()
	flag.Parse()
}

func encryptionInitSetup() {
	flag.StringVar(&algorithm, "algorithm", "aes", "Defines token encryption algorithm")
	flag.StringVar(&algorithm, "a", "aes", "Defines token encryption algorithm")
	randomBytes, err := randutil.GenerateRandomBytes(16)
	if err != nil {
		log.Error.Fatalln(err)
	}
	flag.StringVar(&key, "k", base64util.Encode(randomBytes), "Defines encryption key")
	randomBytes, err = randutil.GenerateRandomBytes(16)
	if err != nil {
		log.Error.Fatalln(err)
	}
	flag.StringVar(&iv, "iv", base64util.Encode(randomBytes), "Defines initialization vector")
}

func encryptionSetup() {
	err := cookieutil.SetExpiration(duration)
	if err != nil {
		log.Error.Fatalln(err)
	}

	webtokenservice.SetCryptoService(algorithm)

	Key, err := base64util.Decode(key)
	if err != nil {
		log.Error.Fatalln(err)
	}

	webtokenservice.SetCryptoServiceKey(Key)

	IV, err := base64util.Decode(iv)
	if err != nil {
		log.Error.Fatalln(err)
	}

	webtokenservice.SetCryptoServiceIV(IV)
}

func main() {
	encryptionSetup()

	router := mux.NewRouter()

	router.HandleFunc("/", index.Render).Methods("GET")
	router.HandleFunc("/auth", api.Auth).Methods("POST")
	router.HandleFunc("/logout", api.Logout).Methods("POST")
	http.Handle("/", router)

	log.Info.Printf("Server started on %s:%s", host, port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	if err != nil {
		log.Error.Fatalln(err)
	}

}
