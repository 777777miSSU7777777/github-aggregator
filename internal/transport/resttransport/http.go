package resttransport

import (
	"github.com/777777miSSU7777777/github-aggregator/internal/encoding"
	"github.com/777777miSSU7777777/github-aggregator/internal/endpoints"
	"github.com/777777miSSU7777777/github-aggregator/internal/service/restservice"

	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeCurrentUserHandler(svc restservice.RESTService) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.MakeCurrentUserEndpoint(svc),
		encoding.DecodeCurrentUser,
		encoding.EncodeResponse,
	)
}

func MakeTokenScopesHandler(svc restservice.RESTService) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.MakeTokenScopesEndpoint(svc),
		encoding.DecodeTokenScopes,
		encoding.EncodeResponse,
	)
}

func MakeUserOrgsHandler(svc restservice.RESTService) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.MakeUserOrgsEndpoint(svc),
		encoding.DecodeUserOrgs,
		encoding.EncodeResponse,
	)
}

func MakeFilteredPullsHandler(svc restservice.RESTService) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.MakeFilteredPullsEndpoint(svc),
		encoding.DecodeFilteredPulls,
		encoding.EncodeResponse,
	)
}
