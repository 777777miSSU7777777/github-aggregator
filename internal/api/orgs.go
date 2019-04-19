package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/tokenservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/orgsfactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/query"
)

// Orgs returns response with current organizations info in json format.
func Orgs(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	tkn := tokenservice.GetToken()

	orgsBytes, err := query.GetDataSource().GetOrgs(context.Background(), tkn)

	if err != nil {
		log.Error.Println(err)
	}

	orgs, err := orgsfactory.New(orgsBytes)

	if err != nil {
		log.Error.Println(err)
	}

	err = json.NewEncoder(rw).Encode(orgs)

	if err != nil {
		log.Error.Println(err)
	}
}
