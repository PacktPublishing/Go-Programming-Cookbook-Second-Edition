package storage

import mgo "gopkg.in/mgo.v2"

// MongoStorage implements our storage interface
type MongoStorage struct {
	*mgo.Session
	DB         string
	Collection string
}

// NewMongoStorage initializes a MongoStorage
func NewMongoStorage(connection, db, collection string) (*MongoStorage, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	ms := MongoStorage{
		Session:    session,
		DB:         db,
		Collection: collection,
	}
	return &ms, nil
}
