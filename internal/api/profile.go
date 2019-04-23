package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/tokenservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/factory/profilefactory"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/query"
)

// Profile returns repsonse with current profile info in json format.
func Profile(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	tkn := tokenservice.GetToken()

	userBytes, err := query.GetDataSource().GetUser(context.Background(), tkn)

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
