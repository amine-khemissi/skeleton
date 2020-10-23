package main

import (
	"context"

	"github.com/amine-khemissi/skeleton/backbone/logger"
	"github.com/amine-khemissi/skeleton/backbone/server"
	"github.com/amine-khemissi/skeleton/endpoints"
	"github.com/amine-khemissi/skeleton/endpoints/count/transportCount"
	"github.com/amine-khemissi/skeleton/endpoints/uppercase/transportUpperCase"
	transportwrite "github.com/amine-khemissi/skeleton/endpoints/write/transport"
)

/* todo :
- add in request context request ID and print it in logger
- check other todos
- add middleware chaining
- add doc generation
- db driver : check other methods are working
- err response encoder to be implemented
*/

func main() {
	ctx := context.Background()
	logger.Instance().Debug(ctx, "starting service")
	srv := server.New(endpoints.New(ctx))
	srv.Register(transportUpperCase.NewEndpoint())
	srv.Register(transportCount.NewEndpoint())
	srv.Register(transportwrite.NewEndpoint())
	srv.Run()
}
