package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// marshalExample()
	// unmarshalExample()
	// sorting()
	// sorting2()
	cryptoTest()
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
		LastName:  "Looker",
		Age:       29,
		Priority:  true,
	}

	personList = append(personList, me, you)
	// personList := []persons{me, you}

	jsonBlob, err := json.Marshal(personList)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(string(jsonBlob))
	os.Stdout.Write(jsonBlob)
	fmt.Printf("\nfound in location: %p\n", &jsonBlob)

}

func unmarshalExample() {

	// this is from https://mholt.github.io/json-to-go/
	type person struct {
		FirstName string `json:"FirstName"`
		LastName  string `json:"LastName"`
		Age       int    `json:"Age"`
		Priority  bool   `json:"Priority"`
	}

	var people []person

	jsonString := `[{"FirstName":"Adeniyi","LastName":"Harrison","Age":30,"Priority":false},{"FirstName":"Chloe","LastName":"Looker","Age":29,"Priority":true}]`

	json.Unmarshal([]byte(jsonString), &people)

	fmt.Println(people)

}

func sorting() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	b := sort.StringsAreSorted(s)
	fmt.Println(b)
}

type person struct {
	firstName string
	lastName  string
	age       int
	priority  bool
}

type byAge []person

func (b byAge) Len() int {
	return len(b)
}
func (b byAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b byAge) Less(i, j int) bool {
	return b[i].age < b[j].age
}

func sorting2() {

	var personList byAge

	me := person{
		firstName: "Adeniyi",
		lastName:  "Harrison",
		age:       30,
		priority:  false,
	}

	you := person{
		firstName: "Chloe",
		lastName:  "Looker",
		age:       29,
		priority:  true,
	}

	another := person{
		firstName: "Amanda",
		lastName:  "Looker",
		age:       34,
		priority:  false,
	}

	personList = append(personList, me, you, another)
	fmt.Printf("PersonList is of type: %T\n", personList)
	fmt.Println(personList)

	sort.Sort(personList)

	fmt.Println("--------")
	fmt.Println(personList)

}

func cryptoTest() {
	pw := "password123"
	a, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}

	err = bcrypt.CompareHashAndPassword(a, []byte(pw))
	if err == nil {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}

}
