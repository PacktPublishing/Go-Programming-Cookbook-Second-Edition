package main

import (
	"net/http"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter11/consensus"
)

func main() {
	consensus.Config(3)

	http.HandleFunc("/", consensus.Handler)
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
