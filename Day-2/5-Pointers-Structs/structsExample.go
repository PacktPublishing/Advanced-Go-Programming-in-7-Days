package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
	age int
}


func main()  {
	alex := User{}
	fmt.Println(alex.age) //0
	alexP := &alex

	fmt.Println(alexP)

	var Worker = struct {
		User
		salary int
	}{
		User: alex,
		salary: 100000,
	}

	var AnotherWorker = struct {
		User
		salary int
	}{
		User: struct {
			name string
			age int
		}{
			"",
			0,
		},
		salary: 100000,
	}

	fmt.Println(Worker.salary)
	fmt.Println(Worker.name)
	fmt.Println(Worker.age)

	fmt.Println(AnotherWorker == Worker) // true

	a := struct {
		name string
		age int
	}{"", 0,}
	fmt.Println(reflect.DeepEqual(a, alex)) // false
}