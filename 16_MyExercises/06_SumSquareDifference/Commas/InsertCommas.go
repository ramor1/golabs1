package Commas

import "fmt"

func InsertCommas(number int) string {
	str1 := fmt.Sprint(number)
	digits := len(str1)
	commas := digits / 3
	if digits%3 == 0  {
		commas--
	}
	blength := digits + commas
	buffer := make([]byte, blength)
	bptr := blength - 1
	for i := 1; i <= digits; i++ {
		buffer[bptr] = str1[digits-i]
		bptr--
		if i%3 == 0 && commas != 0 && bptr >0 {
			buffer[bptr] = ','
			bptr--
		}
	}
	return string(buffer)
}

