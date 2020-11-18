package doc

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func (g *generator) GetRequest() interface{} {
	return &struct {
	}{}
}

func (g *generator) MakeEndpoint(svc interface{}) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return g.doc, nil
	}
}

func (g *generator) GetVerb() string {
	return http.MethodGet
}

func (g *generator) GetPath() string {
	return "/_internal/doc"
}
