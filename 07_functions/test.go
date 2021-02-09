package main

import "fmt"

func main() {
	// p1(
	// 	"test_string")

	// p2Test := p2(
	// 	"test_string_2")

	// fmt.Println(p2Test)

	// fullName, nameCheck := p3("Adeniyi", "Harrison")
	// if nameCheck {
	// 	fmt.Println("His name is", fullName)
	// } else {
	// 	fmt.Println("His name is NOT Adeniyi Harrison")
	// }

	// xi := []int{1, 2, 3, 4}
	// p4(xi...)

	// p5()

	// p6()

	// p7()

	// p8()

	p9()

}

func p1(s string) {
	fmt.Printf("Printing the arguement: %v\n", s)
}

func p2(s string) string {
	return fmt.Sprintf("Returned from p2: %v", s)
}

func p3(fn string, ln string) (string, bool) {

	fullName := fn + " " + ln
	nameCheck := fullName == "Adeniyi Harrison"

	return fullName, nameCheck
}

func p4(x ...int) {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Printf("%d: %d\n", i, v)
	}
	fmt.Println("Total Sum:", sum)
}

func p5() {
	defer p1("Run this last")
	p1("Run this one first!")
}

type person struct {
	firstName string
	lastName  string
	age       int
}

type employee struct {
	person
	title string
}

func (e employee) employeeInfo() bool {
	fmt.Printf(
		"First: %v\n Last: %v\n Age: %d\n Occupation: %v\n",
		e.firstName,
		e.lastName,
		e.age,
		e.title)

	var nameBool bool

	if e.firstName+" "+e.lastName == "Adeniyi Harrison" {
		nameBool = true
	}

	return nameBool
}

func p6() {
	employee1 := employee{
		person: person{
			"Adeniyi",
			"Harrison",
			30},
		title: "Data Engineer",
	}

	isMe := employee1.employeeInfo()

	fmt.Println(isMe)
}

type man struct {
	gender string
}

type woman struct {
	gender string
}

func (m man) speak() {
	fmt.Println("Man speaks")
}

func (w woman) speak() {
	fmt.Println("Woman speaks")
}

type human interface {
	speak()
}

func functionHumanParam(h human) {
	switch h.(type) {
	case man:
		fmt.Println("Gender: ", h.(man).gender)
	case woman:
		fmt.Println("Gender: ", h.(woman).gender)
	}
	fmt.Println("I am still human")
}

func p7() {
	m := man{
		gender: "Male",
	}

	f := woman{
		gender: "Female",
	}

	functionHumanParam(m)
	functionHumanParam(f)
}

// func p8() {
// 	var x int
// 	x = 1
// 	x++
// 	fmt.Println(x)
// 	{
// 		var y int
// 		y = 100
// 		y++
// 		fmt.Println(y)
// 	}

// 	y++
// }

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func p9() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
}
