package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	router := fasthttprouter.New()
	router.GET("/item", GetItems)
	router.POST("/item/:item", AddItems)

	fmt.Println("server starting on localhost:8080")
	log.Fatal(fasthttp.ListenAndServe("localhost:8080", router.Handler))
}
