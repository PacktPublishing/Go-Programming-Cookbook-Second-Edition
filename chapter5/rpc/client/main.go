package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter5/rpc/tweak"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("error dialing:", err)
	}

	args := tweak.Args{
		String:  "this string should be uppercase and reversed",
		ToUpper: true,
		Reverse: true,
	}
	var result string
	err = client.Call("StringTweaker.Tweak", args, &result)
	if err != nil {
		log.Fatal("client call with error:", err)
	}
	fmt.Printf("the result is: %s", result)
}
