package mongo

import (
	"context"
	"net/http"

	"github.com/amine-khemissi/skeleton/backbone/config"
	"github.com/amine-khemissi/skeleton/backbone/db"
	"github.com/amine-khemissi/skeleton/backbone/errorsklt"
	"github.com/amine-khemissi/skeleton/backbone/logger"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoConf struct {
	URI          string `json:"URI"`
	DatabaseName string `json:"name"`
}

func New(ctx context.Context) db.DB {
	tmp, isMSI := config.Instance().Get("database", "mongo").(map[string]interface{})
	if !isMSI {
		logger.Instance().Fatal(ctx, "failed to parse mongo conf, reason: expected json object")
	}
	var conf mongoConf
	if err := mapstructure.Decode(tmp, &conf); err != nil {
		logger.Instance().Fatal(ctx, "failed to parse mongo conf, reason:", err.Error())
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(conf.URI)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Instance().Fatal(ctx, "failed to ping mongo db", conf.URI)
	}

	// Check the connection
	if err = client.Ping(ctx, nil); err != nil {
		logger.Instance().Fatal(ctx, err)
	}

	logger.Instance().Info(ctx, "Connected to mongo:", conf.URI)
	return &wrapper{
		db: client.Database(conf.DatabaseName),
	}
}

type wrapper struct {
	db *mongo.Database
}

func (w *wrapper) SelectOne(ctx context.Context, collection string, filter, projection map[string]interface{}, out interface{}) (err error) {
	opt := options.FindOne()
	opt.Projection = projection
	result := w.db.Collection(collection).FindOne(ctx, filter, opt)
	if result.Err() != nil {
		return errorsklt.WithCode(errorsklt.Stack(result.Err(), "failed to select one result", filter, "in collection", collection), http.StatusInternalServerError)
	}
	if err := result.Decode(out); err != nil {
		return errorsklt.WithCode(errorsklt.Stack(err, "failed to decode one result with filter", filter, "in collection", collection), http.StatusInternalServerError)
	}
	return nil
}

func (w *wrapper) Select(ctx context.Context, collection string, filter, projection map[string]interface{}, results interface{}) (err error) {
	opt := options.Find()
	opt.Projection = projection
	cursor, err := w.db.Collection(collection).Find(ctx, filter, opt)
	if err != nil {
		return errorsklt.WithCode(errorsklt.Stack(err, "failed to select", filter, "in collection", collection), http.StatusInternalServerError)
	}

	if err := cursor.All(ctx, results); err != nil {
		return errorsklt.WithCode(errorsklt.Stack(err, "failed to parse cursor for filter", filter, "in collection", collection), http.StatusInternalServerError)
	}

	return nil
}

func (w *wrapper) Delete(ctx context.Context, collection string, filter map[string]interface{}) (int64, error) {

	deleteResult, err := w.db.Collection(collection).DeleteMany(ctx, filter)
	if err != nil {
		return 0, errorsklt.WithCode(errorsklt.Stack(err, "failed to delete with filter", filter, "in collection", collection), http.StatusInternalServerError)
	}

	return deleteResult.DeletedCount, nil
}

func (w *wrapper) DeleteOne(ctx context.Context, collection string, filter map[string]interface{}) (int64, error) {
	deleteResult, err := w.db.Collection(collection).DeleteOne(ctx, filter)
	if err != nil {
		return 0, errorsklt.WithCode(errorsklt.Stack(err, "failed to delete one with filter", filter, "in collection", collection), http.StatusInternalServerError)
	}
	return deleteResult.DeletedCount, nil
}
func (w *wrapper) InsertOne(ctx context.Context, collection string, item map[string]interface{}) error {
	//todo check how can I use the insertResult
	_, err := w.db.Collection(collection).InsertOne(ctx, item)
	if err != nil {
		return errorsklt.WithCode(errorsklt.Stack(err, "failed to insert one element in collection", collection), http.StatusInternalServerError)
	}
	return nil
}
