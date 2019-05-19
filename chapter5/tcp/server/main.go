package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const addr = "localhost:8888"

func echoBackCapitalized(conn net.Conn) {
	// set up a reader on conn (an io.Reader)
	reader := bufio.NewReader(conn)

	// grab the first line of data encountered
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading data: %s\n", err.Error())
		return
	}
	// print then send back the data
	fmt.Printf("Received: %s", data)
	conn.Write([]byte(strings.ToUpper(data)))
	// close up the finished connection
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Printf("listening on: %s\n", addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("encountered an error accepting connection: %s\n", err.Error())
			// if there's an error try again
			continue
		}
		// handle this asynchronously
		// potentially a good use-case
		// for a worker pool
		go echoBackCapitalized(conn)
	}
}
