package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter8/middleware"
)

func main() {
	// We apply from bottom up
	h := middleware.ApplyMiddleware(
		middleware.Handler,
		middleware.Logger(log.New(os.Stdout, "", 0)),
		middleware.SetID(100),
	)
	http.HandleFunc("/", h)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
