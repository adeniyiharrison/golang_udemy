# Section 5 Notes

### Exercise #1
* Using short declaration, assign the values to variables with identifiers `x`, `y`, and `z`.
    * 42
    * James Bond
    * true
* Now print the values stored in those variables using
    * single print statement
    * multiple print statements

### Exercise #2
* Use var to DECLARE three variables. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the follow TYPE:
    * identifier `x` type int
    * identifier `y` type int
    * identifier `z` type int

* in `func main`
    * print out the values for each identifier
    * the compiler assigned values to the variables. What are these values called? 

### Exercise #3
* At the package level scope, assign the following values to the three variables a. for `a` assign 42 b. for `b` assign "James Bond" c. for `c` assign true
* in `func main` a. use `fmt.Sprintf` to print all of the VALUES to one single string. ASSIGN the returned value of type string using the short declaration operator to a VARIABLE with the IDENTIFIER `s`

### Exercise #4
* create your own type. Have the underlying type be an `int`.
* create a VARIABLE of your new TYPE with the IDENTIFIER `d` using the `var` keyword
* in `func main`
    * print out the value of the variable `d`
    * print out the type of the variable `d`
    * assign `42` to the variable `d` using the `=` OPERATOR
    * print out the value of the variable `d`

### Exercise #5
* at the package level scope, using the `var` keyword, create a VARIABLE with the IDENTIFIER `testType`. The variable should be of the UNDERLYING TYPE of your custom type `f`.
* in `func main`
    * this should already be done
        * print out the value of the variable `f`
        * print out the type of the variable `f`
        * assign your own VALUE to the VARIABLE `f` using the `=` OPERATOR
        * print out the value of the variable `f`
    * now do this
        * now use CONVERSION to convert the TYPE of the VALUE stored in `f` to the UNDERLYING TYPE
        * then use the short declaration operator to ASSIGN that value to `testType`

### Exercise #6
1. The smallest standalone element of a program that expresses some action to be carried out.  

  - [x] statement
  - [ ] expression

2. A combination of one or more explicit values, constants, variables, operators, and functions that the programming language interprets and computes to produce another value. 

  - [ ] statement
  - [x] expression

3. Which are "parentheses" or "parens?"

  - [x] ()
  - [ ] {}
  - [ ] []

4. Which are "curly braces" or "curlies" or "braces?"

  - [ ] ()
  - [x] {}
  - [ ] []

5. Which are "brackets?"  

  - [ ] ()
  - [ ] {}
  - [x] []

6. The "scope" of a variable is where you can access the variable, eg, write to it or read the value from it. *

  - [x] true
  - [ ] false

7. In Go, an `int` is a primitive data TYPE

  - [x] true
  - [ ] false

8. In Go, a `string` is a primitive data TYPE

  - [x] true
  - [ ] false

9. A "composite" data TYPE allows you to compose together values of other data TYPES

  - [x] true
  - [ ] false

10. When a variable is declared in Go using the `var` keyword, and no VALUE is ASSIGNED to that variable, then the compiler assigns a default value to the variable. This is known as the "zero value"

  - [x] true
  - [ ] false

11. Keywords are words that a reserved for use by the Go programming language; they have to be used in a certain way for a certain purpose.

  - [x] true
  - [ ] false

12. Keywords are sometimes called "reserved words."  

  - [x] true
  - [ ] false

13. You can’t use a keyword for anything other than its purpose.  

  - [x] true
  - [ ] false

14. In `2 + 2` the `+` is the OPERATOR

  - [x] true
  - [ ] false

15. In `2 + 2` the `2`s are OPERANDS

  - [x] true
  - [ ] false

16. For finding documentation, what is the difference between documentation found at [golang.org](https://golang.org/) and [godoc.org](http://godoc.org/)?  

17. `package` is a keyword

  - [x] true
  - [ ] false

18. `var` is a keyword

  - [x] true
  - [ ] false

19. The entry point for all programs is in `func main()` which needs to be inside package main

  - [x] true
  - [ ] false

20. The "short declaration operator" can be used anywhere in a program, including at both the package level and at the block level.  

  - [ ] true
  - [x] false

21. What are the three words used to describe good package names in the "effective go" document?  

  - [x] descriptive
  - [x] short
  - [x] concise
  - [ ] evocative

22. What is the name of the website where you can write (most) Go code online and have it run online?  

23. A great place to ask questions is the "golang bridge forum" at https://forum.golangbridge.org/  

  - [ ] true
  - [ ] false

24. When you see something like `fmt.Println()` this is calling the `Println()` function from the `fmt` package.  

  - [x] true
  - [ ] false

25. An "identifier" is the name assigned to a variable or a function or a constant.  

  - [x] true
  - [ ] false

26. To call a func, variable, or constant from a package, use the "package-dot-identifier" syntax. For example, like this, `fmt.Println()`  

  - [ ] true
  - [x] false

27. What is "idiomatic Go code"?


28. Which character allows you to "throw away returns" or "send returns into the void"? Said another way, which character allows you to tell the compiler that you are not going to use a value returned by a function?  

  - [ ] #
  - [ ] @
  - [x] _
  - [ ] This is a trick question

29. In Go, you cannot have a variable which you do not use.  

  - [x] true
  - [ ] false

30. When you see that a func has a parameter of this type `...interface{}` this is called a "variadic parameter" and it means that the func can take as many values of that type as you want to pass in.  

  - [x] true
  - [ ] false

31. Every value in Go is also of type "empty inerface" which is expressed like this: `interface{}`  

  - [x] true
  - [ ] false

32. A statement is an instruction that commands the computer to perform a specified action. Usually statements take up a line in a program. 

  - [ ] true
  - [ ] false

33. An expression is a combination of one or more explicit values, constants, variables, operators, and functions that the programming language interprets and computes to produce another value. For example, 2+3 is an expression which evaluates to 5.  

  - [x] true
  - [ ] false

34. If I wanted to print to a string and then assign that value to a variable, I could use the `func Sprintf()` from the `fmt` package.

  - [x] true
  - [ ] false

35. In Go, you can create your own TYPE

  - [x] true
  - [ ] false

36. We don't say "casting" in Go, we say "conversion"  

  - [x] true
  - [ ] false

37. There is a language which we use to talk about the language.

  - [x] true
  - [ ] false

38. When you create our own TYPE in Go, that TYPE will have an "underlying TYPE".

  - [ ] true
  - [ ] false