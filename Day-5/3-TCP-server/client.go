package main

import (
	"os"
	"log"
	"fmt"
	"net"
	"io/ioutil"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal(fmt.Sprintf("Usage: %s host:port ", os.Args[0]))
	}

	endpoint := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))
}
