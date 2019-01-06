package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter6/storage"

func main() {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}
