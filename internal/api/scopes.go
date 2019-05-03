package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/777777miSSU7777777/github-aggregator/internal/security/tokenservice"
	"github.com/777777miSSU7777777/github-aggregator/pkg/log"
	"github.com/777777miSSU7777777/github-aggregator/pkg/query"
)

// Scopes returns response with current scopes for provided token in json format.
func Scopes(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	tkn := tokenservice.GetTokenService().GetToken()

	scopesArr, err := query.GetDataSource().GetScopes(context.Background(), tkn)

	if err != nil {
		log.Error.Println(err)
	}

	scopesMap := map[string]string{"scopes": strings.Join(scopesArr, ",")}

	err = json.NewEncoder(rw).Encode(scopesMap)

	if err != nil {
		log.Error.Println(err)
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
	} else {
		log.Info.Printf("Sent %s", scopesMap)
	}
}
