package index

import (
	"net/http"
	"log"
	"github.com/777777miSSU7777777/github-aggregator-web/app/view/render"
	"github.com/777777miSSU7777777/github-aggregator-web/app/entity"
	"github.com/777777miSSU7777777/github-aggregator-web/app/util/http/cookie"
	"github.com/777777miSSU7777777/github-aggregator-web/app/query"
)

func Render(rw http.ResponseWriter, req *http.Request){
	session := entity.Session{}
	
	access_token, err := cookie.GetCookieValue(req, "access_token")

	if err != nil {
		log.Println(err.Error())
	}

	session.Authorized = access_token != ""; if session.Authorized {
		session.Username, _ = query.GetUsername(req)
		session.AvatarUrl, _  = query.GetAvatarUrl(req)
		session.ProfileUrl, _ = query.GetProfileUrl(req)
		session.Scopes = query.GetScopes(req)
	}
	
	err = render.GetTemplates().ExecuteTemplate(rw,"index.gohtml", session)

	if err != nil {
		log.Println(err.Error())
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	}
	
}