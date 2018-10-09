package main

import (
	"runtime/debug"
	"fmt"
)

func main()  {
	// main go-routine
	debug.PrintStack()
	defer debug.PrintStack()
	done := make(chan bool)

	// from goroutine
	go func(done chan bool) {
		debug.PrintStack()
		done<- true
		close(done)
	}(done)
	<-done

	// using Stack()
	stackTrace := debug.Stack()
	fmt.Printf("%v", string(stackTrace))
}
