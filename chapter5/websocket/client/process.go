package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func process(c *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter some text: ")
		// this will block ctrl-c, to exit press it then hit enter
		// or kill from another location
		data, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read stdin", err)
		}

		// trim off the space from reading the string
		data = strings.TrimSpace(data)

		// write the message as a byte across the websocket
		err = c.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			log.Println("failed to write message:", err)
			return
		}

		// this is an echo server, so we can always read after the write
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("failed to read:", err)
			return
		}
		log.Printf("received back from server: %#v\n", string(message))
	}
}
