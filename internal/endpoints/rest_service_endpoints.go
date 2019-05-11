package endpoints

import (
	"context"

	servicereq "github.com/777777miSSU7777777/github-aggregator/internal/request"
	serviceresp "github.com/777777miSSU7777777/github-aggregator/internal/response"
	"github.com/777777miSSU7777777/github-aggregator/internal/service/restservice"

	"github.com/go-kit/kit/endpoint"
)

func makeCurrentUserEndpoint(svc restservice.RESTService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		_ = request.(servicereq.CurrentUserRequest)
		val := svc.CurrentUser()

		return serviceresp.CurrentUserResponse{val}, nil
	}
}

func makeTokenScopesEndpoint(svc restservice.RESTService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		_ = request.(servicereq.TokenScopesRequest)
		val, err := svc.TokenScopes()

		if err != nil {
			return serviceresp.TokenScopesResponse{val, err.Error()}, err
		}

		return serviceresp.TokenScopesResponse{val, ""}, nil
	}
}

func makeUserOrgsEndpoint(svc restservice.RESTService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		_ = request.(servicereq.UserOrgsRequest)
		val := svc.UserOrgs()

		return serviceresp.UserOrgsResponse{val}, nil
	}
}

func makeFilteredPullsEndpoint(svc restservice.RESTService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(servicereq.FilteredPullsReq)
		val, err := svc.FilteredPulls(req.Filter, req.SelectedOrgs)

		if err != nil {
			return serviceresp.FilteredPullsResponse{val, err.Error()}, err
		}

		return serviceresp.FilteredPullsResponse{val, ""}, nil
	}
}
