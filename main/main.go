package main

import (
	"context"

	"github.com/amine-khemissi/skeleton/backbone/logger"
	"github.com/amine-khemissi/skeleton/backbone/server"
	"github.com/amine-khemissi/skeleton/endpoints"
	"github.com/amine-khemissi/skeleton/endpoints/count/transportCount"
	"github.com/amine-khemissi/skeleton/endpoints/uppercase/transportUpperCase"
)

/* todo :
- add in request context request ID and print it in logger
- check other todos
- add middleware chaining
- add doc generation
- add db driver
*/

func main() {
	logger.Instance().Debug(context.Background(), "test one")
	srv := server.New(endpoints.New())
	srv.Register(transportUpperCase.NewEndpoint())
	srv.Register(transportCount.NewEndpoint())
	srv.Run()
}
