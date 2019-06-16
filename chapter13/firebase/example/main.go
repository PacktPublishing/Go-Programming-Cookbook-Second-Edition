package main

import (
	"context"
	"fmt"
	"log"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter13/firebase"
)

func main() {
	ctx := context.Background()
	c, err := firebase.Authenticate(ctx, "collection")
	if err != nil {
		log.Fatalf("error initializing client: %v", err)
	}
	defer c.Close()

	if err := c.Set(ctx, "key", []string{"val1", "val2"}); err != nil {
		log.Fatalf(err.Error())
	}

	res, err := c.Get(ctx, "key")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(res)

	if err := c.Set(ctx, "key2", []string{"val3", "val4"}); err != nil {
		log.Fatalf(err.Error())
	}

	res, err = c.Get(ctx, "key2")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(res)
}
