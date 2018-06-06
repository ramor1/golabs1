package main

import "fmt"

func main() {
	var x [58]string

	for i := 65; i <= 122; i++ { // 65 = 'A' 122 ='z'
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(x[42])
}
