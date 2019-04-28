package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

const addr = "localhost:8888"

type connections struct {
	addrs map[string]*net.UDPAddr
	// lock for modifying the map
	mu sync.Mutex
}

func broadcast(conn *net.UDPConn, conns *connections) {
	count := 0
	for {
		count++
		conns.mu.Lock()
		// loop over known addresses
		for _, retAddr := range conns.addrs {

			// send a message to them all
			msg := fmt.Sprintf("Sent %d", count)
			if _, err := conn.WriteToUDP([]byte(msg), retAddr); err != nil {
				fmt.Printf("error encountered: %s", err.Error())
				continue
			}

		}
		conns.mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func main() {
	conns := &connections{
		addrs: make(map[string]*net.UDPAddr),
	}

	log.Printf("serving on %s\n", addr)

	// construct a udp addr
	addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}

	// listen on our specified addr
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	// cleanup
	defer conn.Close()

	// async send messages to all known clients
	go broadcast(conn, conns)

	msg := make([]byte, 1024)
	for {
		// receive a message to gather the ip address
		// and port to send back to
		_, retAddr, err := conn.ReadFromUDP(msg)
		if err != nil {
			continue
		}

		//store it in a map
		conns.mu.Lock()
		conns.addrs[retAddr.String()] = retAddr
		conns.mu.Unlock()
		log.Printf("%s connected\n", retAddr)
	}
}
