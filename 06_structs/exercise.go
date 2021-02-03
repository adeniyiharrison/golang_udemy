package main

import "fmt"

type person struct {
	first, last string
	favFlavors  []string
}

func p1() {

	me := person{
		first: "Adeniyi",
		last:  "Harrison",
		favFlavors: []string{
			"Cherry Garcia",
			"Milk and Cookies"},
	}

	her := person{
		first: "Chloe",
		last:  "Looker",
		favFlavors: []string{
			"Phish Food",
			"Matcha"},
	}

	for i, v := range me.favFlavors {
		fmt.Printf("My #%d favorite: %v\n", i+1, v)
		fmt.Printf("Her #%d favorite: %v\n", i+1, her.favFlavors[i])
	}

}

func p2() {
	me := person{
		first: "Adeniyi",
		last:  "Harrison",
		favFlavors: []string{
			"Cherry Garcia",
			"Milk and Cookies"},
	}

	her := person{
		first: "Chloe",
		last:  "Looker",
		favFlavors: []string{
			"Phish Food",
			"Matcha"},
	}

	employeeList := map[string]person{
		me.last:  me,
		her.last: her,
	}

	for k, v := range employeeList {
		fmt.Println(k)
		fmt.Println(v)
	}

	fmt.Println("Me: ")
	for i, v := range employeeList[me.last].favFlavors {
		fmt.Printf("%d: %v\n", i, v)
	}

	fmt.Println("Her: ")
	for i, v := range employeeList[her.last].favFlavors {
		fmt.Printf("%d: %v\n", i, v)
	}

}

func p3() {
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	fordBronco := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Royal Blue"},
		fourWheel: true,
	}

	fmt.Println(fordBronco)
	fmt.Println(fordBronco.color)

	mazdaMiata := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "Off White"},
		luxury: false,
	}

	fmt.Println(mazdaMiata)
	fmt.Println(mazdaMiata.luxury)

}

func p4() {
	employeeStruct := struct {
		personalInfo map[string]string
		title        string
		fullTime     bool
		managers     []string
	}{
		personalInfo: map[string]string{
			"First": "Adeniyi",
			"Last":  "Harrison",
			"City":  "Oakland, CA",
		},
		title:    "Data Engineer",
		fullTime: true,
		managers: []string{
			"Jonathan Stuckey",
			"Eric Albanese"},
	}

	fmt.Println(
		employeeStruct)
}

func main() {
	p4()
}
