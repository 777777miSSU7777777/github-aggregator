package api

import (
	"net/http"

	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/webtokenservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/profilefactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/query"
)

// Profile returns repsonse with current profile info in json format.
func Profile(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	tkn, err := webtokenservice.GetToken(req)

	if err != nil {
		log.Warning.Println(err)
	}

	userBytes, err := query.GetUser(tkn)

	if err != nil {
		log.Error.Println(err)
	}

	profile, err := profilefactory.New(userBytes)

	if err != nil {
		log.Error.Println(err)
	}

	err = json.NewEncoder(rw).Encode(profile)

	if err != nil {
		log.Error.Println(err)
	}
}
