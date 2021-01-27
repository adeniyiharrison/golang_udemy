# Section 10 Notes

### Arrays
* `var x [5]int`
* Arrays are useful when planning the detailed layout of memory and sometimes can help avoid allocation, but primarily they are a building block for slices, the subject of the next section
* The size of an array is part of its type. The types `[10]int` and `[20]int` are distinct.
    * size of array must be defined
    * All one type
    * LENGTH IS PART OF ITS TYPE!
* Slices are used mostly, arrays have gotten the back seat

### Slices
```
func p2() {
	// x := type{values} // composite literal
	x := []int{4, 5, 6, 7, 8, 9}
	fmt.Printf("Total Lenght of x: %d\n", len(x))

	// BRUTE FORCE
	n := 0
	for {
		if n < len(x) {
			fmt.Printf("n=%d:\t %d\n", n, x[n])
			n++
		} else {
			break
		}
	}

	// FINESCE

	for i, v := range x {
		fmt.Printf("i=%d:\t %d\n", i, v)
	}

	fmt.Println("fin.")

}
```
* A slice allows you to group together values of the same type
    * Like arrays but length isnt specified
* Looping over a slice
    *  `for i, v := range x {}`
* Slicing a slice
    * x[0:2] == [4,5]
    * last number of range is not inclusive (unlike python)
* Appending a slice
    ```
    func p3() {
        x := []int{1, 2, 3}
        y := []int{7, 8, 9}

        x = append(x, 4, 5, 6)
        x = append(x, y...)
        fmt.Println(x)

        x = append(x[:4], x[5:]...)
        fmt.Printf("Deleted the 5: %d\n", x)
    }
    ```
    * `...` occurring after a slice will UNFURL the slice into its inividual numbers
    * Apeend to slices together with index removed to delete from slices
### Make
* Built in function
* you can also make slices with `make`, it is more efficient for the compilier but you need to know the length of the slice you want to make
* `x = make([]int, 10, 12)`
    * make(type, length, capactity)
    * `len(x) == 10`
    * `cap(x) == 12`
    * CAPACITY IS THE SIZE OF THE UNDERLYING ARRAY, You still can append to a make slice but if you append more than 12 that underlying array will double to 24

### Multi-Dimensional Slice
```
func p4() {
	one := []string{
		"Adeniyi",
		"Harrison",
		"Oakland",
		"CA"}

	two := []string{
		"Chloe",
		"Looker",
		"San Francisco",
		"CA"}

	xp := [][]string{
		one,
		two}

	fmt.Println(xp)
}
```
* `[[Adeniyi Harrison Oakland CA] [Chloe Looker San Francisco CA]]`

### Map
```
func p5() {
	m := map[string]int{
		"Adeniyi": 30,
		"Chloe":   29,
		"Olawale": 31,
		"Amanda":  33}

	fmt.Println(m)
	fmt.Println(m["Adeniyi"] == 30)
	_, ok := m["Olukemi"]
	fmt.Println(ok)

	if v, ok := m["Gina"]; ok {
        // called a 'comma okay' idiom
		fmt.Println("This is very common logic for accessing map: ")
	} else {
		fmt.Printf("%d: %t\n", v, ok)
	}
}
```
* Maps are a convenient and powerful built-in data structure that associate values of one type (the key) with values of another type (the element or value). The key can be of any type for which the equality operator is defined, such as integers, floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs and arrays.
* Like python dictionary
* if key is not in map the defualt return value is `0`
    * If you store results returned from the map then you can check if the values are `ok`
* Adding to map
    * `m["Micheal"] = 70`
* Loop Range over map
    ```
        for k, v := range m {
            fmt.Println(k, v)
        }
    ```
* Delete from map
    * `delete(m, "Olawale")`
    * If you delete a key that doesnt exist, no errors are thrown :)





