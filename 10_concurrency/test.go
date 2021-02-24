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
	p4()
}

var wg sync.WaitGroup

func p1() {

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs \t\t", runtime.NumCPU())
	fmt.Println("GO Routines\t", runtime.NumGoroutine())

	wg.Add(1)
	go partOne()

	partTwo()
	fmt.Println("GO Routines\t", runtime.NumGoroutine())

	wg.Wait()
}

func partOne() {
	for i := 0; i < 10; i++ {
		fmt.Println("partOne:", i)
	}

	wg.Done()
}

func partTwo() {
	for i := 0; i < 10; i++ {
		fmt.Println("partTwo:", i)
	}
}

func p2() {

	var waitGroup sync.WaitGroup

	fmt.Println("GO Routines\t", runtime.NumGoroutine())

	const gs = 100
	counter := 0
	waitGroup.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			waitGroup.Done()
		}()
		fmt.Println("GO Routines\t", runtime.NumGoroutine())
	}
	waitGroup.Wait()
	fmt.Println("Counter \t\t", counter)
}

func p3() {

	var waitGroup sync.WaitGroup

	fmt.Println("GO Routines\t", runtime.NumGoroutine())

	const gs = 100
	counter := 0
	waitGroup.Add(gs)

	var mutex sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mutex.Unlock()
			waitGroup.Done()
		}()
		fmt.Println("GO Routines\t", runtime.NumGoroutine())
	}
	waitGroup.Wait()
	fmt.Println("Counter \t\t", counter)
}

func p4() {
	var waitGroup sync.WaitGroup

	fmt.Println("GO Routines\t", runtime.NumGoroutine())
	var counter int64
	const gs = 100
	waitGroup.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter \t", atomic.LoadInt64(&counter))
			waitGroup.Done()
		}()
		fmt.Println("GO Routines\t", runtime.NumGoroutine())
	}
	waitGroup.Wait()
	fmt.Println("Counter \t\t", counter)
}
