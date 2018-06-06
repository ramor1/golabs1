package main

import (
	"fmt"
)

const DEBUG = true

func main(){
	// var small, large string
	var count, ismall, ilarge int
	// var fsmall float64
	var err error

	for {
		fmt.Print("Please enter a small number and a large number :")
		count, err = fmt.Scan(&ismall, &ilarge)
		if DEBUG {fmt.Println(count, err, ismall, ilarge)}
		if err != nil {
			fmt.Println("\t\tError : ", err, "\n")
		}

		if err == nil {break};
	}


	fmt.Println( "The remainder of ", ilarge, "modulo", ismall, "is", ilarge % ismall)


}
