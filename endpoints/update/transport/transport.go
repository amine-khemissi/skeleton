package transportupdate

import (
	"context"
	"net/http"

	"github.com/amine-khemissi/skeleton/def/update"

	"github.com/amine-khemissi/skeleton/backbone/endpointimpl"

	"github.com/amine-khemissi/skeleton/def"
	"github.com/go-kit/kit/endpoint"
)

type ep struct {
}

func (e ep) GetRequest() interface{} {
	return &update.Request{}
}

func NewEndpoint() endpointimpl.EndpointImpl {
	return &ep{}
}

func (e ep) MakeEndpoint(svc interface{}) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*update.Request)
		return svc.(def.Service).Update(ctx, *req)
	}
}

func (e ep) GetVerb() string {
	return http.MethodPatch
}

func (e ep) GetPath() string {
	return "/person/{clientID}"
}
