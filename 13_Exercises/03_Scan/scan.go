package main

import (
	"fmt"
	"bufio"
	"os"
)

const DEBUG = true


func main(){
	reader := bufio.NewReader(os.Stdin)
	var buffer []byte
	var err error
	var prefix bool

	fmt.Print("What's your name ? ")
	buffer, prefix, err = reader.ReadLine( )
	// If prefix is true, the buffer filled up before a newline

	if DEBUG {fmt.Printf("prefix %v\nerr %v\nByte Count: %v\n", prefix,  err, len(buffer))}

	fmt.Println("Hello", string(buffer))}
