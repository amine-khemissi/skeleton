package read

import (
	"context"

	"github.com/amine-khemissi/skeleton/backbone/db"
	"github.com/amine-khemissi/skeleton/backbone/errorsklt"
	"github.com/amine-khemissi/skeleton/backbone/logger"
	"github.com/amine-khemissi/skeleton/def/read"
)

func Read(ctx context.Context, instance db.DB, req read.Request) (read.Response, error) {
	logger.Instance().Debug(ctx, "read ", req.ClientID)

	item := map[string]interface{}{
		"ID": req.ClientID,
	}
	var resp read.Response
	if err := instance.SelectOne(ctx, "people", item, map[string]interface{}{
		"name": true,
		"age":  true,
	}, &resp); err != nil {
		return read.Response{}, errorsklt.Stack(err, "failed to select person", req.ClientID)
	}
	return resp, nil
}
