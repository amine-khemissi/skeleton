package endpointimpl

import (
	"github.com/go-kit/kit/endpoint"
)

type EndpointImpl interface {
	MakeEndpoint(svc interface{}) endpoint.Endpoint
	GetVerb() string
	GetPath() string
	GetRequest() interface{}
}
