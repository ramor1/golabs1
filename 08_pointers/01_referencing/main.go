package main

import (
	"fmt"
)

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(b)
	fmt.Println(*b)
	fmt.Printf("a is of type %T, \nb is of type %T\n",a, b)

	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int
}
