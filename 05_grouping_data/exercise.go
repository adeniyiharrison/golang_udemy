package main

import "fmt"

func p1() {
	// var testArray [5]int
	// testArray[1] = 1
	// testArray[2] = 2
	// testArray[3] = 3
	// testArray[4] = 4

	testArray := [5]int{0, 1, 2, 3, 4}
	// defining lenght from the get go

	for i, v := range testArray {
		fmt.Printf("%d:\t%d\n", i, v)
	}

	fmt.Printf("%T\n", testArray)
}

func p2() {
	testArray := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range testArray {
		fmt.Printf("%d:\t%d\n", i, v)
	}

	fmt.Printf("%T\n", testArray)
}

func p3() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 41}
	s1 := x[:5]
	s2 := x[5:]
	s3 := x[2:7]
	s4 := x[1:6]

	fmt.Printf("%v\t%v\t%v\t%v\n", s1, s2, s3, s4)
}

func p4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	x = append(x, []int{56, 57, 58, 59, 60}...)

	fmt.Println(x)
}

func p5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)

	fmt.Println(x)
}

func p6() {
	x := make([]string, 50, 50)
	fmt.Printf("%T\n", x)
	fmt.Print(len(x))
	fmt.Print(cap(x))
}

func p7() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Martini"}
	xp := [][]string{x, y}

	fmt.Printf("%T\n", xp)
	fmt.Println(xp)
}

func p8() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Women"},
		"miss_moneypenny": []string{"Bond", "literature"},
		"dr_no":           []string{"Being Evil", "Ice cream"}}

	for k, v := range m {
		fmt.Println("Key: ", k)
		for i, v := range v {
			fmt.Printf("\ti:%d\tv:%v\n", i, v)
		}
	}
}

func p9() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Women"},
		"miss_moneypenny": []string{"Bond", "literature"},
		"dr_no":           []string{"Being Evil", "Ice cream"}}

	m["odd_job"] = []string{
		"Being Evil too", "Beaches", "Hats"}

	for k, v := range m {
		fmt.Println("Key: ", k)
		for i, v := range v {
			fmt.Printf("\ti:%d\tv:%v\n", i, v)
		}
	}
}

func p10() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Women"},
		"miss_moneypenny": []string{"Bond", "literature"},
		"dr_no":           []string{"Being Evil", "Ice cream"}}

	m["odd_job"] = []string{
		"Being Evil too", "Beaches", "Hats"}

	delete(m, "dr_no")

	for k, v := range m {
		fmt.Println("Key: ", k)
		for i, v := range v {
			fmt.Printf("\ti:%d\tv:%v\n", i, v)
		}
	}
}

func main() {
	p10()
}
