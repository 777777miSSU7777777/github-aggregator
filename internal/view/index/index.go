package index

import (
	"log"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/query"
	"github.com/777777miSSU7777777/github-aggregator/pkg/http/cookieutil"
	"github.com/777777miSSU7777777/github-aggregator/internal/view"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/profilefactory"
)


func Render(rw http.ResponseWriter, req *http.Request) {
	tkn, err := cookieutil.GetCookieValue(req, "access_token"); if err != nil {
		log.Println(err)
	}

	if tkn != "" {
		userBytes, err := query.QueryUser(tkn); if err != nil {
			log.Println(err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
		}

		scopes, err := query.QueryScopes(tkn); if err != nil {
			log.Println(err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
		}

		profile, err := profilefactory.New(userBytes, scopes); if err != nil {
			log.Println(err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
		}
		
		err = view.GetTemplates().ExecuteTemplate(rw, "index-authorized.gohtml", profile); if err != nil {
			log.Println(err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
		}
	} else {
		err = view.GetTemplates().ExecuteTemplate(rw, "index.gohtml", nil); if err != nil {
			log.Println(err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
		}
	}
}
