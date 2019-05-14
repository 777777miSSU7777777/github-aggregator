package rest

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	FILTER_PARAM        = "filter"
	SELECTED_ORGS_PARAM = "selected_orgs"
)

func DecodeCurrentUser(_ context.Context, r *http.Request) (interface{}, error) {
	return currentUserRequest{}, nil
}

func DecodeTokenScopes(_ context.Context, r *http.Request) (interface{}, error) {
	return tokenScopesRequest{}, nil
}

func DecodeUserOrgs(_ context.Context, r *http.Request) (interface{}, error) {
	return userOrgsRequest{}, nil
}

func DecodeFilteredPulls(_ context.Context, r *http.Request) (interface{}, error) {
	filter := r.FormValue(FILTER_PARAM)

	var selectedOrgs []string

	err := json.Unmarshal([]byte(r.FormValue(SELECTED_ORGS_PARAM)), &selectedOrgs)

	if err != nil {
		return nil, err
	}

	return filteredPullsReq{filter, selectedOrgs}, nil
}

func EncodeResponse(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(rw).Encode(response)
}
