package transportUpperCase

import (
	"context"
	"net/http"

	"github.com/amine-khemissi/skeleton/backbone/endpointimpl"

	"github.com/amine-khemissi/skeleton/def"
	"github.com/amine-khemissi/skeleton/def/uppercase"
	"github.com/go-kit/kit/endpoint"
)

type ep struct {
}

func (e ep) GetRequest() interface{} {
	return &uppercase.Request{}
}

func NewEndpoint() endpointimpl.EndpointImpl {
	return &ep{}
}

func (e ep) MakeEndpoint(svc interface{}) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*uppercase.Request)
		return svc.(def.Service).Uppercase(ctx, *req)
	}
}

func (e ep) GetVerb() string {
	return http.MethodGet
}

func (e ep) GetPath() string {
	return "/uppercase"
}
