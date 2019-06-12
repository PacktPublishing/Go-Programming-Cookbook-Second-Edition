package mongodb

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
)

// State is our data model
type State struct {
	Name       string `bson:"name"`
	Population int    `bson:"pop"`
}

// Exec creates then queries an Example
func Exec(address string) error {
	ctx := context.Background()
	db, err := Setup(ctx, address)
	if err != nil {
		return err
	}

	coll := db.Database("gocookbook").Collection("example")

	vals := []interface{}{&State{"Washington", 7062000}, &State{"Oregon", 3970000}}

	// we can inserts many rows at once
	if _, err := coll.InsertMany(ctx, vals); err != nil {
		return err
	}

	var s State
	if err := coll.FindOne(ctx, bson.M{"name": "Washington"}).Decode(&s); err != nil {
		return err
	}

	if err := coll.Drop(ctx); err != nil {
		return err
	}

	fmt.Printf("State: %#v\n", s)
	return nil
}
