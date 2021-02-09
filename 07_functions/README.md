# Section 14 Notes

## Functions
```
func (r reciever) identifier(parameter) (return(s)) {
    ...
}

func p1(s string) {
	fmt.Printf("Printing the arguement: %v\n", s)
}

func p2(s string) string {
	return fmt.Sprintf("Returned from p2: %v", s)
}

func p3(fn string, ln string) (string, bool) {

	fullName := fn + " " + ln
	nameCheck := fullName == "Adeniyi Harrison"

	return fullName, nameCheck
}
```

### Variadic Parameters
```
func p4(x ...int) {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Printf("%d: %d\n", i, v)
	}
	fmt.Println("Total Sum:", sum)
}
```
* you can pass any number of variable you'd like of a specified type
* `fmt.Printf(%T, x) == []int`
* Variadic means 0 or more. You can pass in no values for parameter if you'd like such as `p4()`
    * if no value is passed in for the variadic param then an empty slice will be created for the param within the function
* The variadic parameter must be the final parameter if there are multiple params for the fuction

__Unfurling a slice__
```
func main(){
	xi := []int{1, 2, 3, 4}
	p4(xi...)
}

func p4(x ...int) {
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Printf("%d: %d\n", i, v)
	}
	fmt.Println("Total Sum:", sum)
}
```
* Because `p4` is accepting any number of ints and NOT a slice of ints we must "unfurl"
* a slice plus `...` will unfurl a slice into its actual individual contents

### Defer
* A defer statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body or the corresponding GoRoutine is panicking.

```
func main() {
    p5()
}

func p1(s string) {
	fmt.Printf("Printing the arguement: %v\n", s)
}

func p5() {
	defer p1("Run this last")
	p1("Run this one first!")
}
```
* This will return
```
Printing the arguement: Run this one first!
Printing the arguement: Run this last
```
* Used to ensure a specific function run at the very end of a process. Such as closing files after a function that opens it

### Methods
```
type person struct {
	firstName string
	lastName  string
	age       int
}

type employee struct {
	person
	title string
}

func (e employee) employeeInfo() bool {
	fmt.Printf(
		"First: %v\n Last: %v\n Age: %d\n Occupation: %v\n",
		e.firstName,
		e.lastName,
		e.age,
		e.title)

	var nameBool bool

	if e.firstName+" "+e.lastName == "Adeniyi Harrison" {
		nameBool = true
	}

	return nameBool
}

func p6() {
	employee1 := employee{
		person: person{
			"Adeniyi",
			"Harrison",
			30},
		title: "Data Engineer",
	}

	isMe := employee1.employeeInfo()

	fmt.Println(isMe)
}
```
* Again the following is the function format!
```
func (r reciever) identifier(parameter) (return(s)) {
    ...
}
```
* Receiver being the type that the function belongs to
* In this case the `employee` type recieves the method of `employeeInfo` no other type has access to this method

### Interfaces & Polymorphism
* A value can be of more than one type 🤯
```
type human interface {
    employeeInfo()
}
```
* Any value that has the `employeeInfo` method will also be of type `human`

```
type man struct {
	gender string
}

type woman struct {
	gender string
}

func (m man) speak() {
	fmt.Println("Man speaks")
}

func (w woman) speak() {
	fmt.Println("Woman speaks")
}

type human interface {
	speak()
}

func functionHumanParam(h human) {
	switch h.(type) {
	case man:
		fmt.Println("Gender: ", h.(man).gender)
	case woman:
		fmt.Println("Gender: ", h.(woman).gender)
	}
	fmt.Println("I am still human")
}

func p7() {
	m := man{
		gender: "Male",
	}

	f := woman{
		gender: "Female",
	}

	functionHumanParam(m)
	functionHumanParam(f)
}
```
* Polymorphism (Change more than one)
    * Basically function that can operate on many types
    * See how `functionHumanParam` can accept both the `man` and `woman` types

__Assertion__
* In the `functionHumanParam` Im calling a switch statement (basically an if statement) on the type (`h.(type)`) of `human`. 
* `h.(woman).gender` gives you access to the parameter of the type

__Empty Interface__
* `type emptyInterface interface{}`
    * This says that every type will also become the `emptyInterface` type too
* For example if we look at the `fmt` package docs
    * `func Println(a...interface{}) (n int, err error)`
        * This function can accept any number of any type

### Anonymous Functions
```
func(){
    fmt.Println("Anonymous func ran")
}()

func(x int){
    fmt.Println("This is the number", 42)
}(68)
```
* Instead of creating stand alot functions you can define and run a function all in the same code block.
    * `()` the parenthesis at the end here call the functions
* AKA self executing function

__Function Expressions__
```
f := func(){
    fmt.Println("Function Expression, saving function as f var")
}

f()
```

### Returning a function
```
func main() {
    x := returnFunction()

    x()
}

func returnFunction() func() int {
    return func() int {
        return 911
    }
}
```
* This function returns a function which returns `int`
* Not sure why you would want to do this

### Callback
* Passing a function as an arguement
* Functional programming
```
func main() {
    ii :- []int{1,2,3,4,5,6,7,8}
    s := sum(ii...)
    fmt.Println(s)

    s2 := evenSum(sum, ii...)
    fmt.Println(s2)
}

func sum(xi...int) int {
    total := 0
    for _, v in range xi {
        total += v
    }

    return total
}

func evenSum(f func(xi...int) int, vi...int) int {
    var yi int
    for _, v := range vi {
        if v % 2 == 0{
            yi = append(yi, v)
        }
    }

    return f(yi...)
}
```

### Closure
```
func p8() {
	var x int
	x = 1
	x++
	fmt.Println(x)
	{
		var y int
		y = 100
		y++
		fmt.Println(y)
	}

	y++
}
```
* This fails because `y++` is out of scope
* enclosing a variable to limit its scope
```
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func p9() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
}
```