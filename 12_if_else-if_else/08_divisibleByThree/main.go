package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
	// Reverse a
	a  := []byte("Hello World!")
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		fmt.Println(i,j)
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(string(a))
}
