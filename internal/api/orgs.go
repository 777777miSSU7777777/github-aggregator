package api

import (
	"net/http"

	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/webtokenservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/orgsfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/query"
)

// Orgs returns response with current organizations info in json format.
func Orgs(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	tkn, err := webtokenservice.GetToken(req)

	if err != nil {
		log.Warning.Println(err)
	}

	orgsBytes, err := query.GetOrgs(tkn)

	if err != nil {
		log.Error.Println(err)
	}

	orgs, err := orgsfactory.New(orgsBytes)

	if err != nil {
		log.Error.Println(err)
	}

	json.NewEncoder(rw).Encode(orgs)
}
