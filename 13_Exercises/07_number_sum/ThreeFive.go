package main

import (
		"fmt"
)

func main() {
	var sum int = 0
	for i := 1; i <1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	// p := message.NewPrinter(language.English)

	fmt.Printf("The sum of all of the multiples of 3 and 5 less than 1000 is %d\n", sum)
}
