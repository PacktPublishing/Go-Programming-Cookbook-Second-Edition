package main

import "github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter6/pools"

func main() {
	if err := pools.ExecWithTimeout(); err != nil {
		panic(err)
	}
}
