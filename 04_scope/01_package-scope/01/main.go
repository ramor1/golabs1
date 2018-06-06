package main

import "fmt"

var x int = 42
var y = 18

func main() {
	fmt.Println(x)
	y := 17
	fmt.Println(y)
	foo()
}

func foo() {
	fmt.Println(x)
	fmt.Println(y)
}
