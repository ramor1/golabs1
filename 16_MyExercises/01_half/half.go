package main

import "fmt"

func checkEven(number int) (halfNumber int, evenFlag bool) {
	halfNumber = number / 2
	if number%2 == 0 && number != 0 {
		evenFlag = true
	} else {
		evenFlag = false
	}
	return
}

func main() {
	for i := 0; i <= 10; i++ {
		half, eflag := checkEven(i)
		fmt.Printf("i: %v \t half: %v \t Even Status: %v\n", i, half, eflag)
	}
}
