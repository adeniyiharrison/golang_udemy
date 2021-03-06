package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {
	// p1()
	// p2()
	// p3()
	// p4()
	// p5()
	p6()
}

func p1() {
	c := make(chan int)
	c2 := make(chan int, 1)

	go func() {
		c <- 1
	}()

	c2 <- 2

	fmt.Println("method #1:", <-c)
	fmt.Println("method #2:", <-c2)
}

func p2() {
	cs := make(chan int)
	go func() {
		cs <- 1
	}()
	fmt.Println(<-cs)
	fmt.Printf("%T\n", cs)
}

func p3() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for value := range c {
		fmt.Println(value)
	}

	// for {
	// 	select {
	// 	case value, ok := <-c:
	// 		if ok {
	// 			fmt.Println(value)
	// 		} else {
	// 			return
	// 		}
	// 	}
	// }
}

func p4() {
	q := make(chan int)
	gen(q)
	recieve(q)
	fmt.Println("exit func")
}

func gen(c chan<- int) {
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
}

func recieve(c <-chan int) {
	for {
		select {
		case value, ok := <-c:
			if ok {
				fmt.Println(value)
			} else {
				fmt.Println("exit select")
				return
			}
		}
	}
}

func p5() {
	channel := make(chan int)

	go func(c chan<- int) {
		for i := 0; i < 100; i++ {
			c <- rand.Intn(100)
		}
		close(c)
	}(channel)

	func(c <-chan int) {
		for {
			select {
			case value, ok := <-c:
				if ok {
					fmt.Println(value)
				} else {
					return
				}
			}
		}
	}(channel)
}

func p6() {

	threads := 10
	var waiter sync.WaitGroup

	channel := make(chan int)
	go func() {
		for i := 0; i < threads; i++ {
			fmt.Println("Number of Routines:", runtime.NumGoroutine())
			waiter.Add(1)
			go func() {
				for j := 0; j < 10; j++ {
					channel <- rand.Intn(100)
				}
				waiter.Done()
			}()
		}
		waiter.Wait()
		close(channel)
	}()

	for v := range channel {
		fmt.Println(v)
	}
}
