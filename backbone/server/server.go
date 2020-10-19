package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/amine-khemissi/skeleton/backbone/transport"

	"github.com/amine-khemissi/skeleton/backbone/endpointimpl"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type Server interface {
	Run()
	Register(impl endpointimpl.EndpointImpl)
}

func New(svc interface{}) Server {
	s := &srv{
		r:   mux.NewRouter(),
		svc: svc,
	}
	http.Handle("/", s.r)
	return s
}

type srv struct {
	r   *mux.Router
	svc interface{}
}

func (s *srv) Register(impl endpointimpl.EndpointImpl) {
	implHandler := httptransport.NewServer(
		impl.MakeEndpoint(s.svc),
		transport.DecodeRequest(impl.GetRequest()),
		func(_ context.Context, w http.ResponseWriter, response interface{}) error {
			return json.NewEncoder(w).Encode(response)
		},
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			bts, _ := json.Marshal(err)
			w.Write(bts)
		}),
	)
	s.r.Methods(impl.GetVerb()).Path(impl.GetPath()).Handler(implHandler)
}
func (s *srv) Run() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
