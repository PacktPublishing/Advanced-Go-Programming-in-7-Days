package main

import "fmt"

func between(from, to int) []int {
	var result []int

	if from > to {
		result = []int{}
	} else {
		for i := from; i < to; i++ {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	for i := range between(0, 10) {
		switch i % 5 {
		case 1:
			fmt.Println("fizz")
		case 2:
			fmt.Println("bazz")
		case 3:
			fmt.Println("gizz")
			fallthrough
		default:
			fmt.Println("fizzbazz")
		}
	}
}
