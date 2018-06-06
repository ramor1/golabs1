package main

import "fmt"

func foo(iarg ...int) int {
	counter := 0
	for _, i := range iarg {
		counter += i
	}
	return counter
}

func main() {
	fmt.Println(foo(1, 2))
	fmt.Println(foo(1, 2, 3))
	aSlice := []int{1, 2, 3, 4}
	fmt.Println(foo(aSlice...))
	fmt.Println(foo())
}
