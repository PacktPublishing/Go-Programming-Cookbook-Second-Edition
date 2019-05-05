package main

import (
	mongodb "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter11/orchestrate"
)

func main() {
	if err := mongodb.Exec("mongodb://mongodb:27017"); err != nil {
		panic(err)
	}
}
