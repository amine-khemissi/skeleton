package config

import (
	"context"
	"net/http"

	"github.com/amine-khemissi/skeleton/backbone/endpointimpl"

	"github.com/go-kit/kit/endpoint"
)

type ep struct {
}

type Request struct {
}

func (e ep) GetRequest() interface{} {
	return nil
}

func NewEndpoint() endpointimpl.EndpointImpl {
	return &ep{}
}

func (e ep) MakeEndpoint(svc interface{}) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return gConf.(*config).Content, nil
	}
}

func (e ep) GetVerb() string {
	return http.MethodGet
}

func (e ep) GetPath() string {
	return "/_internal/config"
}
