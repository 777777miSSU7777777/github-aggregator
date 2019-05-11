package encoding

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/request"
)

const (
	FILTER_PARAM        = "filter"
	SELECTED_ORGS_PARAM = "selected_orgs"
)

func DecodeCurrentUser(_ context.Context, r *http.Request) (interface{}, error) {
	return request.CurrentUserRequest{}, nil
}

func DecodeTokenScopes(_ context.Context, r *http.Request) (interface{}, error) {
	return request.TokenScopesRequest{}, nil
}

func DecodeUserOrgs(_ context.Context, r *http.Request) (interface{}, error) {
	return request.UserOrgsRequest{}, nil
}

func DecodeFilteredPulls(_ context.Context, r *http.Request) (interface{}, error) {
	filter := r.FormValue(FILTER_PARAM)

	var selectedOrgs []string

	err := json.Unmarshal([]byte(r.FormValue(SELECTED_ORGS_PARAM)), &selectedOrgs)

	if err != nil {
		return nil, err
	}

	return request.FilteredPullsReq{filter, selectedOrgs}, nil
}
