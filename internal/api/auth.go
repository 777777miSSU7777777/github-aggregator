package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/constants"
	"github.com/777777miSSU7777777/github-aggregator/internal/security/webtokenservice"
)


func Auth(rw http.ResponseWriter, req *http.Request) {
	tkn := req.FormValue(constants.AccessToken)
	resp, err := http.Get( fmt.Sprintf("%s%s?%s%s", constants.GHApiURL, constants.User, constants.AccessTokenParam, tkn))

	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode == 200 {
		err = webtokenservice.SaveToken(rw, tkn); if err != nil {
			log.Println(err)
		}
		log.Println("Authentication is successful")
	} else if resp.StatusCode == 401 {
		log.Println("Authentication is failed")
	}

	http.Redirect(rw, req, "/", 301)
}
