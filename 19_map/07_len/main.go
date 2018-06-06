package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(len(myGreeting))
 	for key, value := range(myGreeting){
		fmt.Println("Key", key, "\tValue:", value, "\tEntry", myGreeting[key])
	}
}
