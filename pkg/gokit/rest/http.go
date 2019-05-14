package rest

import (
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeCurrentUserHandler(svc RESTService) *httptransport.Server {
	return httptransport.NewServer(
		MakeCurrentUserEndpoint(svc),
		DecodeCurrentUser,
		EncodeResponse,
	)
}

func MakeTokenScopesHandler(svc RESTService) *httptransport.Server {
	return httptransport.NewServer(
		MakeTokenScopesEndpoint(svc),
		DecodeTokenScopes,
		EncodeResponse,
	)
}

func MakeUserOrgsHandler(svc RESTService) *httptransport.Server {
	return httptransport.NewServer(
		MakeUserOrgsEndpoint(svc),
		DecodeUserOrgs,
		EncodeResponse,
	)
}

func MakeFilteredPullsHandler(svc RESTService) *httptransport.Server {
	return httptransport.NewServer(
		MakeFilteredPullsEndpoint(svc),
		DecodeFilteredPulls,
		EncodeResponse,
	)
}
