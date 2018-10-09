package main

import (
	"fmt"
	"sync"
	"time"
)

type T int

func IsClosed(c chan T) bool {
	select {
	case <-c:
		return true
	default:
		return false
	}
}

type ClosableChan struct {
	ch chan T
	once sync.Once
	isClosed bool
}

func (cc * ClosableChan) Close() {
	cc.once.Do(func() {
		close(cc.ch)
		cc.isClosed = true
	})
}

func (cc ClosableChan) IsClosed() bool {
	return cc.isClosed
}

func ping1(ch chan string) {
	ch <- "ping-1 successful"
}

func ping2(ch chan string) {
	ch <- "ping-2 successful"
}

func main()  {
	var c = make(chan T, 1)
	fmt.Println(IsClosed(c))

	// build-in methods
	fmt.Println(cap(c))
	fmt.Println(len(c))
	c <- 10
	fmt.Println(cap(c))
	fmt.Println(len(c))

	v, ok := <- c
	fmt.Println(v, ok)

	close(c)
	fmt.Println(IsClosed(c))

	v, ok = <- c
	fmt.Println(v, ok)

	// panics
	//c <- 10

	var cc = &ClosableChan{
		ch:make(chan T),
	}

	fmt.Println(cc.isClosed)
	cc.Close()
	cc.Close()

	// random select
	out1 := make(chan string)
	out2 := make(chan string)
	go ping1(out1)
	go ping2(out2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-out1:
		fmt.Println(s1)
	case s2 := <-out2:
		fmt.Println(s2)
	}
}