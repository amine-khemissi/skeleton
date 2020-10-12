package transportCount

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amine-khemissi/skeleton/def"
	"github.com/amine-khemissi/skeleton/def/count"
	"github.com/go-kit/kit/endpoint"
)

func MakeEndpoint(svc def.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(count.Request)
		return svc.Count(ctx, req)
	}
}

func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request count.Request
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
	return "/count"
}
