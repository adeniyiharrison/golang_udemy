package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	// errorTest()
	// errorTest2()
	// ioTest()
	// ioTest2()
	// ioTest3()
	// ioTest4()
	// loggingTest()
	_, err := errorTest3(-10)
	if err != nil {
		log.Println(err)
	}

}

func errorTest() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}

func errorTest2() {
	const name, id = "adeniyi", 30
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	if err != nil {
		fmt.Print(err)
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func ioTest() {
	// need to be in proper directory when calling go run test.go
	// temp needs to be created
	file, err := os.Create("temp/test1.txt")
	checkError(err)

	_, err = file.Write([]byte("this is a test"))
	checkError(err)

	defer file.Close()
}

func ioTest2() {
	file, err := os.Create("temp/test2.txt")
	checkError(err)

	r := strings.NewReader("another test, writting using new reader")
	_, err = r.WriteTo(file)
	checkError(err)

	defer file.Close()
}

func ioTest3() {
	readFile, err := ioutil.ReadFile("temp/test1.txt")
	checkError(err)

	fmt.Println(string(readFile))
}

func ioTest4() {
	f, err := os.Open("temp/test2.txt")
	checkError(err)
	defer f.Close()

	readFile, err := ioutil.ReadAll(f)
	checkError(err)

	fmt.Println(string(readFile))
}

func ioTest5() {
	// this was me messing around // dont like this
	// https://gobyexample.com/reading-files
	readFile := make([]byte, 15)
	f, err := os.Open("temp/test1.txt")
	checkError(err)
	defer f.Close()

	_, err = f.Read(readFile)
	checkError(err)

	fmt.Println("Read File:", string(readFile))
}

func loggingTest() {

	logF, err := os.Create("temp/log.txt")
	checkError(err)
	log.SetOutput(logF)

	f, err := os.Open("temp/logging_test.txt")
	if err != nil {
		log.Println(err)
	}

	readFile, err := ioutil.ReadAll(f)
	fmt.Println(string(readFile))

	defer f.Close()
	defer logF.Close()
}

func errorTest3(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("cant calculate square root of negative number")
		// or we could use fmt.Errorf("test test one two")
		// you can do formatting with Errorf though
	}
	return math.Sqrt(n), nil
}
