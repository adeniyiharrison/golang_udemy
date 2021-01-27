package main

import (
	"fmt"
)

func p1() {
	var x [5]int
	fmt.Printf("%T: %d\n", x, x)
}

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
}

func p3() {
	x := []int{1, 2, 3}
	y := []int{7, 8, 9}

	x = append(x, 4, 5, 6)
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:4], x[5:]...)
	fmt.Printf("Deleted the 5: %d\n", x)
}

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
		fmt.Println("This is very common logic for accessing map: ")
	} else {
		fmt.Printf("%d: %t\n", v, ok)
	}
}

func main() {
	p5()
}
