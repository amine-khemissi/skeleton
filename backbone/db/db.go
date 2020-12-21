package db

import "context"

type DB interface {
	DeleteOne(ctx context.Context, collection string, filter map[string]interface{}) (int64, error)
	Delete(ctx context.Context, collection string, filter map[string]interface{}) (deleteCount int64, err error)
	SelectOne(ctx context.Context, collection string, filter, projection map[string]interface{}, out interface{}) (err error)
	Select(ctx context.Context, collection string, filter, projection map[string]interface{}, results interface{}) (err error)
	InsertOne(ctx context.Context, collection string, item map[string]interface{}) (err error)
	UpdateOne(ctx context.Context, collection string, where map[string]interface{}, item map[string]interface{}) error
}
