package transportCount

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amine-khemissi/skeleton/backbone/endpointimpl"

	"github.com/amine-khemissi/skeleton/def"
	"github.com/amine-khemissi/skeleton/def/count"
	"github.com/go-kit/kit/endpoint"
)

type ep struct {
}

func NewEndpoint() endpointimpl.EndpointImpl {
	return &ep{}
}
func (e ep) MakeEndpoint(svc interface{}) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(count.Request)
		return svc.(def.Service).Count(ctx, req)
	}
}

func (e ep) DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request count.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func (e ep) EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
func (e ep) GetVerb() string {
	return http.MethodGet
}

func (e ep) GetPath() string {
	return "/count"
}
