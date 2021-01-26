package main

import "fmt"

func testLoop() {
	for i := 0; i < 100; i++ {
		fmt.Printf("i = %d\n", i)
	}
}

func testLoop1() {
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
}

func testLoop2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Outer Loop = %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t Inner Loop = %d\n", j)
		}
	}
}

func testLoop3() {
	x := 0
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
}

func testLoop4() {

	for i := 33; i <= 122; i++ {
		fmt.Printf("%d: %#U\n", i, i)
	}
}

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
	// fmt.Println(x)
}

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

func testLoop8() {
	switch {
	case false:
		fmt.Println("This wont print")
	case false:
		fmt.Println("Not true so wont print")
	case false:
		fmt.Println("This will print!")
	default:
		fmt.Println("If nothing is true, default runs")
	}
}

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

func main() {
	testLoop9()
}
