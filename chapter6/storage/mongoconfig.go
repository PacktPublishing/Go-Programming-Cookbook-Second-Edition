package storage

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// MongoStorage implements our storage interface
type MongoStorage struct {
	*mongo.Client
	DB         string
	Collection string
}

// NewMongoStorage initializes a MongoStorage
func NewMongoStorage(ctx context.Context, connection, db, collection string) (*MongoStorage, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, "mongodb://localhost")
	if err != nil {
		return nil, err
	}

	ms := MongoStorage{
		Client:     client,
		DB:         db,
		Collection: collection,
	}
	return &ms, nil
}
