package main

import "fmt"

func main(){

	for i:=0 ; i < 101 ; i++ {
		if i%2 == 1 || i == 0 { // is it odd
			continue
		} else {
			fmt.Println(i)
		}
	}
}