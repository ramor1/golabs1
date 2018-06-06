package main

import "fmt"

func findGreatest(fpargs ...float64) float64 {
	var greatest float64 = fpargs[0]
	for i := range fpargs[1:] {
		if fpargs[i] > greatest {
			greatest = fpargs[i]
		}
	}
	return greatest
}

func main() {
	great := findGreatest(10, 20, 30, 15, 100.0/3, 2)
	fmt.Println("The largest number is", great)
}
