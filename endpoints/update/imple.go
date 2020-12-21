package update

import (
	"context"

	"github.com/amine-khemissi/skeleton/backbone/db"
	"github.com/amine-khemissi/skeleton/backbone/errorsklt"
	"github.com/amine-khemissi/skeleton/backbone/logger"
	"github.com/amine-khemissi/skeleton/def/update"
)

func Update(ctx context.Context, instance db.DB, req update.Request) (update.Response, error) {
	logger.Instance().Debug(ctx, "Update", req.Name, req.Age)

	where := map[string]interface{}{
		"ID": req.ClientID,
	}
	item := map[string]interface{}{}
	if req.Age != 0 {
		item["age"] = req.Age
	}
	if req.Name != "" {
		item["name"] = req.Name
	}
	if err := instance.UpdateOne(ctx, "people", where, item); err != nil {
		return update.Response{}, errorsklt.Stack(err, "failed to update person", req.ClientID)
	}
	return update.Response{}, nil
}
