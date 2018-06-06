package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for i  := range numbers {
		callback(numbers[i])
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

// callback: passing a func as an argument
