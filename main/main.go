package main

import (
	"log"
	"net/http"

	"github.com/amine-khemissi/skeleton/endpoints"
	"github.com/amine-khemissi/skeleton/endpoints/count/transportCount"
	"github.com/amine-khemissi/skeleton/endpoints/uppercase/transportUpperCase"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

/* todo :
- add gorilla mux
- add middle ware chaining
- add logger chain
- add interface handler
- add doc generation
*/
func main() {
	svc := endpoints.New()
	r := mux.NewRouter()
	http.Handle("/", r)

	uppercaseHandler := httptransport.NewServer(
		transportUpperCase.MakeEndpoint(svc),
		transportUpperCase.DecodeRequest,
		transportUpperCase.EncodeResponse,
	)
	r.Path(transportUpperCase.GetPath()).Methods(transportUpperCase.GetVerb()).Handler(uppercaseHandler)

	countHandler := httptransport.NewServer(
		transportCount.MakeEndpoint(svc),
		transportCount.DecodeRequest,
		transportCount.EncodeResponse,
	)
	r.Path(transportCount.GetPath()).Methods(transportCount.GetVerb()).Handler(countHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
