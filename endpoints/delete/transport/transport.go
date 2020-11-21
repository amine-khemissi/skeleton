package transportdelete

import (
	"context"
	"net/http"

	"github.com/amine-khemissi/skeleton/backbone/endpointimpl"
	"github.com/amine-khemissi/skeleton/def"
	"github.com/amine-khemissi/skeleton/def/delete"
	"github.com/go-kit/kit/endpoint"
)

type ep struct {
}

func (e ep) GetRequest() interface{} {
	return &delete.Request{}
}

func NewEndpoint() endpointimpl.EndpointImpl {
	return &ep{}
}

func (e ep) MakeEndpoint(svc interface{}) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*delete.Request)
		return svc.(def.Service).Delete(ctx, *req)
	}
}

func (e ep) GetVerb() string {
	return http.MethodDelete
}

func (e ep) GetPath() string {
	return "/person/{clientID}"
}
