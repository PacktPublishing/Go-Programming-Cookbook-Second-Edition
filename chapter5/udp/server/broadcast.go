package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

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
