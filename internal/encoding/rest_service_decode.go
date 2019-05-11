package encoding

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/777777miSSU7777777/github-aggregator/internal/request"
)

func DecodeCurrentUser(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.CurrentUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func DecodeTokenScopes(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.TokenScopesRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func DecodeUserOrgs(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.UserOrgsRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func DecodeFilteredPulls(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.FilteredPullsReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}
