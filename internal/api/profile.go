package api

import (
	"encoding/json"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/session"
)

// Profile returns repsonse with current profile info in json format.
func Profile(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	profile := session.GetSessionService().GetSession().GetCurrentUser()

	err := json.NewEncoder(rw).Encode(profile)

	if err != nil {
		log.Error.Println(err)
	}
}
