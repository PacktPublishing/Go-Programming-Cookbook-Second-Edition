package main

import (
	"fmt"
	"log"

	sarama "github.com/Shopify/sarama"
)

func sendMessage(producer sarama.SyncProducer, value string) {
	msg := &sarama.ProducerMessage{Topic: "example", Value: sarama.StringEncoder(value)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
		return
	}
	log.Printf("> message \"%s\" sent to partition %d at offset %d\n", msg.Value, partition, offset)
}

func main() {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	for i := 0; i < 10; i++ {
		sendMessage(producer, fmt.Sprintf("Message %d", i))
	}
}
