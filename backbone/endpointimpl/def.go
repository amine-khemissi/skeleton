package endpointimpl

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type EndpointImpl interface {
	MakeEndpoint(svc interface{}) endpoint.Endpoint
	DecodeRequest(_ context.Context, r *http.Request) (interface{}, error)
	EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error
	GetVerb() string
	GetPath() string
}
