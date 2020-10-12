package transportUpperCase

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amine-khemissi/skeleton/def"
	"github.com/amine-khemissi/skeleton/def/uppercase"
	"github.com/go-kit/kit/endpoint"
)

func MakeEndpoint(svc def.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercase.Request)
		return svc.Uppercase(ctx, req)
	}
}
func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercase.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func GetVerb() string {
	return http.MethodGet
}

func GetPath() string {
	return "/uppercase"
}
