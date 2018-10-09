package main

import (
	"unsafe"
	"fmt"
)

// internal representation
type MySlice struct {
	elems unsafe.Pointer
	len int
	cap int
}

const e = 2

func main()  {
	var mapping = make(map[int]int)
	mapping[e] = e
	//var myptr = &ma[1]

	value, ok := mapping[e]
	fmt.Println(value, ok)

	//var array = [...]byte{1,2,3}
	var array2 = [...]byte{2: 1, 3: 2,4: 3}
	fmt.Println(array2) // [0 0 1 2 3]
	// _ = array2[-1] no negative index

	// build-in methods
	fmt.Println(len(array2), cap(array2)) // 5 5
	fmt.Println(len(mapping)) // 1
	delete(mapping, e)
	fmt.Println(len(mapping)) // 0

	m := new(map[int]int) // makes no sense
	fmt.Println(m)
	//&m[1] = 1
}
