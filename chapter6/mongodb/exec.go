package mongodb

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// State is our data model
type State struct {
	Name       string `bson:"name"`
	Population int    `bson:"pop"`
}

// Exec creates then queries an Example
func Exec() error {
	db, err := Setup()
	if err != nil {
		return err
	}

	conn := db.DB("gocookbook").C("example")

	// we can inserts many rows at once
	if err := conn.Insert(&State{"Washington", 7062000}, &State{"Oregon", 3970000}); err != nil {
		return err
	}

	var s State
	if err := conn.Find(bson.M{"name": "Washington"}).One(&s); err != nil {
		return err
	}

	if err := conn.DropCollection(); err != nil {
		return err
	}

	fmt.Printf("State: %#v\n", s)
	return nil
}
