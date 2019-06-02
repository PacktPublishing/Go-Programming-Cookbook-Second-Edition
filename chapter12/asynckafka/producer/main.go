package main

import (
	"fmt"
	"net/http"

	sarama "github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer producer.AsyncClose()

	go ProcessResponse(producer)

	c := KafkaController{producer}
	http.HandleFunc("/", c.Handler)
	fmt.Println("Listening on port :3333")
	panic(http.ListenAndServe(":3333", nil))
}
