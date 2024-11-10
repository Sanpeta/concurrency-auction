package auction

import (
	"context"

	"os"
	"strconv"
	"time"

	"github.com/Sanpeta/concurrency-auction/config/logger"
	"github.com/Sanpeta/concurrency-auction/internal/entity/auction_entity"
	"github.com/Sanpeta/concurrency-auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}
type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}
	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	go ar.startAuctionExpirationRoutine(auctionEntityMongo.Id)

	return nil
}

func (ar *AuctionRepository) startAuctionExpirationRoutine(auctionId string) {
	durationMinutes, _ := strconv.Atoi(os.Getenv("AUCTION_DURATION_MINUTES"))
	checkIntervalSeconds, _ := strconv.Atoi(os.Getenv("CHECK_INTERVAL_SECONDS"))

	expirationTime := time.Now().Add(time.Duration(durationMinutes) * time.Minute)

	ticker := time.NewTicker(time.Duration(checkIntervalSeconds) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if time.Now().After(expirationTime) {
			ar.closeAuction(auctionId)
			return
		}
	}
}

func (ar *AuctionRepository) closeAuction(auctionId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": auctionId}
	update := bson.M{"$set": bson.M{"status": auction_entity.Active}}

	_, err := ar.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error("Error closing auction: ", err)
	}
}
