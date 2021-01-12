# Section 4 Notes

### Hello World
```
package main

import (
	"fmt"
)

var y = 42

func main() {
	NumberOfBytes, _ := fmt.Println("this is a test")
	fmt.Println(NumberOfBytes)
    fmt.Println(y)
	// "_" is throwing away the error statement returned (if there was an error)
	// This is being thrown away because you cannot declare a variable and not use it in go
    // No code polution

}
```
* Requirement for every program, every program will end when the main function is completed..
* https://godoc.org/fmt#Println
    * `fmt.Println(a...interface{})(n int, err error)`
        * `...<data_type>` variadic parameter. meaning it can take any number of arguements
        * `interface{}` empty interface meaning it could take any data type
        * `(n int, err error)` what data is returned from function

### Identifiers / Short declaration operator
* Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter
```
    identifier = letter { letter | unicode_digit} . 
```
* The following keywords are reserved and cannot be used as identifiers
    * `var`, `type`, `go`, `if`, `break`, `map`, `default`, etc.
* The first time you create a variable you must declaring it
    * Here we are declaring and assigning a new varaiable `x := 42`
    * later you can reasign a value to x using just `x = 0`
* variables
    * `var x = 42` OR `var z int`
    * defualt values
        * int = 0
        * string = ""
        * bool = false
        * pointers = nil
        * float = 0.0
    * used globaly outside of fuctions
    * BEST PRACTICE IS TO LIMIT THE SCOPE OF YOUR VARIABLES. aka use `var` when necessary
* statements
    * In programming a statement is the smallest standalone element of a program that expresses some action to be carried out. It is an instruction that commands the computer to perform a specified action. A progra, is formed by a sequence of one or more statements
* expression
    * in programming an expression is a combination of one or more explicit values, constants, variables, opereators and function that the programming language interprets and computes to produce another value. For example 2 + 3 is an expression that evaluates to 5

### Types
* Go is a STATIC programming language, once youve declared a variable to a particular type it will now will only accept that type
* `fmt.Printf("%T\n", z)`
    * Returns `int`
    * `\n` is just for a new line
* When decaring a string you can also used "`" (back qoutes) you can keep returns and double qoutes. Basically keeps all the formatting
* Primative Data Types
    * basic types or buit in type such as string, int, float, bool, etc
* Composite Data Types
    * Constructed using multple data types or other composite types
    * also called structure or aggregate data types
    * Slice, struct

### The fmt Package
* https://godoc.org/fmt
* `fmt.Printf("%#x\t%b\t%x", y, y, y)`
    * https://godoc.org/fmt#Printf
    * https://yourbasic.org/golang/multiline-string/
    * Here im printing the same y variable in different formats with tab spacing them out
* `s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)`
    * string print. Ive now made a variable with all the formatting turned into a string
* `Fprint`
    * prints to a file

### Creating your own type
* `type hotdog int`
    * hotdog now operates like an int
* Honestly not sure why you want to do this

### Conversion
* Convert type to another
```
    var a int
    type hotdog int
    var b hotdog
    func main() {
        b = 43
        a = int(b)
    }

```