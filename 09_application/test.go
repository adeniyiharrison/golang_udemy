package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	marshalExample()
}

func marshalExample() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
		Priority  bool
	}

	var personList []Person

	me := Person{
		FirstName: "Adeniyi",
		LastName:  "Harrison",
		Age:       30,
		Priority:  false,
	}

	you := Person{
		FirstName: "Chloe",
		LastName:  "Harrison",
		Age:       29,
		Priority:  true,
	}

	personList = append(personList, me, you)

	jsonBlob, err := json.Marshal(personList)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(string(jsonBlob))
	os.Stdout.Write(jsonBlob)
	fmt.Printf("found in location: %p\n", &jsonBlob)

}
