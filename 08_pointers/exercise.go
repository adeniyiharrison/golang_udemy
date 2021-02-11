package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// a := p1(100)
	// fmt.Println(a)

	b := person{
		firstName: "Adeniyi",
		lastName:  "Harrison",
	}

	c := changeMe(&b)

	fmt.Println(b == c)

}

func p1(value int) *int {
	return &value
}

func changeMe(ogLocation *person) person {
	ogLocation.firstName = "Chloe"
	ogLocation.lastName = "Looker"

	return *ogLocation
}
