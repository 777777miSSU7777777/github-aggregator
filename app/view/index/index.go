package index

import (
	"net/http"
	"log"
	"github-aggregator-web-closed/app/view/render"
	"github-aggregator-web-closed/app/entity"
	"github-aggregator-web-closed/app/util/http/cookie"
	"github-aggregator-web-closed/app/query"
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