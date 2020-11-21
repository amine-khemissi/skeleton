package endpoints

import (
	"context"

	"github.com/amine-khemissi/skeleton/backbone/db"
	"github.com/amine-khemissi/skeleton/backbone/db/mongo"
	"github.com/amine-khemissi/skeleton/def"
	"github.com/amine-khemissi/skeleton/def/count"
	"github.com/amine-khemissi/skeleton/def/delete"
	"github.com/amine-khemissi/skeleton/def/read"
	"github.com/amine-khemissi/skeleton/def/uppercase"
	"github.com/amine-khemissi/skeleton/def/write"
	count2 "github.com/amine-khemissi/skeleton/endpoints/count"
	delete2 "github.com/amine-khemissi/skeleton/endpoints/delete"
	read2 "github.com/amine-khemissi/skeleton/endpoints/read"
	upperCase2 "github.com/amine-khemissi/skeleton/endpoints/uppercase"
	write2 "github.com/amine-khemissi/skeleton/endpoints/write"
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

func (s *stringService) Delete(ctx context.Context, req delete.Request) (delete.Response, error) {
	return delete2.Delete(ctx, s.DBInstance, req)
}

func New(ctx context.Context) def.Service {
	return &stringService{
		DBInstance: mongo.New(ctx),
	}
}
