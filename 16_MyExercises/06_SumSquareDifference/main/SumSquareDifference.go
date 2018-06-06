package main

import (
		"fmt"
		"github.com/ramor1/golabs1/16_MyExercises/06_SumSquareDifference/Commas"
		)

const MAX = 100

func sumOfSquare(limit int) int {
	var sum int
	for i := 1; i <= limit; i++ {
		sum += i * i
	}
	return sum
}
func squareOfSum(limit int) int {
	sum := 0
	for i := 1; i <= limit; i++ {
		sum += i
	}
	return sum * sum
}

func main() {
	sum1 := sumOfSquare(MAX)
	sum2 := squareOfSum(MAX)
	diff := sum2 - sum1
	fmt.Println("Sum of the first", MAX, "squares is :", Commas.InsertCommas(sum1))
	fmt.Println("Sum of the first", MAX, "numbers squared is :", Commas.InsertCommas(sum2))
	fmt.Println("Difference of the two is :", Commas.InsertCommas(diff))
	fmt.Println("----------------------------------------------")
}

/*
Ref: https://projecteuler.net/problem=6

Sum square difference
Problem 6
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
 */