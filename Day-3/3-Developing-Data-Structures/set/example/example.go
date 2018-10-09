package main

import (
	"fmt"
	"github.com/PacktPublishing/Advanced-Go-Programming-in-7-Days/Day-3/3-Developing-Data-Structures/set"
)

func main()  {
	var s = set.New()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s.Add(5)

	fmt.Println(s.Size())// 5
	s.Add(1)
	s.Remove(6)
	fmt.Println(s.Size())// 5
	fmt.Println(s.IsElementOf(3)) // true
	fmt.Println(s.Values())
}