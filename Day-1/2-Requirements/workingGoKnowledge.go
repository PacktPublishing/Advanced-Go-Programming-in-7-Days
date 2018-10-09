package main

import (
	"fmt"
	"math"
	"sync"
)

// Variable declarations
var atom int
const monad = "monad"

// Know when init is called
func init()  {
	atom = 1
}

// function declarations
func calculateCeil(input float64) (result float64, err error)  {
	// Usage of math methods
	return math.Ceil(input), nil
}

// type definitions and aliases
type ByteSlice []byte
type Bytes = ByteSlice

// main
func main()  {
	// some basic printf flags
	fmt.Sprintf("%d\n", atom)
	fmt.Sprintf("%s", monad)

	// short declaration assignment
	stringSlice := [3]string{"1", "2", "3"}

	// basic slice operations
	fmt.Println(len(stringSlice))

	// underscore for ignoring results
	_, err := calculateCeil(1.2)

	// Checking for errors
	if err != nil {
		panic("Oops")
	}

	// infinite loop
	counter := 0
	for {
		if counter > 1000 {
			break
		} else {
			counter +=1
		}
	}

	// type switches
	var boolRef interface{}
	boolean := true
	boolRef = &boolean

	switch t := boolRef.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf("pointer to integer %d\n", *t)
	}

	// defer calls as lifo order
	defer fmt.Printf("%d ", 1)
	defer fmt.Printf("%d ", 2)

	l := new(sync.Mutex)
	l.Lock()
	defer l.Unlock()

	// Interfaces and export rules
	type Iterator interface {
		Next() interface{}
		hasNext() bool
	}

	type Guard struct {}

	// Channels
	ch := make(chan Guard, 1)

	// Goroutines
	go func(chan Guard) {
		fmt.Println("%T", ch)
		ch <- Guard{}
		close(ch)
	}(ch)

	<-ch
}
