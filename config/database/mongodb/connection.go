package mongodb

import (
	"context"
	"os"

	"github.com/Sanpeta/concurrency-auction/config/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// MongoDBHost is the key for the MongoDB host
	MongoDBHost = "MONGODB_HOST"
	// MongoDBPort is the key for the MongoDB port
	MongoDBPort = "MONGODB_PORT"
	// MongoDBDatabase is the key for the MongoDB database
	MongoDBDatabase = "MONGODB_DATABASE"
	// MongoDBPassword is the key for the MongoDB password
	MongoDBPassword = "MONGODB_PASSWORD"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongoHost := os.Getenv(MongoDBHost)
	mongoDatabase := os.Getenv(MongoDBDatabase)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoHost))
	if err != nil {
		logger.Error("error connecting to mongodb", err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Error("error pinging mongodb", err)
		return nil, err
	}

	return client.Database(mongoDatabase), nil
}
