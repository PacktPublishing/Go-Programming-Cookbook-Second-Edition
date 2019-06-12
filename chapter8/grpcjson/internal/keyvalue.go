package internal

import (
	"sync"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter8/grpcjson/keyvalue"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// KeyValue is a struct that holds a map
type KeyValue struct {
	mutex sync.RWMutex
	m     map[string]string
}

// NewKeyValue initializes the KeyValue struct and its map
func NewKeyValue() *KeyValue {
	return &KeyValue{
		m: make(map[string]string),
	}
}

// Set sets a value to a key, then returns the value
func (k *KeyValue) Set(ctx context.Context, r *keyvalue.SetKeyValueRequest) (*keyvalue.KeyValueResponse, error) {
	k.mutex.Lock()
	k.m[r.GetKey()] = r.GetValue()
	k.mutex.Unlock()
	return &keyvalue.KeyValueResponse{Value: r.GetValue()}, nil
}

// Get gets a value given a key, or say not found if
// it doesn't exist
func (k *KeyValue) Get(ctx context.Context, r *keyvalue.GetKeyValueRequest) (*keyvalue.KeyValueResponse, error) {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	val, ok := k.m[r.GetKey()]
	if !ok {
		return nil, grpc.Errorf(codes.NotFound, "key not set")
	}
	return &keyvalue.KeyValueResponse{Value: val}, nil
}
