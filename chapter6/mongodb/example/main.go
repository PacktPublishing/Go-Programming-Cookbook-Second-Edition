package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter6/mongodb"

func main() {
	if err := mongodb.Exec(); err != nil {
		panic(err)
	}
}
