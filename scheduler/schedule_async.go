package scheduler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ch374n/vehicles-app/internal/models"
	"github.com/ch374n/vehicles-app/logger"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Scheduler struct {
	collection  *mongo.Collection
	redisClient *redis.Client
}

func NewScheduler(collection *mongo.Collection, redisClient *redis.Client) *Scheduler {
	return &Scheduler{
		collection:  collection,
		redisClient: redisClient,
	}
}

func (s *Scheduler) SyncToMongodb() {
	var cursor uint64 = 0
	log := logger.Get()

	for {
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		time.Sleep(next.Sub(now))

		keys, cursor, err := s.redisClient.Scan(cursor, "mfr:*", 50).Result()

		if err != nil {
			log.Fatal().Msg("error while scanning redis")
			break
		}

		for _, key := range keys {

			val, err := s.redisClient.Get(key).Result()

			if err != nil {
				log.Fatal().Msg("error while fetching record")
				break
			}

			var mfr models.Manufacturer

			err = json.Unmarshal([]byte(val), &mfr)

			if err != nil {
				log.Printf("Failed to unmarshal manufacturer: %v", err)
				continue
			}

			_, err = s.collection.UpdateOne(
				context.Background(),
				bson.M{"mfr": mfr.MfrID},
				bson.M{"$set": mfr},
				options.Update().SetUpsert(true),
			)
			if err != nil {
				log.Printf("Failed to sync manufacturer to MongoDB: %v", err)
			}

		}

		if cursor == 0 {
			break
		}
	}
}
