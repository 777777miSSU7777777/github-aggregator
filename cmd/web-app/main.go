package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/api"
	"github.com/777777miSSU7777777/github-aggregator/internal/view/index"

	"github.com/gorilla/mux"
)

var host string
var port string

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "Defines host ip")
	flag.StringVar(&host, "h", "127.0.0.1", "Defines host ip")
	flag.StringVar(&port, "port", "8080", "Defines host port")
	flag.StringVar(&port, "p", "8080", "Defines host port")
	flag.Parse()
}

func main() {
	log.Printf("Server started on %s:%s", host, port)

	router := mux.NewRouter()

	router.HandleFunc("/", index.Render).Methods("GET")
	router.HandleFunc("/auth", api.Auth).Methods("POST")
	router.HandleFunc("/logout", api.Logout).Methods("POST")
	http.Handle("/", router)

	err := http.ListenAndServe( fmt.Sprintf("%s:%s", host, port), nil); if err != nil {
		log.Fatalln(err)

	}
}
