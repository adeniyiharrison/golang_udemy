# Section 18 Notes

## Application

### JSON documentation
__Marshal__
* Marshal returns the JSON encoding of v. `func Marshal(v interface{}) ([]byte, error)`
* `err` will `nil` if there are no errors
* Capitalize names of the keys in the `struct` for things to be exported
```
import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
```
* Output
```
{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
```

__Unmarshal__
* Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
* `func Unmarshal(data []byte, v interface{}) error`
    * Slice of bytes is very close to a string acorriding to Todd
    * v is a pointer to a location
```
func main() {
    var jsonBlob := []byte(`[
        {"FirstName": "Adeniyi", "Age": 30},
        {"FirstName": "Chloe", "Age": 29}
    ]`)

    type person struct {
        firstName string
        age int
    }

    var persons []person

    err := json.Unmarshal(jsonBlob, &persons)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Printf("%+v", persons)
}
```
* Output
```
[{FirstName: Adeniyi Age: 30} {FirstName: Chloe Age: 30}]
```
* `jsonBlob` creates a `[]byte` (since strings are just an array of bytes) of string representation of 2 json blobs
* Then turns them into `person` structs within the program

## Deeper Dive to Marshal
```
import "encoding/json"

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
```
* Package location is `encoding/json`
* When marshalling your keys within a struct type need to be title cased, this syntax indicates that the data will be available outside of the program
* `func Marshal(v interface{}) ([]byte, error)`
    * empty interface means it can take any type
    * Returns `[]byte` version of json object which can be converted to a human-legible form by using `string()` conversion or `os.Stdout.Write()`

## Deeper Dive into Unmarshal
* `func Unmarshal(data []byte, v interface{}) error`
    * Unmarshal parses the JSON encoded data and stores the result in the value pointed to by v
* https://mholt.github.io/json-to-go/
    * Paste in json and the tool will create what that data type (`struct`) should look like
    * copied in `[{"FirstName":"Adeniyi","LastName":"Harrison","Age":30,"Priority":false},{"FirstName":"Chloe","LastName":"Looker","Age":29,"Priority":true}]`
    ```
        type AutoGenerated []struct {
            FirstName string `json:"FirstName"`
            LastName  string `json:"LastName"`
            Age       int    `json:"Age"`
            Priority  bool   `json:"Priority"`
        }
    ```
    * __A field declaration may be followed by an optional string literal tag, which becomes an attribute for all the fields in the corresponding field declaration. An empty tag string is equivalent to an absent tag. The tags are made visible through a reflection interface and take part in type identity for structs but are otherwise ignored.__
    * Basically whats happening in the `AutoGenerated` struct type, you are mapping fields from json to Go. In this case the mapping would have been happened organically because the field names are the same but you can force it if the field names were different

```
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
```
* updated the `autoGenerated` type to be a `struct` and not a `[]struct` and called it `person`
* Then created `people` to be a `[]person` to handle the incoming json file which had multiple json objects

### JSON Encode/Decode

### Writer Interface
```
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
* Any other type with the `Write` method attached to it is also a `Writer`

### Sort
* `sort` package
```
import (
    "fmt",
    "sort"
)

testArray := []int{5,6,8,1}
sort.Ints(testArray)
fmt.Println(testArray)
```
* Returns `[1 5 6 8]`
```
import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)
}
```
* sort strings
* Returns `[Alpha Bravo Delta Go Gopher Grin]`

#### __CUSTOM SORTING (Sorting stucts)__
```
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
```
* Here what ive done is created a `byAge` type which is just an array of the `people` type
* Then I create 3 methods called `Len`, `Swap` and `Less` because an `Interface` type calls for those 3 functions. Remember to create an interface you specify a method and say any type with those methods also becomes another type
* So Im creating an `Interface` type to include the `byAge` type
* This is all done because `Sort` requieres an `Interface` type 
    * `func Sort(data Interface)`
### Bcrypt
* Increption
* from package `bcrypt`
* https://pkg.go.dev/golang.org/x/crypto/bcrypt
* Need to `go get golang.org/x/crypto`
    * "golang.org/x/crypto" directory comes from the go doc page (see link above)
    * `go get -u` will update
* Adds package to env because this isnt in the standard library




