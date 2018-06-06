package main

import "fmt"

func main() {
	var x [65600]string
	var i uint32

	fmt.Println(len(x))
	fmt.Println(x[0])
	for i = 0; i < 65600; i+= 64 {
		x[i] = string(rune(i))
	}
	var v string
	for i = 0; i < 65600; i+= 64 {
		v = x[i]
		fmt.Printf("%d %c -%T %T - % X ",i, v, i, v,  []byte(v))
		fmt.Println()
	}
}
