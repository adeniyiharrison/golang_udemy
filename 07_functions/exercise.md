# Ninja Level 6

1.
    * Create a function with identifier q1f1 that returns and int
    * Create a function with identifier q1f2 that returns an int and a string
    * Call both functions
    * Print out their results

2.
    * Create a function with identifier q2f1 that
        * takes variadic parameter of type int
        * pass in a value of type []int that you unfurl
        * returns the sum of all values of type int
    * Create a function with identifier q2f2 that
        * takes in a parmeter of type []int
        * returns the sum of all values of type int passed in

3. Use the `defer` keyword to show that the deferred function runs after the func containing it exits

4. Create a user defined `struct` with
    * the `identifier` "person"
    * the fields
        * `first`
        * `last`
        * `age`
    * Attach a method to type person with 
        * the `identifier` "speak"
        * the method should have the person say their name and age
    * Create a value of type person
    * Call the method from the value of type person

5. 
    * create  type `square` and `circle`
    * attach a method to each that calculates the area and returns it
    * create a type `shape` which defines an interface as anything which has the area method
    * create a function `info` which takes type shape and then prints the area
    * create a value of the type `square` and `circle`
    * use func info to print the area of the `square` and `circle`

6. Build and use an anonymous function
7. Create a function expression
8.
    * Create a func which returns a func
    * assign the returned func to a variable
    * call the returned func
9. Pass a func into a func as an arguement
