package main

import (
	"log"
	"net/http"

	"github.com/amine-khemissi/skeleton/endpoints"
	"github.com/amine-khemissi/skeleton/endpoints/count/transportCount"
	"github.com/amine-khemissi/skeleton/endpoints/uppercase/transportUpperCase"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := endpoints.New()

	uppercaseHandler := httptransport.NewServer(
		transportUpperCase.MakeEndpoint(svc),
		transportUpperCase.DecodeRequest,
		transportUpperCase.EncodeResponse,
	)

	http.Handle(transportUpperCase.GetPath(), uppercaseHandler)

	countHandler := httptransport.NewServer(
		transportCount.MakeEndpoint(svc),
		transportCount.DecodeRequest,
		transportCount.EncodeResponse,
	)

	http.Handle(transportCount.GetPath(), countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
