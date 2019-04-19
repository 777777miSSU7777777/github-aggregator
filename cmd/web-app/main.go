package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/tokenservice"

	"github.com/777777miSSU7777777/github-aggregator/internal/api"
	"github.com/777777miSSU7777777/github-aggregator/internal/view"
	"github.com/777777miSSU7777777/github-aggregator/internal/view/index"
	"github.com/777777miSSU7777777/github-aggregator/internal/view/login"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logutil"

	"github.com/gorilla/mux"
)

var host string
var port string
var duration string
var algorithm string
var key string
var iv string

const STATIC_DIR = "/web/static/"

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "Defines host ip")
	flag.StringVar(&host, "h", "0.0.0.0", "Defines host ip")
	flag.StringVar(&port, "port", "8080", "Defines host port")
	flag.StringVar(&port, "p", "8080", "Defines host port")
	flag.Parse()
	view.SetTemplates(template.Must(template.ParseGlob("web/templates/*.gohtml")))
	logutil.SetProjectName("github-aggregator")
	tokenservice.TryLoadToken()
}

func main() {
	router := mux.NewRouter()

	router.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))

	router.HandleFunc("/", index.Render).Methods("GET")
	router.HandleFunc("/login", login.Render).Methods("GET")

	router.HandleFunc("/auth", api.Auth).Methods("POST")
	router.HandleFunc("/logout", api.Logout).Methods("POST")

	router.HandleFunc("/profile", api.Profile).Methods("GET")
	router.HandleFunc("/scopes", api.Scopes).Methods("GET")
	router.HandleFunc("/orgs", api.Orgs).Methods("GET")

	http.Handle("/", router)

	log.Info.Printf("Server started on %s:%s", host, port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	if err != nil {
		log.Error.Fatalln(err)
	}

}
