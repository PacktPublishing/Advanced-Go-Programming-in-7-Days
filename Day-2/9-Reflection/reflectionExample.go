package main

import (
	"reflect"
	"time"
	"fmt"
)

type UserType reflect.Type

type User struct {
	FirstName string
	LastName  string
	Birthday  time.Time
}

func (u User) String() string  {
	return fmt.Sprintf("User: %v, %v", u.FirstName, u.LastName)
}

func main()  {
	alex := User{}
	userType := reflect.TypeOf(alex)

	//fmt.Println(userType.Elem()) // panics
	fmt.Println(userType.NumField()) // 3
	fmt.Println(userType.Comparable()) // true
	fmt.Println(userType.Kind()) // struct
	fmt.Println(userType.NumMethod()) // 1, Value vs Ref receiver matters
	fmt.Println(userType.MethodByName("String")) // case matters

	// Create slices via reflection
	intSlice := reflect.MakeSlice(reflect.TypeOf([]int{}), 0, 0)
	fmt.Println(intSlice)
	intSlice = reflect.Append(intSlice, reflect.ValueOf(1))
	fmt.Println(intSlice) // [1]

	intArrayType := reflect.ArrayOf(5, reflect.TypeOf(0))
	intArray := reflect.New(intArrayType)
	fmt.Println(intArray) // &[0 0 0 0 0]

	var n = []int{1,2,3}
	var p = reflect.ValueOf(&n)
	fmt.Println(p.CanSet()) // false
	fmt.Println(p.CanAddr()) // false
	var nv = p.Elem()
	fmt.Println(nv.CanSet()) // true
	fmt.Println(nv.CanAddr()) // true
}
