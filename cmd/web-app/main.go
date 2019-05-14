package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/777777miSSU7777777/github-aggregator/pkg/session"

	"github.com/777777miSSU7777777/github-aggregator/internal/api"
	"github.com/777777miSSU7777777/github-aggregator/internal/gokit/rest"
	"github.com/777777miSSU7777777/github-aggregator/internal/view"
	"github.com/777777miSSU7777777/github-aggregator/internal/view/index"
	"github.com/777777miSSU7777777/github-aggregator/internal/view/login"
	"github.com/777777miSSU7777777/github-aggregator/internal/view/pulls"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/datasrcfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/query"
	"github.com/777777miSSU7777777/github-aggregator/pkg/token"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var host string
var port string
var dataSrc string
var logger *log.Logger

const STATIC_DIR = "/web/static/"
const timeFormat = "2006-01-02 15:04:05"

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "Defines host ip")
	flag.StringVar(&host, "h", "0.0.0.0", "Defines host ip")
	flag.StringVar(&port, "port", "8080", "Defines host port")
	flag.StringVar(&port, "p", "8080", "Defines host port")
	flag.StringVar(&dataSrc, "data-source", "rest-api", "Defines data source")
	flag.Parse()

	logger = log.New()
	jsonFormatter := &log.JSONFormatter{}
	jsonFormatter.TimestampFormat = timeFormat
	logger.SetFormatter(jsonFormatter)
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)

	view.SetTemplates(template.Must(template.ParseGlob("web/templates/*.gohtml")))
	query.SetDataSource(datasrcfactory.New(dataSrc))
	err := token.GetTokenService().TryLoadToken()

	if err != nil {
		logger.Warnln(err)
	} else {
		logger.Infoln("Token initalized from .token file")
	}

	token := token.GetTokenService().GetToken()
	if token != "" {
		err = session.GetSessionService().StartSession(token)

		if err != nil {
			logger.Warnln(err)
		}
	}

	api.SetLogger(logger)
	view.SetLogger(logger)
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

	restService := rest.NewRestServiceImpl()
	restService = rest.WrapLoggingMiddleware(restService, logger)

	currentUserHandler := rest.MakeCurrentUserHandler(restService)
	tokenScopesHandler := rest.MakeTokenScopesHandler(restService)
	userOrgsHandler := rest.MakeUserOrgsHandler(restService)
	filteredPullsHandler := rest.MakeFilteredPullsHandler(restService)

	apiRouter.Handle("/profile", currentUserHandler).Methods("GET")
	apiRouter.Handle("/scopes", tokenScopesHandler).Methods("GET")
	apiRouter.Handle("/orgs", userOrgsHandler).Methods("GET")
	apiRouter.Handle("/pulls", filteredPullsHandler).Methods("GET")

	http.Handle("/", router)

	logger.Infof("Server started listening on %s:%s", host, port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	if err != nil {
		logger.Fatalln(err)
	}

}
