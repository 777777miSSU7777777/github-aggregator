package api

import (
	"net/http"
	"log"
	"time"
	"github.com/777777miSSU7777777/github-aggregator-web/app/util/http/cookie"
)

func Auth(rw http.ResponseWriter, req *http.Request){
	if req.Method == http.MethodPost{
		access_token := req.FormValue("access_token")
		log.Printf("Token: %s",access_token)
		resp, err := http.Get("https://api.github.com/user?access_token=" + access_token)

		if err != nil {
			log.Println(err.Error())
		}

		if (resp.StatusCode == 200){
			cookie.SaveCookie(rw, "access_token", access_token, time.Hour)
			log.Println("Authentication is successful")
		} else if (resp.StatusCode == 401){
			log.Println("Authentication is failed")
		}

		http.Redirect(rw, req,"/", 301)
	}
}