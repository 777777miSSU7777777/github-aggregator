package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/http/cookieutil"
	"github.com/777777miSSU7777777/github-aggregator/internal/api"
	"github.com/777777miSSU7777777/github-aggregator/internal/view/index"

	"github.com/gorilla/mux"
)

var host string
var port string
var duration string

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "Defines host ip")
	flag.StringVar(&host, "h", "127.0.0.1", "Defines host ip")
	flag.StringVar(&port, "port", "8080", "Defines host port")
	flag.StringVar(&port, "p", "8080", "Defines host port")
	flag.StringVar(&duration, "duration", "1h", "Defines cookie expiration duration")
	flag.StringVar(&duration, "d", "1h", "Defines cookie expiration duration")
	flag.Parse()
}

func main() {	
	err := cookieutil.SetExpiration(duration); if err != nil {
		log.Fatalln(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/", index.Render).Methods("GET")
	router.HandleFunc("/auth", api.Auth).Methods("POST")
	router.HandleFunc("/logout", api.Logout).Methods("POST")
	http.Handle("/", router)

	log.Printf("Server started on %s:%s", host, port)
	
	err = http.ListenAndServe( fmt.Sprintf("%s:%s", host, port), nil); if err != nil {
		log.Fatalln(err)
	}

}
