package main

import (
	"fmt"
	"runtime"
	"time"
	"sync"
	"sync/atomic"
)

func main()  {
	fmt.Println(runtime.NumCPU()) // logical CPUs
	go func() {
		select {}
	}()

	fmt.Println(runtime.NumGoroutine()) // 2 = main + inf select

	ir := func() int {
		var i = 1
		fmt.Println(i)
		return i
	}

	go func(n int) {
		fmt.Println("Got,", n)
	}(ir())

	time.Sleep(1 * time.Second)

	var wg sync.WaitGroup
	var v int32 = 0

	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1) // wrong place, must be called before firing goroutine
			atomic.AddInt32(&v, 1)
			wg.Done() // or wg.Add(-1)
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&v)) // might print < 100
}
