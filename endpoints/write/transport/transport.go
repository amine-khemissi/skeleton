package transportwrite

import (
	"context"
	"net/http"

	"github.com/amine-khemissi/skeleton/def/write"

	"github.com/amine-khemissi/skeleton/backbone/endpointimpl"

	"github.com/amine-khemissi/skeleton/def"
	"github.com/go-kit/kit/endpoint"
)

type ep struct {
}

func (e ep) GetRequest() interface{} {
	return &write.Request{}
}

func NewEndpoint() endpointimpl.EndpointImpl {
	return &ep{}
}

func (e ep) MakeEndpoint(svc interface{}) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*write.Request)
		return svc.(def.Service).Write(ctx, *req)
	}
}

func (e ep) GetVerb() string {
	return http.MethodPost
}

func (e ep) GetPath() string {
	return "/person"
}
