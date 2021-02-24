package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// p1()
	// p2()
	// p3()
	// p4()
	p5()
}

func p1() {

	var wg sync.WaitGroup

	fmt.Println("Starting Routines:", runtime.NumGoroutine())
	wg.Add(2)

	go func() {
		fmt.Println("Go Routines After 1st func:", runtime.NumGoroutine())
		wg.Done()
	}()
	go func() {
		fmt.Println("Go Routines After 2nd func:", runtime.NumGoroutine())
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Finishing Routines:", runtime.NumGoroutine())

}

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) speak() {
	fmt.Printf("Human says 'I am %v'\n", p.firstName)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func p2() {
	me := person{
		firstName: "Adeniyi",
		lastName:  "Harrison",
		age:       30,
	}

	// these both work!
	// but if the speak method had a receiver of (p *person)
	// then saySomething(me) wouldnt work
	saySomething(me)
	saySomething(&me)
}

func p3() {
	var counter int
	var waitGroup sync.WaitGroup
	const totalRoutines = 100
	waitGroup.Add(totalRoutines)

	for i := 0; i < totalRoutines; i++ {
		go func() {
			runtime.Gosched()
			value := counter
			value++
			counter = value
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Println("Counter:", counter)

}

func p4() {
	var counter int
	var waitGroup sync.WaitGroup
	var mu sync.Mutex
	const totalRoutines = 100
	waitGroup.Add(totalRoutines)

	for i := 0; i < totalRoutines; i++ {
		go func() {
			mu.Lock()
			value := counter
			value++
			counter = value
			waitGroup.Done()
			mu.Unlock()
		}()
	}
	waitGroup.Wait()
	fmt.Println("Counter:", counter)

}

func p5() {
	var counter int64
	var waitGroup sync.WaitGroup
	const totalRoutines = 100
	waitGroup.Add(totalRoutines)

	for i := 0; i < totalRoutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter:", atomic.LoadInt64(&counter))
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Println("Final Counter:", counter)

}
