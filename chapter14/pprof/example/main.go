package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter14/pprof/crypto"
)

func main() {

	http.HandleFunc("/guess", crypto.GuessHandler)
	fmt.Println("server started at localhost:8080")
	log.Panic(http.ListenAndServe("localhost:8080", nil))
}
