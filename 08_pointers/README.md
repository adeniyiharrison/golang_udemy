# Section 16 Notes

## Pointers
```
a := 69
fmt.Println(*a)
// returns 0x1040a124 which is the location the value is stored

fmt.Printf("%T\n", &a)
// returns *int as its type which is a pointer to an int

b := &a
fmt.Println(*b)
// Dereferrencing the pointer to convert back to value
```
* Pointing to some location in memory where a value is stored
    * `&` before the variable name will return the memory location
* `int` is a completely different type as `*int`
    * int vs pointer to an int
* For a value of `*int` type, I can convert it back to `int` by using the following syntax
    * type(`*b`) == `int`
* Not sure why you would want to do this but
    * `*&a` == `int`
    * Turn `int` into `pointer`, then dereference it with `*` to turn it back to `int`

```
type person struct {
	firstName string
	lastName  string
}

func main() {
	b := person{
		firstName: "Adeniyi",
		lastName:  "Harrison",
	}
	c := changeMe(&b)
	fmt.Println(b == c)

}

func changeMe(ogLocation *person) person {
	ogLocation.firstName = "Chloe"
	ogLocation.lastName = "Looker"

	return *ogLocation
}
```
* To dereference a struct use `(*value).fieldName`

## When to use Pointers
* Pointers are good when you have a big chunk of data that you dont want to pass around your program, you can just pass the address where its stored

## Method Sets
* Method sets determine what methods attach to a type
* __NON-POINTER RECIEVER__
    * Normal assigning of method for a type is using the syntax below
    * `func (r receiver) indetifier(parameters) returns {...}`
    * Works with values that are pointer or non-pointers
* __POINTER RECIEVER__
    * only works with values that are pointers
    * `func (r *receiver) indetifier(parameters) returns {...}`
    * `a := 5`, `b := identifer(&a)`
    
