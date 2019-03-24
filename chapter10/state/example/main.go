package main

import (
	"context"
	"fmt"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter10/state"
)

func main() {
	in := make(chan *state.WorkRequest, 10)
	out := make(chan *state.WorkResponse, 10)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go state.Processor(ctx, in, out)

	req := state.WorkRequest{state.Add, 3, 4}
	in <- &req

	req2 := state.WorkRequest{state.Subtract, 5, 2}
	in <- &req2

	req3 := state.WorkRequest{state.Multiply, 9, 9}
	in <- &req3

	req4 := state.WorkRequest{state.Divide, 8, 2}
	in <- &req4

	req5 := state.WorkRequest{state.Divide, 8, 0}
	in <- &req5

	for i := 0; i < 5; i++ {
		resp := <-out
		fmt.Printf("Request: %v; Result: %v, Error: %v\n", resp.Wr, resp.Result, resp.Err)
	}
}
