package main

import "fmt"

func main() {
	x := 13 % 2
	fmt.Println(x)
	fmt.Println(x == 1)
	if x ==  1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}
