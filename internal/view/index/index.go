package index

import (
	"log"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator-web/app/entity"
	"github.com/777777miSSU7777777/github-aggregator-web/app/query"
	"github.com/777777miSSU7777777/github-aggregator-web/app/util/http/cookie"
	"github.com/777777miSSU7777777/github-aggregator-web/app/view/render"
)

//Render renders index page.
func Render(rw http.ResponseWriter, req *http.Request) {
	session := entity.Session{}

	accessToken, err := cookie.GetCookieValue(req, "access_token")

	if err != nil {
		log.Println(err.Error())
	}

	session.Authorized = accessToken != ""
	if session.Authorized {
		session.Username, _ = query.GetUsername(req)
		session.AvatarURL, _ = query.GetAvatarURL(req)
		session.ProfileURL, _ = query.GetProfileURL(req)
		session.Scopes = query.GetScopes(req)
	}

	err = render.GetTemplates().ExecuteTemplate(rw, "index.gohtml", session)

	if err != nil {
		log.Println(err.Error())
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	}

}
