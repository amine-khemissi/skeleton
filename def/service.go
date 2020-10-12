package def

import (
	"context"
	"github.com/amine-khemissi/skeleton/def/count"
	"github.com/amine-khemissi/skeleton/def/upperCase"
)

type Service interface {
	Uppercase(ctx context.Context, req upperCase.Request) (upperCase.Response, error)
	Count(ctx context.Context, req count.Request) (count.Response, error)
}
