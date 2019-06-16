package firebase

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
)

// Client Interface for mocking
type Client interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}) error
	Close() error
}

// firestore.Client implements Close()
// we create Get and Set
type firebaseClient struct {
	*firestore.Client
	collection string
}

func (f *firebaseClient) Get(ctx context.Context, key string) (interface{}, error) {
	data, err := f.Collection(f.collection).Doc(key).Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get failed")
	}
	return data.Data(), nil
}

func (f *firebaseClient) Set(ctx context.Context, key string, value interface{}) error {
	set := make(map[string]interface{})
	set[key] = value
	_, err := f.Collection(f.collection).Doc(key).Set(ctx, set)
	return errors.Wrap(err, "set failed")
}
