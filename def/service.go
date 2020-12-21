package def

import (
	"context"

	"github.com/amine-khemissi/skeleton/def/delete"
	"github.com/amine-khemissi/skeleton/def/read"
	"github.com/amine-khemissi/skeleton/def/update"
	"github.com/amine-khemissi/skeleton/def/write"
)

type Service interface {
	Write(ctx context.Context, req write.Request) (write.Response, error)
	Update(ctx context.Context, req update.Request) (update.Response, error)
	Read(ctx context.Context, req read.Request) (read.Response, error)
	Delete(ctx context.Context, request delete.Request) (delete.Response, error)
}
