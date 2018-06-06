package main

import (
	"fmt"
)

func main() {
	var student []string
	var students [][]string
	fmt.Println(student)
	fmt.Println(students)
	student = append(student, "Ren","Bob")
	students = append(students, student)
	fmt.Println("Capacity:", cap(student), cap(students), "Length: ", len(student), len(students))
	fmt.Println(students)
	for i := range(students[0]){
		fmt.Println(students[0][i])
	}
	fmt.Println(student == nil)
}
