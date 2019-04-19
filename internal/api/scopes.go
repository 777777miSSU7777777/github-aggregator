package api

import (
	"net/http"
	"strings"

	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/tokenservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/query"
)

// Scopes returns response with current scopes for provided token in json format.
func Scopes(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	tkn := tokenservice.GetToken()

	scopesArr, err := query.GetScopes(tkn)

	if err != nil {
		log.Error.Println(err)
	}

	scopesMap := map[string]string{"scopes": strings.Join(scopesArr, ",")}

	err = json.NewEncoder(rw).Encode(scopesMap)

	if err != nil {
		log.Error.Println(err)
	}
}
