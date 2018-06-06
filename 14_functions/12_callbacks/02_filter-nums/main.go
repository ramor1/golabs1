package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	var xs2 []int
	for _, n := range numbers {
		if callback(n) {
			xs2 = append(xs2, n)
		}
	}
	return xs2
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs) // [2 3 4]
}
