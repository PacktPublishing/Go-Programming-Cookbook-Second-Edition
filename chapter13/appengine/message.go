package main

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
)

// Message is the object we store
type Message struct {
	Timestamp time.Time
	Message   string
}

func (c *Controller) storeMessage(ctx context.Context, message string) error {
	m := &Message{
		Timestamp: time.Now(),
		Message:   message,
	}

	k := datastore.IncompleteKey("Message", nil)
	_, err := c.store.Put(ctx, k, m)
	return err
}

func (c *Controller) queryMessages(ctx context.Context, limit int) ([]*Message, error) {
	q := datastore.NewQuery("Message").
		Order("-Timestamp").
		Limit(limit)

	messages := make([]*Message, 0)
	_, err := c.store.GetAll(ctx, q, &messages)
	return messages, err
}
