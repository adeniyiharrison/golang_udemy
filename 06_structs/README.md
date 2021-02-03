# Section 12 Notes

## Structs
```
type person struct {
    first, last string,
    age int,
}
```
* Data Structure
    * Data structure with values of differing types
    * Refered to as aggregate datatype, complex datatype
* Fields of the same type can be declared on the same line

```
type employee struct {
	person
	title  string
	salary int
}

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
```
* Embedded Structs
    * "Class Inheritance" for python
    * Object Oriented Programming
* `person: person{}`
    * This is called an embedded type or anonymous field
* When the `person` struct is inherited into the `employee` struct you would say "the `person` struct is promoted"

```
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
```
* Called an anonymous struct, defined on the fly
