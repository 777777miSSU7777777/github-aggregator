package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/query"
	"github.com/777777miSSU7777777/github-aggregator/pkg/token"
)

// Scopes returns response with current scopes for provided token in json format.
func Scopes(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	tkn := token.GetTokenService().GetToken()

	scopesArr, err := query.GetDataSource().GetScopes(context.Background(), tkn)

	if err != nil {
		log.Error.Println(err)
	}

	scopesMap := map[string]entity.Scope{"scopes": strings.Join(scopesArr, ",")}

	err = json.NewEncoder(rw).Encode(scopesMap)

	if err != nil {
		log.Error.Println(err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	} else {
		log.Info.Printf("Sent %s", scopesMap)
	}
}
