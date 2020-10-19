package count

import (
	"context"

	"github.com/amine-khemissi/skeleton/backbone/logger"
	"github.com/amine-khemissi/skeleton/def/count"
)

func Count(ctx context.Context, req count.Request) (count.Response, error) {
	logger.Instance().Debug(ctx, "Count", req.Value)
	return count.Response{Count: len(req.Value)}, nil
}
