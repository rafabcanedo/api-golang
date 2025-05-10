package mongodb

import (
	"context"
	"os"

	"github.com/rafabcanedo/api-golang/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// global var, but how we use constructor in project, we go to use this line
//var (
//	Connection mongo.Client
//)

var (
	MONGODB_URL = "MONGODB_URL"
	MONGO_USER_DB = "MONGO_USER_DB"
)

// ctx test, we don't have use this
func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongodb_ui := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGO_USER_DB)
	
	//client, _ := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_ui))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_ui))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Connect to database")

	return client.Database(mongodb_database), nil
}