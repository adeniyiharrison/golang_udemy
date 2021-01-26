# Section 8 Notes

### Understanding Control Flow
* Control Flow: In CS control flow is the order in which individual statements, instructions or function calls of an imperative program are executed or evaluated. (Order of Operations)

* Helpful Resource for Syntax: www.gobyexample.com

## Loops
```
    for init; condition; post{
        fmt.Println("This is how you make a loop)
    }
```
```
for i := 0; i < 100: i++ {
    fmt.Printf("i = %d\n", i)
}
```
* __Kind of like a while loop__
* `init`: initialize a variable
    * `i := 0`
* `condition`: while `i` matches a condition the loop will continue to run
    * `i < 100`
* `post`: what to do after each loop logic is completed
    * `i ++`
    * add one to variable `i`

```
func testLoop1() {
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
}
```
* this works as well

### Nested Loops
```
func testLoop2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Outer Loop = %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t Inner Loop = %d\n", j)
		}
	}
}
```

### Break
```
func testLoop3() {
	x := 0
	for {
		if x > 9 {
			break
		}
		fmt.Print(x)
		x++
	}
	fmt.Println("done")
}
```
* This loop is a forever loop if the `break` in the if statement wasnt there
* `break` ends a loop. Remove you from loop
* For ever loop is effectively a listener

### Continue
* `continue` allows code to skip to next iteration of loop

## If Statements
```
func testLoop5() {
	if true {
		fmt.Println("This is constantly true")
	}

	if !false {
		fmt.Println(
			"This is also true, because the 'not' operator")
	}

	if x := 42; x == 42 {
		fmt.Println(`
		Initialize statement within if statement.
		The scope of x is limited to this if statement.
		Which is why the print statement below doesnt work`)
	}
	fmt.Println(x)
}
```
```
func testLoop6() {

	for x := 40; x <= 42; x++ {

		if x == 40 {
			fmt.Println("Almost there")
		} else if x == 41 {
			fmt.Println("Pretty close!")
		} else if x == 42 {
			fmt.Printf("Finally we made it to %d!\n", x)
		} else {
			fmt.Println("Something has gone wrong!")
			break
		}
	}
	fmt.Println("Exiting Function")
}
```

## Switch
```
func testLoop7() {
	switch {
	case false:
		fmt.Println("This wont print")
	case 2 != 2:
		fmt.Println("Not true so wont print")
	case true:
		fmt.Println("This will print!")
	case 2 == 2:
		fmt.Println("This is also true, but why doesnt it print??")
		// By defualt, there isnt "fallthrough"
	}
}
```

```
func testLoop9() {
	n := 1
	switch n {
	case 0:
		fmt.Println("Miss")
	case n:
		fmt.Println("We gotta match!")
	case 2:
		fmt.Println("Not true so wont print")
	default:
		fmt.Println("wont print")
	}
}
```

* Adding "fallthrough" to the case statements will ensure everything underneath will run. Regardless if its true or not.. Funky logic in my opinion
* If all case statements are all false then the `default` runs
* Also not sure how this really differs from an if statement

## Conitional Operators
* `&&` == "AND"
* `||` == "OR"
