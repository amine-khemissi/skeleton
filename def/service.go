package def

import (
	"context"

	"github.com/amine-khemissi/skeleton/def/count"
	"github.com/amine-khemissi/skeleton/def/uppercase"
	"github.com/amine-khemissi/skeleton/def/write"
)

type Service interface {
	Uppercase(ctx context.Context, req uppercase.Request) (uppercase.Response, error)
	Count(ctx context.Context, req count.Request) (count.Response, error)
	Write(ctx context.Context, req write.Request) (write.Response, error)
}
