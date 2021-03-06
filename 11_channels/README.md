# Channel Notes

## Understanding Channels
```
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

```
* Channels allow use to pass values between goroutines
* P1 will not run because channels block
    * Theyre like runners in a relay race, they are sycrhonized
    * They have to pass/recieve the values at the same time
* P2 is good to go because channels are now unblocked because a goroutine is kicked off
    * Channels cannot handoff unless go routine is initiated to ensure read and write happen similtaneously
* P3 also runs because its a buffered channel, if you specify the number of items that can live on a channel it becomes unnlocked
* p4 is blocked because youve exceeded the size of the buffered channel
* You can create large channel buffers but you never know when you'll reach the limit. This is not good practice to create massive buffers
    * Bill Kennedy recommends not to use buffer channels

### __Do not communicate by sharing memory. Instead, share memory by communicating.__
* Channels allow you to pass references to data structures between goroutines
*  If you consider this as passing around ownership of the data (the ability to read and write it), they become a powerful and expressive synchronization mechanism. 

## Directional Channels
```
func setChannelValue(c chan<- int, i int) {
	c <- i
}

func recieveChannelValue(c <-chan int) {
	fmt.Println(<-c)
}

func directionalChannels() {
	// this channel can only set values
	setChannel := make(chan<- int)
	// this channel can only recieve values
	recieveChannel := make(<-chan int)
	// this can do both
	channel := make(chan int)

	fmt.Printf("setChannel:\t%T\n", setChannel)
	fmt.Printf("recieveChannel:\t%T\n", recieveChannel)

	go setChannelValue(channel, 69)

	recieveChannelValue(channel)
}
```
* It looks like the directional channels are to be used when defining functions. Basically saying in the function I should only be setting or reterieve
* The reason why `recieveChannelValue` is not a goroutine is to ensure that the channel is set before pulling anything off
* The following works too with wait groups
```
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
```
## Ranging Channels
```
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
```
* Not sure why but I thought channels could only accept one value at a time. Not true, this is only with buffered channels where limits are specified.

## Select Statement
```
func selectChannels() {
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	exitChannel := make(chan bool)

	go func() {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				evenChannel <- i
			} else {
				oddChannel <- i
			}
		}
		exitChannel <- true
	}()

	func() {
		for {
			select {
			case v := <-evenChannel:
				fmt.Println("This is Even:", v)
			case v := <-oddChannel:
				fmt.Println("This is Odd:", v)
			case v := <-exitChannel:
				if v {
					fmt.Println("Exit Channel has been hit")
					return
				}
			}
		}
	}()
}
```

### Comma OK idiom
* if `exitChannel` was `chan int` instead of `chan bool`, and instead of `exitChannel <- true` I put `close(exitChannel)`
    * in the select statement I could so something like `case i, ok := <-exitChannel`
    * this would return 0 for `i` and `false` for okay
* __Basically you can do the comma, okay idiom to check if a channel was closed__
    * Maybe the exit channel isnt required anymore because I could close the even or odd channel

```
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
```
* Also if you pull the last value off a closed channel, the nex time you try to pull off you should get a false for `ok`

```
c := make(chan int)
go func(){
    c <- 100
    close(c)
}()

i, ok := <-c
fmt.Println(i, ok)
// returns 100, true

i, ok = <-c
// returns 0, false
```

## Fan in
* Taking values from many channels and put those values onto one channel
```
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
```

## Fan Out
* Taking some work and putting the chunks of work onto many goroutines

```
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
```

## Context
* Is a tool that goes along with concurrent design patterns that allows you to cancel all associated go routines if you cancel it's parent process
* You dont want to "leak goroutines" // waste of resources
* https://blog.golang.org/context
* https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922360#overview
```
In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional goroutines to access backends such as databases and RPC services. The set of goroutines working on a request typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the request's deadline. When a request is canceled or times out, all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using.

At Google, we developed a context package that makes it easy to pass request-scoped values, cancelation signals, and deadlines across API boundaries to all the goroutines involved in handling a request. The package is publicly available as context. This article describes how to use the package and provides a complete working example.
```


