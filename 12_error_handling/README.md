# Error Handling

## Understanding
* Why does Go not have exceptions?
    * __We believe that coupling exceptions to a control structure, as in the "try, catch, finally" idiom results in convoluted code."__ 
    * It also encourages programmers to label too many originary errors, such as opening a file as exceptional.
    
* Go takes a different approach. For plain error handling, Go's multi-value returns make it easy to report an error without overloading the return value.
*  A canonical error type, coupled with Go's other features, makes error handling pleasant but quite different from that in other languages.
* Go also has a couple of built-in functions to signal and recover from truly exceptional conditions. The recovery mechanism is executed only as part of a function's state being torn down after an error, which is sufficient to handle catastrophe but requires no extra control structures and, when used well, can result in clean error-handling code.
* https://golang.org/doc/faq#exceptions

* `error` is a built in `interface` in GoLang
* Any type with the `Error()` method will also be an error type
* The predeclared type error is defined as
```
    type error interface {
        Error() string
    }
```
```
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}
```

```
    // New returns an error that formats as the given text.
    // Each call to New returns a distinct error value even if the text is identical.
    func New(text string) error {
        return &errorString{text}
    }
    
    // errorString is a trivial implementation of error.
    type errorString struct {
        s string
    }
    
    func (e *errorString) Error() string {
        return e.s
    }
```

## Checking Errors
```
if err != nil {
    // DO SOMETHING WHEN ERRORS EXIST
}
```
* you almost always want to be checking your errors

## Printing and Logging
You have a few options to choose from when it comes to printing out or logging an error message:
* `fmt.Println()`
    *  Println formats using the default formats for its operands and writes to standard output.
    ``` 
        func Println(a ...interface{}) (n int, err error) {
        return Fprintln(os.Stdout, a...)
    ```
* `log.Println()`
    * You get timestamp information
    * could write to a file
    ```
        logF, err := os.Create("temp/log.txt")
        checkError(err)
        log.SetOutput(logF) 

        f, err := os.Open("temp/logging_test.txt")
        if err != nil {
            log.Println(err)
	    }           
    ```
* `log.Fatalln()`
    * Fatalln is equivalent to `Println()` followed by a call to `os.Exit(1)`.
    * Game over, ends program __DEFERRED FUNCTIONS WILL NOT RUN__
* `log.Panicln()`
    * The panic built-in function stops normal execution of the current goroutine. When a function F calls panic, normal execution of F stops immediately. Any functions whose execution was deferred by F are run in the usual way, and then F returns to its caller. 
        * https://pkg.go.dev/builtin#panic
    * Panic is equivalent to `Print()` followed by a call to `panic()`
    * __DEFERRED FUNCTIONS WILL RUN__
    * can use "recover"

## Recover
* https://blog.golang.org/defer-panic-and-recover
* Recover is a built-in function that regains control of a panicking goroutine. __Recover is only useful inside deferred functions__. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
```
package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
```

## Errors with Info
```
func main() {
    n, err := sqrt(-10)
}

func sqrt(n float64) float64, error {
    if n > 0 {
        return math.sqrt(n), nil
    } else {
        return 0, error.New("cant calculate square root of negative number")
    }
}
```