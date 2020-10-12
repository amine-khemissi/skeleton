package count

import (
	"context"
	"github.com/amine-khemissi/skeleton/def/count"
)

func Count(ctx context.Context, req count.Request) (count.Response, error) {
	return count.Response{Count: len(req.Value)}, nil
}
