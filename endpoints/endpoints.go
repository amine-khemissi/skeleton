package endpoints

import (
	"context"

	"github.com/amine-khemissi/skeleton/def"
	"github.com/amine-khemissi/skeleton/def/count"
	"github.com/amine-khemissi/skeleton/def/uppercase"
	count2 "github.com/amine-khemissi/skeleton/endpoints/count"
	upperCase2 "github.com/amine-khemissi/skeleton/endpoints/uppercase"
)

type stringService struct {
}

func (stringService) Uppercase(ctx context.Context, req uppercase.Request) (uppercase.Response, error) {
	return upperCase2.Uppercase(ctx, req)
}
func (stringService) Count(ctx context.Context, req count.Request) (count.Response, error) {
	return count2.Count(ctx, req)
}

func New() def.Service {
	return stringService{}
}
