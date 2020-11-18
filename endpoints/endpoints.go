package endpoints

import (
	"context"

	read2 "github.com/amine-khemissi/skeleton/endpoints/read"

	"github.com/amine-khemissi/skeleton/def/read"

	"github.com/amine-khemissi/skeleton/def/write"
	write2 "github.com/amine-khemissi/skeleton/endpoints/write"

	"github.com/amine-khemissi/skeleton/backbone/db"
	"github.com/amine-khemissi/skeleton/backbone/db/mongo"

	"github.com/amine-khemissi/skeleton/def"
	"github.com/amine-khemissi/skeleton/def/count"
	"github.com/amine-khemissi/skeleton/def/uppercase"
	count2 "github.com/amine-khemissi/skeleton/endpoints/count"
	upperCase2 "github.com/amine-khemissi/skeleton/endpoints/uppercase"
)

type stringService struct {
	DBInstance db.DB
}

//todo : modify or add endpoint that select or insert on the DB in order to test it
func (s *stringService) Uppercase(ctx context.Context, req uppercase.Request) (uppercase.Response, error) {
	return upperCase2.Uppercase(ctx, req)
}
func (s *stringService) Count(ctx context.Context, req count.Request) (count.Response, error) {
	return count2.Count(ctx, req)
}

func (s *stringService) Write(ctx context.Context, req write.Request) (write.Response, error) {
	return write2.Write(ctx, s.DBInstance, req)
}

func (s *stringService) Read(ctx context.Context, req read.Request) (read.Response, error) {
	return read2.Read(ctx, s.DBInstance, req)
}

func New(ctx context.Context) def.Service {
	return &stringService{
		DBInstance: mongo.New(ctx),
	}
}
