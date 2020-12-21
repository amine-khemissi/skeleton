package endpoints

import (
	"context"

	"github.com/amine-khemissi/skeleton/backbone/db"
	"github.com/amine-khemissi/skeleton/backbone/db/mongo"
	"github.com/amine-khemissi/skeleton/def"
	"github.com/amine-khemissi/skeleton/def/delete"
	"github.com/amine-khemissi/skeleton/def/read"
	"github.com/amine-khemissi/skeleton/def/update"
	"github.com/amine-khemissi/skeleton/def/write"
	delete2 "github.com/amine-khemissi/skeleton/endpoints/delete"
	read2 "github.com/amine-khemissi/skeleton/endpoints/read"
	update2 "github.com/amine-khemissi/skeleton/endpoints/update"
	write2 "github.com/amine-khemissi/skeleton/endpoints/write"
)

type stringService struct {
	DBInstance db.DB
}

func (s *stringService) Write(ctx context.Context, req write.Request) (write.Response, error) {
	return write2.Write(ctx, s.DBInstance, req)
}

func (s *stringService) Update(ctx context.Context, req update.Request) (update.Response, error) {
	return update2.Update(ctx, s.DBInstance, req)
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
