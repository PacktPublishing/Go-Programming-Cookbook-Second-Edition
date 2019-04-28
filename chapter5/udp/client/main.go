package main

import (
	"log"
	"net"
)

const addr = "localhost:8888"

func main() {
	log.Printf("client for server url: %s\n", addr)

	addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	msg := make([]byte, 512)
	n, err := conn.Write([]byte("connected"))
	if err != nil {
		panic(err)
	}
	for {
		n, err = conn.Read(msg)
		if err != nil {
			continue
		}
		log.Printf("%s\n", string(msg[:n]))
	}

}
