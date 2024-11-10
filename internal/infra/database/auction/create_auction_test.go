// internal/infra/database/auction/create_auction_test.go
package auction

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Sanpeta/concurrency-auction/internal/entity/auction_entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAuctionExpiration(t *testing.T) {
	os.Setenv("AUCTION_DURATION_MINUTES", "1")
	os.Setenv("CHECK_INTERVAL_SECONDS", "1")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	assert.NoError(t, err)

	database := client.Database("testdb")
	repo := NewAuctionRepository(database)

	auction := &auction_entity.Auction{
		Id:          "test-auction",
		ProductName: "Test Product",
		Category:    "Test Category",
		Description: "Test Description",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   time.Now(),
	}

	err = repo.CreateAuction(context.Background(), auction)
	assert.NoError(t, err)

	time.Sleep(2 * time.Minute)

	var result AuctionEntityMongo
	err = repo.Collection.FindOne(context.Background(), bson.M{"_id": "test-auction"}).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, auction_entity.Completed, result.Status)
}
