package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/session"

	"github.com/777777miSSU7777777/github-aggregator/internal/api"
	"github.com/777777miSSU7777777/github-aggregator/internal/service/restservice"
	"github.com/777777miSSU7777777/github-aggregator/internal/transport/resttransport"
	"github.com/777777miSSU7777777/github-aggregator/internal/view"
	"github.com/777777miSSU7777777/github-aggregator/internal/view/index"
	"github.com/777777miSSU7777777/github-aggregator/internal/view/login"
	"github.com/777777miSSU7777777/github-aggregator/internal/view/pulls"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/datasrcfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log/logutil"
	"github.com/777777miSSU7777777/github-aggregator/pkg/query"
	"github.com/777777miSSU7777777/github-aggregator/pkg/token"

	"github.com/gorilla/mux"
)

var host string
var port string
var dataSrc string

const STATIC_DIR = "/web/static/"

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "Defines host ip")
	flag.StringVar(&host, "h", "0.0.0.0", "Defines host ip")
	flag.StringVar(&port, "port", "8080", "Defines host port")
	flag.StringVar(&port, "p", "8080", "Defines host port")
	flag.StringVar(&dataSrc, "data-source", "rest-api", "Defines data source")
	flag.Parse()
	view.SetTemplates(template.Must(template.ParseGlob("web/templates/*.gohtml")))
	logutil.SetProjectName("github-aggregator")
	query.SetDataSource(datasrcfactory.New(dataSrc))
	token.GetTokenService().TryLoadToken()
	token := token.GetTokenService().GetToken()
	if token != "" {
		session.GetSessionService().StartSession(token)
	}
}

func main() {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()

	router.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))

	router.HandleFunc("/", index.Render).Methods("GET")
	router.HandleFunc("/login", login.Render).Methods("GET")
	router.HandleFunc("/pulls", pulls.Render).Methods("GET")

	apiRouter.HandleFunc("/auth", api.Auth).Methods("POST")
	apiRouter.HandleFunc("/logout", api.Logout).Methods("POST")

	restService := restservice.NewRestServiceImpl()

	currentUserHandler := resttransport.MakeCurrentUserHandler(restService)
	tokenScopesHandler := resttransport.MakeTokenScopesHandler(restService)
	userOrgsHandler := resttransport.MakeUserOrgsHandler(restService)
	filteredPullsHandler := resttransport.MakeFilteredPullsHandler(restService)

	apiRouter.Handle("/profile", currentUserHandler).Methods("GET")
	apiRouter.Handle("/scopes", tokenScopesHandler).Methods("GET")
	apiRouter.Handle("/orgs", userOrgsHandler).Methods("GET")
	apiRouter.Handle("/pulls", filteredPullsHandler).Methods("GET")

	http.Handle("/", router)

	log.Info.Printf("Server started on %s:%s", host, port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	if err != nil {
		log.Error.Fatalln(err)
	}

}
