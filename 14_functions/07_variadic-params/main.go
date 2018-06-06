package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	for key := range sf {
		fmt.Println(key)
		fmt.Println(sf[key])
	}

	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
