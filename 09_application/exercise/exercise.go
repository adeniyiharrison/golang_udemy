package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	employeeJSON := p1()
	p2(employeeJSON)
	p3(employeeJSON)
	p4()
	p5(employeeJSON)
}

func p1() *[]byte {

	fmt.Println("Problem #1")

	type personalInfo struct {
		FirstName, LastName string
		Age                 int
	}

	type employeeInfo struct {
		personalInfo
		Title  string
		Active bool
	}

	employee1 := employeeInfo{
		personalInfo: personalInfo{
			FirstName: "Adeniyi",
			LastName:  "Harrison",
			Age:       30},
		Title:  "Data Engineer",
		Active: true}

	employee2 := employeeInfo{
		personalInfo: personalInfo{
			FirstName: "Jonathan",
			LastName:  "Stuckey",
			Age:       35},
		Title:  "Manager",
		Active: true}

	employeeList := []employeeInfo{employee1, employee2}
	employeeJSON, err := json.Marshal(employeeList)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(employeeJSON))
	}

	return &employeeJSON
}

func p2(j *[]byte) {

	fmt.Println("Problem #2")

	type employeeJSON struct {
		FirstName string
		LastName  string
		Age       int
		Title     string
		Active    bool
	}

	var employeeList []employeeJSON

	err := json.Unmarshal(*j, &employeeList)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(employeeList)
	}

}

func p3(j *[]byte) {

	fmt.Println("Problem #3")

	type employeeJSON struct {
		FirstName string
		LastName  string
		Age       int
		Title     string
		Active    bool
	}

	var employeeList []employeeJSON

	// strings.NewReader(string) turns strings into Reader Type
	// Which json.NewDecoder needs a Reader Type
	decoderType := json.NewDecoder(strings.NewReader(string(*j)))

	err := decoderType.Decode(&employeeList)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(employeeList)
	}

	// WITHOUT DECODING!
	json.NewEncoder(os.Stdout).Encode(*j)
}

func p4() {

	fmt.Println("Problem #4")

	xi := []int{9, 7, 5, 2, 7}
	xs := []string{"Niyi", "Adeniyi", "Harrison", "Chloe", "Ola"}

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)

}

type employeeInfo struct {
	FirstName string
	LastName  string
	Age       int
	Active    bool
}

type employeeList []employeeInfo

func (el employeeList) Len() int {
	return len(el)
}
func (el employeeList) Swap(i, j int) {
	el[i], el[j] = el[j], el[i]
}
func (el employeeList) Less(i, j int) bool {
	return el[i].Age < el[j].Age
}

func p5(j *[]byte) {

	fmt.Println("Problem #5")

	var employeeStructArray employeeList

	newEmployee := employeeInfo{
		FirstName: "Ryan",
		LastName:  "Quigley",
		Age:       31,
		Active:    false,
	}

	jsonDecoder := json.NewDecoder(strings.NewReader(string(*j)))
	jsonDecoder.Decode(&employeeStructArray)
	employeeStructArray = append(employeeStructArray, newEmployee)
	sort.Sort(employeeStructArray)

	for i, v := range employeeStructArray {
		fmt.Printf("Employee #%d\n", i)
		fmt.Printf(
			"-- First: %s, Last: %s, Age: %d\n",
			v.FirstName, v.LastName, v.Age,
		)
	}

}
