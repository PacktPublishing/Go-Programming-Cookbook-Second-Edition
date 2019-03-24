package main

import (
	"fmt"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter10/pool"
)

func main() {
	cancel, in, out := pool.Dispatch(10)
	defer cancel()

	for i := 0; i < 10; i++ {
		in <- pool.WorkRequest{Op: pool.Hash, Text: []byte(fmt.Sprintf("messages %d", i))}
	}

	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		in <- pool.WorkRequest{Op: pool.Compare, Text: res.Wr.Text, Compare: res.Result}
	}

	for i := 0; i < 10; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}
		fmt.Printf("string: \"%s\"; matched: %v\n", string(res.Wr.Text), res.Matched)
	}
}
