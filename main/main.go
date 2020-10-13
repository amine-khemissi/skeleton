package main

import (
	"github.com/amine-khemissi/skeleton/backbone/server"
	"github.com/amine-khemissi/skeleton/endpoints"
	"github.com/amine-khemissi/skeleton/endpoints/count/transportCount"
	"github.com/amine-khemissi/skeleton/endpoints/uppercase/transportUpperCase"
)

/* todo :
- add middle ware chaining
- add logger chain
- add doc generation
*/

func main() {
	srv := server.New(endpoints.New())
	srv.Register(transportUpperCase.NewEndpoint())
	srv.Register(transportCount.NewEndpoint())
	srv.Run()
}
