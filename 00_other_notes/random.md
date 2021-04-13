# Random Notes going through DSOM UP

## Difference using pointer in struct fields
* https://github.com/shipt/dsom-universal-products/blob/staging/pkg/models/models.go#L58
```
type PartnerDefault struct {
	PrimaryStoreID *int64     `json:"primary_store_id" db:"primary_store_id" permissions:"r,w" post:"nonzero" patch:"nonzero" put:"nonzero"`
	SkipOverride   *bool      `json:"skip_override" db:"skip_override" permissions:"r,w" post:"nonnil" patch:"nonnil" put:"nonnil"`
	CreatedAt      *time.Time `json:"created_at" db:"created_at" permissions:"r"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at" permissions:"r"`

	SQLHelpers base_sql.SQLHelpers `db:"-" json:"-"`
}
```
* https://stackoverflow.com/questions/59964619/difference-using-pointer-in-struct-fields
    ```
    type Foo struct {
        Bar string  `json:"bar"`
        Foo *string `json:"foo,omitempty"`
    }
    ```
    * When I'm unmarshalling a message that has no bar value, the Bar field will just be an empty string. That makes it kind of hard to work out whether or not the field was sent or not. Especially when dealing with integers: how do I tell the difference between a field that wasn't sent vs a field that was sent with a value of 0?
    * Using a pointer field, and specify omitempty allows you to do that. If the field wasn't specified in the JSON data, then the field in your struct will be nil, if not: it'll point to an integer of value 0.

## JSON struct tags
* https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go
*  When you read this information from systems such as databases, or APIs, you can use struct tags to control how this information is assigned to the fields of a struct. Struct tags are small pieces of metadata attached to fields of a struct that provide instructions to other Go code that works with the struct.
* The JSON encoder in the standard library makes use of struct tags as annotations indicating to the encoder how you would like to name your fields in the JSON output. These JSON encoding and decoding mechanisms can be found in the `encoding/json` [package](https://godoc.org/encoding/json)
```
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "time"
)

type User struct {
    Name          string    `json:"name"`
    Password      string    `json:"password"`
    PreferredFish []string  `json:"preferredFish"`
    CreatedAt     time.Time `json:"createdAt"`
}

func main() {
    u := &User{
        Name:      "Sammy the Shark",
        Password:  "fisharegreat",
        CreatedAt: time.Now(),
    }

    out, err := json.MarshalIndent(u, "", "  ")
    if err != nil {
        log.Println(err)
        os.Exit(1)
    }

    fmt.Println(string(out))
}
```
* This will output: (notice the camelCase)
```
-- Output
{
  "name": "Sammy the Shark",
  "password": "fisharegreat",
  "preferredFish": null,
  "createdAt": "2019-09-23T18:16:17.57739-04:00"
}
```
### Removing Empty
```
type User struct {
    Name          string    `json:"name"`
    Password      string    `json:"password"`
    PreferredFish []string  `json:"preferredFish,omitempty"`
    CreatedAt     time.Time `json:"createdAt"`
}
```
* Most commonly, we want to suppress outputting fields that are unset in JSON. 
### Ignoring private fields
```
type User struct {
    Name      string    `json:"name"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"createdAt"`
}
```