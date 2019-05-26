package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening on port :8000")
	// we mount our single handler on port localhost:8000 to handle all
	// requests
	log.Panic(http.ListenAndServe("localhost:8000", http.HandlerFunc(wsHandler)))
}
