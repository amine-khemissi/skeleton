package read

import (
	"context"

	"github.com/amine-khemissi/skeleton/backbone/db"
	"github.com/amine-khemissi/skeleton/backbone/errorsklt"
	"github.com/amine-khemissi/skeleton/backbone/logger"
	"github.com/amine-khemissi/skeleton/def/delete"
)

func Delete(ctx context.Context, instance db.DB, req delete.Request) (delete.Response, error) {
	logger.Instance().Debug(ctx, "delete ", req.ClientID)

	item := map[string]interface{}{
		"ID": req.ClientID,
	}
	var resp delete.Response
	nbDeleted, err := instance.DeleteOne(ctx, "people", item)
	if err != nil {
		return delete.Response{}, errorsklt.Stack(err, "failed to delete person", req.ClientID)
	}
	logger.Instance().Debug(ctx, "deleted ", nbDeleted, "items")
	return resp, nil
}
