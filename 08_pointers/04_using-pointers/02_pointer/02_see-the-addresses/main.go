package main

import "fmt"

func zero(z *int) {
	x := 0
	var y int = 0

	fmt.Println(&z)
	fmt.Println(&x, &y)
	fmt.Println(z)
	*z = 0

}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x) // x is 0
}
