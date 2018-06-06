package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}

func main() {
	greet := makeGreeter()
	greet1 := makeGreeter()
	fmt.Println(greet(), greet1())
}
