package main

import "fmt"

func main() {
	ce := func(number int) (float32, bool) {
		halfNumber := float32(0)
		evenFlag := false
		halfNumber = float32(number) / 2	
		if number%2 == 0 && number != 0 {
			evenFlag = true
		} else {
			evenFlag = false
		}
		return halfNumber, evenFlag
	}

	for i := 0; i <= 10; i++ {
		half, eflag := ce(i)
		fmt.Printf("i: %v \t half: %v \t Even Status: %v\n", i, half, eflag)
	}

}
