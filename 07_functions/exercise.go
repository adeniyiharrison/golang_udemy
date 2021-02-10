package main

import (
	"fmt"
	"math"
)

func main() {
	// question1()
	// question2()
	// question3()
	// question4()
	// question5(7.45)
	// question6()
	// question8()
	question9(returnFunction())
}

func q1f1() int {
	return 69
}

func q1f2() (int, string) {
	return 1, "hello"
}

func question1() {
	foo := q1f1()
	bar1, bar2 := q1f2()

	fmt.Printf("Foo: %d, Bar1: %d, Bar2: %s\n", foo, bar1, bar2)

}

func q2f1(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

func q2f2(xi []int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}

func question2() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultOne := q2f1(xi...)
	resultTwo := q2f2(xi)

	fmt.Printf("Result One: %d\n Result Two: %d\n", resultOne, resultTwo)
}

func question3() {
	onTheFly := func() {
		fmt.Println("and we out!")
	}

	defer onTheFly()

	fmt.Println("this is where we start..")
}

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) speak() string {
	x := fmt.Sprintf(
		"Hi, I am %s %s and I am %d years old.\n",
		p.firstName,
		p.lastName,
		p.age,
	)

	return x
}

func question4() {
	me := person{
		firstName: "Adeniyi",
		lastName:  "Harrison",
		age:       30,
	}

	words := me.speak()

	fmt.Println(words)
}

type square struct {
	length float64
}

type circle struct {
	diameter float64
}

func (s square) returnArea() float64 {
	return math.Pow(s.length, 2)
}

func (c circle) returnArea() float64 {
	radius := (c.diameter / 2)
	return math.Pi * math.Pow(radius, 2)
}

type shapes interface {
	returnArea() float64
}

func info(s shapes) {
	switch s.(type) {
	case square:
		area := s.(square).returnArea()
		fmt.Println("I am a square with an area of", area)
	case circle:
		area := s.(circle).returnArea()
		fmt.Println("I am a circle with an area of", area)
	}
}

func question5(length float64) {
	s := square{
		length: length}

	c := circle{
		diameter: length}

	info(s)
	info(c)

}

func question6() {

	// contains 6 and 7

	func() {
		fmt.Println("This is an anonymous function")
	}()

	funcExpression := func() {
		fmt.Println("This is an function expression")
	}

	funcExpression()
}

func returnFunction() func() {
	return func() {
		fmt.Println("Originally from returnFunc")
	}
}

func question8() {
	f := returnFunction()

	f()
}

func question9(f func()) {
	f()
}
