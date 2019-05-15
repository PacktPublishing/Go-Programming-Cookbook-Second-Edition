package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const addr = "localhost:8888"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// grab a string input from the clie
		fmt.Printf("Enter some text: ")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("encountered an error reading input: %s\n", err.Error())
			continue
		}
		// connect to the addr
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("encountered an error connecting: %s\n", err.Error())
		}

		// write the data to the connection
		fmt.Fprintf(conn, data)

		// read back the response
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("encountered an error reading response: %s\n", err.Error())
		}
		fmt.Printf("Received back: %s", status)
		// close up the finished connection
		conn.Close()
	}
}
