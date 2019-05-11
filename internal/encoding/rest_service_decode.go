package encoding

import (
	"context"
	"encoding/json"

	"github.com/777777miSSU7777777/github-aggregator/internal/request"
)

func decodeCurrentUser(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.CurrentUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeTokenScopes(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.TokenScopesRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeUserOrgs(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.UserOrgsRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeFilteredPulls(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.FilteredPullsReq

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}
