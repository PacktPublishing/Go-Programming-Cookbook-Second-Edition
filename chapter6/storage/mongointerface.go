package storage

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
)

// GetByName queries mongodb for an item with
// the correct name
func (m *MongoStorage) GetByName(ctx context.Context, name string) (*Item, error) {
	c := m.Client.Database(m.DB).Collection(m.Collection)
	var i Item
	if err := c.FindOne(ctx, bson.M{"name": name}).Decode(&i); err != nil {
		return nil, err
	}

	return &i, nil
}

// Put adds an item to our mongo instance
func (m *MongoStorage) Put(ctx context.Context, i *Item) error {
	c := m.Client.Database(m.DB).Collection(m.Collection)
	_, err := c.InsertOne(ctx, i)
	return err
}
