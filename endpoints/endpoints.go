package endpoints

import (
	"context"
	"github.com/amine-khemissi/skeleton/def"
	"github.com/amine-khemissi/skeleton/def/count"
	"github.com/amine-khemissi/skeleton/def/upperCase"
	upperCase2 "github.com/amine-khemissi/skeleton/endpoints/upperCase"
)

type stringService struct {
}

func (stringService) Uppercase(ctx context.Context, req upperCase.Request) (upperCase.Response, error) {

	return upperCase2.UpperCase(ctx, req)
}
func (stringService) Count(ctx context.Context, req count.Request) (count.Response, error) {
	return count.Count(ctx, req)
}

func New() def.Service {
	return stringService{}
}
