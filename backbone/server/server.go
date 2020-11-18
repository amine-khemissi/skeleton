package server

import (
	"flag"
	"log"
	"net/http"

	"github.com/amine-khemissi/skeleton/backbone/config"
	"github.com/amine-khemissi/skeleton/backbone/endpointimpl"
	"github.com/amine-khemissi/skeleton/backbone/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type Server interface {
	Run()
	Register(impl endpointimpl.EndpointImpl)
}

func init() {
	flag.Parse()
}
func New(svc interface{}) Server {
	s := &srv{
		r:   mux.NewRouter(),
		svc: svc,
	}

	http.Handle("/", s.r)
	s.Register(config.NewEndpoint())
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
		genericEncoder,
		httptransport.ServerErrorEncoder(genericErrorEncoder),
		httptransport.ServerBefore(addContextID),
	)
	s.r.Methods(impl.GetVerb()).Path(impl.GetPath()).Handler(implHandler)
}

func (s *srv) Run() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
