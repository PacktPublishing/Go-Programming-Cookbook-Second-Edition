package log

import (
	"bytes"
	"fmt"
	"log"
)

// Log uses the setup logger
func Log() {
	// we'll configure the logger to write
	// to a bytes.Buffer
	buf := bytes.Buffer{}

	// second argument is the prefix last argument is about options
	// you combine them with a logical or.
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ldate)

	logger.Println("test")

	logger.SetPrefix("new logger: ")

	logger.Printf("you can also add args(%v) and use Fataln to log and crash", true)

	fmt.Println(buf.String())
}
