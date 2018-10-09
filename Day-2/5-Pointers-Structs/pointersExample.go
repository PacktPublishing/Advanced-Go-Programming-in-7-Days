package main

import "fmt"

type P int

func add2(a int) func() int  {
	return func() int {
		return a + 2
	}
}

func add2Ref(a *int) func() int  {
	return func() int {
		return *a + 2
	}
}

func add2Ref2(a *int) func() int  {
	var b = *a // save inside closure
	return func() int {
		return b + 2
	}
}

func main()  {
	n := new(P)
	*n = 1
	fmt.Println(&*n == n)
	//fmt.Println(*n++) // no direct pointer arithmetic (use unsafe.Pointer)

	a2 := add2(1)
	fmt.Println(a2()) // 3

	var a = 3
	a2r := add2Ref(&a)
	a = 5
	fmt.Println(a2r()) //7

	a = 3
	a2r2 := add2Ref2(&a)
	a = 5
	fmt.Println(a2r2()) //5

	s1 := [...]int{1,2,3,4,5}
	// s is pointer to array
	for _, v := range &s1 {
		fmt.Println(v)
	}

	//s2 := []int{1,2,3,4,5}
	// wont work on slices
	//for _, v := range &s2 {
	//	fmt.Println(v)
	//}
}