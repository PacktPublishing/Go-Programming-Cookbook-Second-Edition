package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/gorilla/websocket"
)

// catchSig cleans up our websocket conenction if we kill the program
// with a ctrl-c
func catchSig(ch chan os.Signal, c *websocket.Conn) {
	// block on waiting for a signal
	<-ch
	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close:", err)
	}
	return
}

func main() {
	// connect the os signal to our channel
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// use the ws:// Scheme to connect to the websocket
	u := "ws://localhost:8000/"
	log.Printf("connecting to %s", u)

	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// dispatch our signal catcher
	go catchSig(interrupt, c)

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
