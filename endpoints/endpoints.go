package endpoints

import (
	"context"

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
func (stringService) Uppercase(ctx context.Context, req uppercase.Request) (uppercase.Response, error) {
	return upperCase2.Uppercase(ctx, req)
}
func (stringService) Count(ctx context.Context, req count.Request) (count.Response, error) {
	return count2.Count(ctx, req)
}

func New(ctx context.Context) def.Service {
	return stringService{
		DBInstance: mongo.New(ctx),
	}
}
