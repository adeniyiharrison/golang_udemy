package main

import "fmt"

type person struct {
	first, last string
	age         int
}

type employee struct {
	person
	title  string
	salary int
}

func p1() {

	person1 := person{
		first: "Adeniyi",
		last:  "Harrison",
	}

	person2 := person{
		first: "Chloe",
		last:  "Looker",
	}

	fmt.Println(person1.first)
	fmt.Println(person2.last)
}

func p2() {

	person1 := employee{
		person: person{
			first: "Adeniyi",
			last:  "Harrison",
			age:   30},
		title:  "data engineer",
		salary: 5,
	}

	person2 := employee{
		person: person{
			first: "Chloe",
			last:  "Looker",
			age:   29},
		title:  "MFA student",
		salary: 0,
	}

	fmt.Println(person1.person)
	fmt.Println(person2.title)
}

func p3() {
	person := struct {
		first, last, title string
		age, salary        int
	}{
		first:  "Olawale",
		last:   "Harrison",
		age:    31,
		title:  "Surfer",
		salary: 100,
	}

	fmt.Println(person)
}

func main() {
	p3()
}
