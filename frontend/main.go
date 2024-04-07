package main

import (
	"fmt"
	"os"
)

func main () {
	bytes, err := os.ReadFile("../examples/expressions.nar")
	if err != nil {
		fmt.Println("Error Occured While Reading The Source Code")
	}
	source_string := string(bytes) 

	// print the source code 
	fmt.Println(source_string)
}