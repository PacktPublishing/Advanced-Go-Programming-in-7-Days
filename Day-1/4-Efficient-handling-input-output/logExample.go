package main

import (
	"log"
	"fmt"
	"io/ioutil"
)

func main() {
	log.Println("Log entry")

	log.SetFlags(0)

	for i := 0; i< 100; i+=1 {
		go log.Println(i)
	}

	for i := 0; i< 100; i+=1 {
		go fmt.Println(i)
	}

	log.SetOutput(ioutil.Discard)

	log.Println("ENtry 2")

	defer log.Println("Will not be logged")
	log.Fatal("Exit")
}
