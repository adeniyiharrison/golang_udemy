package main

import "fmt"

func qOne() {
	for x := 1; x <= 10000; x++ {
		fmt.Println(x)
	}
}

func qTwo() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d: %#U\t %#U\t %#U\n", i, i, i, i)
	}
}

func qThree() {
	x := 0
	for x <= 30 {
		fmt.Println(x)
		x++
	}
}

func qFour() {
	x := 0
	for {
		if x > 30 {
			break
		}
		fmt.Println(x)
		x++
	}

	fmt.Println("fin.")
}

func qFive() {
	for x := 10; x <= 100; x++ {
		fmt.Printf("%d: %d\n", x, x%4)
	}
}

func qSix() {
	if x := false; x {
		fmt.Println("wont print")
	} else {
		fmt.Println("X was false!")
	}
}

func qSeven() {
	if x := 1; x == 3 {
		fmt.Println("wont print")
	} else if x == 1 {
		fmt.Println("will print")
	} else {
		fmt.Println("wont print")
	}
}

func qEight() {
	switch {
	case true:
		fmt.Println("this will print")
	default:
		fmt.Println("this wont")
	}
}

func qNine() {
	favSport := "basketball"
	switch favSport {
	case "baseball":
		fmt.Println("nah fam")
	case "basketball":
		fmt.Println("yezzir")
	}
}

func main() {
	qNine()
}
