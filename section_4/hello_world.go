package main

import (
	"fmt"
)

var y = 42

func main() {
	NumberOfBytes, _ := fmt.Println("this is a test")
	fmt.Println(NumberOfBytes)
	// "_" is throwing away the error statement returned (if there was an error)
	// This is being thrown away because you cannot have a variable and not use it in go

	z := 43
	fmt.Println("global var, ", y)
	fmt.Printf("%T\n", z)
}
