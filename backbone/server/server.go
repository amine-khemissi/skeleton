package server

import (
	"log"
	"net/http"

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
		impl.DecodeRequest,
		impl.EncodeResponse,
	)
	s.r.Methods(impl.GetVerb()).Path(impl.GetPath()).Handler(implHandler)
}
func (s *srv) Run() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
