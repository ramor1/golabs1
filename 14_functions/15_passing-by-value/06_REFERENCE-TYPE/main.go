package main

import "fmt"

func main() {
	m := make([]string, 6, 10)
	fmt.Println(m) // [ ]
	for index, value := range m {
		fmt.Println(index, value, len(value))
	}
	fmt.Printf("%T\n", m)
	m = changeMe(m)
	fmt.Println(len(m))
	fmt.Println(m) // [Todd]
	for index, value := range m {
		fmt.Println(index, value, len(value))
	}
}

func changeMe(z []string) []string {
	z[0] = "Todd"
	z[1] = "Ren"
	z[5] = "Bobbbbbbbbaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	z = append(z, "Newentry6")
	z = append(z, "Newentry7")
	z = append(z, "Newentry8")
	z = append(z, "Newentry9")
	z = append(z, "Newentry10")
	z = append(z, "Newentry11")
	fmt.Println(len(z))
	fmt.Println(z) // [Todd]
	return (z)
}
