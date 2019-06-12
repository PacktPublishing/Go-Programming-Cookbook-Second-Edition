package mongodb

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Setup initializes a mongo client
func Setup(ctx context.Context, address string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// cancel will be called when setup exits
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	return client, nil
}
