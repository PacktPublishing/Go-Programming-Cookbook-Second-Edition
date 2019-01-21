package main

import (
	"fmt"
	"net/http"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter7/twirp/rpc/greeter"
)

func main() {
	server := &Greeter{}
	twirpHandler := greeter.NewGreeterServiceServer(server, nil)

	fmt.Println("Listening on port :4444")
	http.ListenAndServe(":4444", twirpHandler)
}
