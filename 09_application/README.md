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
