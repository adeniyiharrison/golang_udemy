package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// p1()
	// p2()
	// p3()
	// p4()
	// p5()
	// p6()
	// directionalChannels()
	// rangingChannels()
	// selectChannels()
	// fanInFunc()
	fanOutFunc()
}

func p1() {
	// fatal error: all goroutines are asleep - deadlock!
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)
}

func p2() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}

func p3() {
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
}

func p4() {
	// This fails
	c := make(chan int, 1)
	c <- 42
	c <- 43
	fmt.Println(<-c)
}

func p5() {
	c := make(chan int, 2)
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func p6() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
	go func() {
		c <- 43
	}()
	fmt.Println(<-c)
}

var wg sync.WaitGroup

func setChannelValue(c chan<- int, i int) {
	c <- i
	wg.Done()
}

func recieveChannelValue(c <-chan int) {
	fmt.Println(<-c)
	wg.Done()
}

func directionalChannels() {
	wg.Add(2)
	// this channel can only set values
	setChannel := make(chan<- int)
	// this channel can only recieve values
	recieveChannel := make(<-chan int)
	// this can do both
	channel := make(chan int)

	fmt.Printf("setChannel:\t%T\n", setChannel)
	fmt.Printf("recieveChannel:\t%T\n", recieveChannel)

	go setChannelValue(channel, 69)
	go recieveChannelValue(channel)

	wg.Wait()
}

func rangingChannels() {
	channel := make(chan int)

	go func() {
		// channel <- 1
		// channel <- 2
		// channel <- 3
		for i := 0; i < 5; i++ {
			channel <- i
		}
		close(channel)
	}()

	for v := range channel {
		fmt.Println(v)
	}

	fmt.Println("done")
}

func setChannelFunc(
	evenChannel,
	oddChannel,
	exitChannel chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			evenChannel <- i
		} else {
			oddChannel <- i
		}
	}
	exitChannel <- 0
}

func selectChannelFunc(
	evenChannel,
	oddChannel,
	exitChannel <-chan int) {
	for {
		select {
		case v := <-evenChannel:
			fmt.Println("This is Even:", v)
		case v := <-oddChannel:
			fmt.Println("This is Odd:", v)
		case v := <-exitChannel:
			fmt.Println("Exit Channel has been hit", v)
			return
		}
	}
}

func selectChannels() {
	evenChannel := make(chan int)
	oddChannel := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				evenChannel <- i
			} else {
				oddChannel <- i
			}
		}
		close(evenChannel)
		close(oddChannel)
	}()

	func() {
		for {
			select {
			case i, ok := <-evenChannel:
				if ok {
					fmt.Println("This is Even:", i)
				} else {
					fmt.Println("exiting from even chan")
					return
				}
			case i, ok := <-oddChannel:
				if ok {
					fmt.Println("This is Odd:", i)
				} else {
					fmt.Println("exiting from odd chan")
					return
				}
			}
		}
	}()
}

func receive(
	evenChannel, oddChannel <-chan int,
	fanChannel chan<- int) {
	wg.Add(2)
	go func() {
		for x := range evenChannel {
			fanChannel <- x

		}
		wg.Done()
	}()

	go func() {
		for x := range oddChannel {
			fanChannel <- x

		}
		wg.Done()
	}()

	wg.Wait()
	close(fanChannel)
}

func fanInFunc() {

	evenChannel := make(chan int)
	oddChannel := make(chan int)
	fanChannel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				evenChannel <- i
			} else {
				oddChannel <- i
			}
		}
		close(evenChannel)
		close(oddChannel)
	}()

	// go receive(
	// 	evenChannel,
	// 	oddChannel,
	// 	fanChannel)

	go func() {
		wg.Add(2)
		go func() {
			for x := range evenChannel {
				fanChannel <- x

			}
			wg.Done()
		}()

		go func() {
			for x := range oddChannel {
				fanChannel <- x

			}
			wg.Done()
		}()

		wg.Wait()
		close(fanChannel)
	}()

	for x := range fanChannel {
		fmt.Println("Fan_Channel:", x)
	}
}

func fanOutFunc() {
	inputChannel := make(chan string)
	outputChannel := make(chan string)
	people := []string{
		"Adeniyi",
		"Chloe",
		"Ola",
		"Amanda",
		"Teddy",
		"Gina",
		"Peter",
		"Kemi",
		"Micheal",
	}

	go func() {
		for _, v := range people {
			inputChannel <- v
		}
		close(inputChannel)
	}()

	go func() {
		fmt.Println("Goroutines Launched")
		for person := range inputChannel {
			wg.Add(1)
			go func(p string) {
				randInt := int(rand.Intn(10))
				waitTime := fmt.Sprintf("%s will wait for %d seconds", p, randInt)
				time.Sleep(time.Duration(randInt) * time.Second)
				outputChannel <- waitTime
				wg.Done()
			}(person)
		}
		wg.Wait()
		close(outputChannel)
	}()

	for {
		select {
		case waitTime, ok := <-outputChannel:
			if ok {
				fmt.Println(waitTime)
			} else {
				fmt.Println("end of program")
				return
			}
		}
	}

}
