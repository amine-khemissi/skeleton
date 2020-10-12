package def

import (
	"context"

	"github.com/amine-khemissi/skeleton/def/count"
	"github.com/amine-khemissi/skeleton/def/uppercase"
)

type Service interface {
	Uppercase(ctx context.Context, req uppercase.Request) (uppercase.Response, error)
	Count(ctx context.Context, req count.Request) (count.Response, error)
}
