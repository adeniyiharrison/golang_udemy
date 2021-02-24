# Ninja Level 9

1. 
    * In addition to the main goroutine, launch tow additional goroutines
        * Each additional Goroutine will print something out
    * Use waitgroups to make sure the goroutines finish before the program exits

2. This exercise will reinforce our understanding of method sets:
    * Create a type person struct
    * attach a method `speak` to type pointer of a person
    * create a human interface
        * human types must have a `speak` method
    * create func `saySomething`
        * have it take human as a param
        * have it call `speak`
    * show the following in your code
        * you can pass type `*person` into `saySomething`
        * you cannot pass type `person` into `saySomething`

3. Using goroutine, create an incrementer program
    * have a variable hold the incrementer value
    * launch a bunch of go routines
    * each routine should
        * reach the incrementer value
            * store it as a new variable
        * yield the processor with `runtime.Gosched()`
        * increment the new variable
        * write the value in the new variable back to the incrementor
    * use waitgroups to wait wait till all routines have finished
    * the above will create a race condition
    * prove that it is a race condition

4. Fix race condition using Mutex
5. Fix race condition using atomic