package mongodb

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// Setup initializes a redis client
func Setup(ctx context.Context) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, "mongodb://localhost")
	if err != nil {
		return nil, err
	}
	return client, nil
}
