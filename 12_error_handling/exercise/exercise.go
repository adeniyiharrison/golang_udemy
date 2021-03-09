package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	FirstName string
	LastName  string
	Sayings   []string
}

type customError struct {
	errorString string
}

func (c customError) Error() error {
	return errors.New(c.errorString)
}

func main() {
	person1 := person{
		FirstName: "Adeniyi",
		LastName:  "Harrison",
		Sayings: []string{
			"They have a way of working themselves out",
			"Shooters gon' shoot",
		},
	}
	p1(person1)
	p2(person1)
	p3()
}

func toJSON(p person) (string, error) {
	if p.FirstName == "Adeniyi" {
		// return "", fmt.Errorf("oh, %s is not allowed here", p.FirstName)
		log.Fatalf("oh, %s is not allowed here", p.FirstName)
	}
	bs, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}

	return string(bs), nil
}

func p1(p person) {
	bs, err := json.Marshal(p)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(bs))
}

func p2(p person) {
	js, err := toJSON(p)
	if err != nil {
		log.Println(err)
	}

	log.Println(js)

}

func foo(e error) {
	if e != nil {
		log.Println(e)
	}
}

func p3() {
	err := customError{
		errorString: "this is a test",
	}.Error()

	foo(err)

}
