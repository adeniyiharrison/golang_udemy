package main

import "fmt"

const (
	c        = 42
	d string = "@niyiwi"
)

const (
	a1 = 2021 + iota
	a2 = 2021 + iota
	a3 = 2021 + iota
	a4 = 2021 + iota
)

// partOne response, question on markdown
func partOne() {
	a := 31
	fmt.Printf("%d\t%b\t%X\n", a, a, a)
}

// partTwo response
func partTwo() {
	a := 5
	b := 10

	same := a == b
	lessThanEqual := a <= b
	greaterThanEqual := a >= b
	lessThan := a < b
	greaterThan := a > b
	notEqual := a != b

	fmt.Printf("a: %d\nb: %d\n", a, b)
	fmt.Println("a == b: ", same)
	fmt.Println("a <= b: ", lessThanEqual)
	fmt.Println("a >= b: ", greaterThanEqual)
	fmt.Println("a < b: ", lessThan)
	fmt.Println("a > b: ", greaterThan)
	fmt.Println("a != b", notEqual)

}

// partThree response
func partThree() {
	fmt.Println("untyped constant: ", c)
	fmt.Println("typed constant: ", d)
}

// partFour response
func partFour() {
	a := 10
	fmt.Printf("%d\t%b\t%X\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%X\n", b, b, b)

}

//partFive response
func partFive() {
	a :=
		`
This is a raw string literal.

I could do multi line printing

with "qoutes" and stuff
	`
	fmt.Println(a)
}

// partSix response
func partSix() {
	fmt.Println(a1, a2, a3, a4)
}

func main() {
	partOne()
	partTwo()
	partThree()
	partFour()
	partFive()
	partSix()
}
