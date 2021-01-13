package main

import (
	"fmt"
)

type testType int

var x int
var y string
var z bool
var a = 42
var b = "James Bond"
var c = true
var d testType
var f testType

// PartOne of Hands on Exercise see README.md for more details
func PartOne() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("%v\t%v\t%v\n", x, y, z)
	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)
	fmt.Println("Z: ", z)
}

// PartTwo of Hands on Exercise see README
func PartTwo() {
	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)
	fmt.Println("Z: ", z)

	// These values are called The zero value (default values)
}

// PartThree of Hands on Exercise see README
func PartThree() {
	s := fmt.Sprint(a, b, c)
	fmt.Println("S: ", s)
}

// PartFour of Hands on Exercise see README
func PartFour() {
	fmt.Println("D: ", d)
	fmt.Printf("%T\n", d)
	d = 42
	fmt.Println("D: ", d)

}

// PartFive of Hands on Exercise see README
func PartFive() {
	fmt.Println("F: ", f)
	fmt.Printf("%T\n", f)
	f = 99
	fmt.Println("F: ", f)
	f = testType(69)
	fmt.Println("F: ", f)
}

func main() {
	fmt.Println("Part One:")
	PartOne()

	fmt.Println("Part Two:")
	PartTwo()

	fmt.Println("Part Three:")
	PartThree()

	fmt.Println("Part Four:")
	PartFour()

	fmt.Println("Part Five:")
	PartFive()

}
