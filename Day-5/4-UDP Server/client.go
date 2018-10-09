package main

import (
	"net"
	"os"
	"fmt"
	"log"
)

const defaultBufferSize  = 512

func main() {

	if len(os.Args) != 2 {
		log.Fatal(fmt.Sprintf("Usage: %s host:port ", os.Args[0]))
	}

	endpoint := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", endpoint)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte("lalalala"))
	if err != nil {
		log.Fatal(err)
	}

	var buf [defaultBufferSize]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf[0:n]))

	os.Exit(0)
}
