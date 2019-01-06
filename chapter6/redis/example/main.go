package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter6/redis"

func main() {
	if err := redis.Exec(); err != nil {
		panic(err)
	}

	if err := redis.Sort(); err != nil {
		panic(err)
	}
}
