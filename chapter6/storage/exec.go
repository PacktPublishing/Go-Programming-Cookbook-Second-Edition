package storage

import (
	"context"
	"fmt"
)

// Exec initializes storage, then performs operations
// using the storage interface
func Exec() error {
	ctx := context.Background()
	m, err := NewMongoStorage(ctx, "localhost", "gocookbook", "items")
	if err != nil {
		return err
	}
	if err := PerformOperations(m); err != nil {
		return err
	}

	if err := m.Client.Database(m.DB).Collection(m.Collection).Drop(ctx); err != nil {
		return err
	}

	return nil
}

// PerformOperations creates a candle item
// then gets it
func PerformOperations(s Storage) error {
	ctx := context.Background()
	i := Item{Name: "candles", Price: 100}
	if err := s.Put(ctx, &i); err != nil {
		return err
	}

	candles, err := s.GetByName(ctx, "candles")
	if err != nil {
		return err
	}
	fmt.Printf("Result: %#v\n", candles)
	return nil
}
