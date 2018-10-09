package main

import (
	"net"
	"log"
	"time"
	"fmt"
)

const defaultHostPort  = ":9000"

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", defaultHostPort)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		// conn.(*net.TCPConn).SetKeepAlive(true) Set keepalive
		// timeoutDuration := 5 * time.Second
		// conn.(*net.TCPConn).SetDeadline(time.Now().Add(timeoutDuration))

		if err != nil {
			continue
		}

		datetime := time.Now().String()
		go func() {
			conn.Write([]byte(fmt.Sprintf("Time is: %q", datetime)))
			defer conn.Close()
		}()
	}
}
