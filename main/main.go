package main

import (
	"context"

	transportupdate "github.com/amine-khemissi/skeleton/endpoints/update/transport"

	transportdelete "github.com/amine-khemissi/skeleton/endpoints/delete/transport"

	transportread "github.com/amine-khemissi/skeleton/endpoints/read/transport"

	"github.com/amine-khemissi/skeleton/backbone/logger"
	"github.com/amine-khemissi/skeleton/backbone/server"
	"github.com/amine-khemissi/skeleton/endpoints"
	transportwrite "github.com/amine-khemissi/skeleton/endpoints/write/transport"
)

/* todo :
- check other todos
- add middleware chaining
- add doc generation
- db driver : check other methods are working
- err response encoder to be implemented
- add ContentType int the response and the request and change marshalling process accordingly
*/

func main() {
	ctx := context.Background()
	logger.Instance().Debug(ctx, "starting service")
	srv := server.New(endpoints.New(ctx))
	srv.Register(transportwrite.NewEndpoint())
	srv.Register(transportread.NewEndpoint())
	srv.Register(transportdelete.NewEndpoint())
	srv.Register(transportupdate.NewEndpoint())
	srv.Run()
}
