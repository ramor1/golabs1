package main

import "fmt"

func main() {
	for i := 50; i <= 512; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []rune(string(i)))
	}

	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)

}
