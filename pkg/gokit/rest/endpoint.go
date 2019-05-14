package rest

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeCurrentUserEndpoint(svc RESTService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		_ = request.(currentUserRequest)
		val := svc.CurrentUser()

		return currentUserResponse{val}, nil
	}
}

func MakeTokenScopesEndpoint(svc RESTService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		_ = request.(tokenScopesRequest)
		val, err := svc.TokenScopes()

		if err != nil {
			return tokenScopesResponse{val, err.Error()}, err
		}

		return tokenScopesResponse{val, ""}, nil
	}
}

func MakeUserOrgsEndpoint(svc RESTService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		_ = request.(userOrgsRequest)
		val := svc.UserOrgs()

		return userOrgsResponse{val}, nil
	}
}

func MakeFilteredPullsEndpoint(svc RESTService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(filteredPullsReq)
		val, err := svc.FilteredPulls(req.Filter, req.SelectedOrgs)

		if err != nil {
			return filteredPullsResponse{val, err.Error()}, err
		}

		return filteredPullsResponse{val, ""}, nil
	}
}
