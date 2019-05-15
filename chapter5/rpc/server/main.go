package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter5/rpc/tweak"
)

func main() {
	s := new(tweak.StringTweaker)
	if err := rpc.Register(s); err != nil {
		log.Fatal("failed to register:", err)
	}

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	fmt.Println("listening on :1234")
	log.Panic(http.Serve(l, nil))
}
