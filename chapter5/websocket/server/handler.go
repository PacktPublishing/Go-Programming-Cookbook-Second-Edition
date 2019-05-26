package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrader takes an http connection and converts it
// to a websocket one, we're using some recommended
// basic buffer sizes
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// upgrade the connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("failed to upgrade connection: ", err)
		return
	}
	for {
		// read and echo back messages in a loop
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("failed to read message: ", err)
			return
		}
		log.Printf("received from client: %#v", string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("failed to write message: ", err)
			return
		}
	}
}
