package write

import (
	"context"
	"crypto/md5"
	"fmt"
	"strconv"

	"github.com/amine-khemissi/skeleton/backbone/db"
	"github.com/amine-khemissi/skeleton/backbone/errorsklt"
	"github.com/amine-khemissi/skeleton/backbone/logger"
	"github.com/amine-khemissi/skeleton/def/write"
)

func Write(ctx context.Context, instance db.DB, req write.Request) (write.Response, error) {
	logger.Instance().Debug(ctx, "Write", req.Name, req.Age)

	hash := fmt.Sprintf("%x", md5.Sum([]byte(req.Name+strconv.Itoa(req.Age))))
	item := map[string]interface{}{
		"name": req.Name,
		"age":  req.Age,
	}
	if err := instance.InsertOne(ctx, "people", item); err != nil {
		return write.Response{}, errorsklt.Stack(err, "failed to insert new person")
	}
	return write.Response{
		ClientID: hash,
	}, nil
}
