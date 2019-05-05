package main

import (
	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter11/orchestrate"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("mongodb")
	if err != nil {
		panic(err)
	}
	if err := orchestrate.ConnectAndQuery(session); err != nil {
		panic(err)
	}
}
