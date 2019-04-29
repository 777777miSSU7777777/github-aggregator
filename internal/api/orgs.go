package api

import (
	"encoding/json"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/session"
)

// Orgs returns response with current organizations info in json format.
func Orgs(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	orgs := session.GetSessionService().GetSession().GetUserOrgs()

	err := json.NewEncoder(rw).Encode(orgs)

	if err != nil {
		log.Error.Println(err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	} else {
		log.Info.Printf("Sent %s", orgs)
	}
}
