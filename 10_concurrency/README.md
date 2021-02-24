# Concurrency Notes

## Concurrency vs parallelism
* Go is the first major programming language to natively take advantage of multiple cores
* If you run a Go script on a machine with one CPU there is no way that code can run in parallel, itll run sequentially. However if you have more than one, you have the OPPORTUNITY for code to run in parallel. 
* Concurrency is a design pattern, it is a way that you can write your code. If you have multiple cores, the concurrent code can now run in PARALLEL
* Writing concurrent code doesnt guarantee that your code is ran in parallel, what will now determine this is having a machine with multiple cores (CPUs)

## WaitGroup
```
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
```
* `runtime.NumCPU()` number of CPUs on machine
* add `go` in front of the calling of a function to create another "go routine"
* created WaitGroup type `var wg sync.WaitGroup`
    * https://pkg.go.dev/sync#WaitGroup
        * AS you can see `WaitGroup` is a type
* Then you let the program know that we are waiting for `1` goroutine to complete
    * `wg.Add()`
* You let the program know that the goroutine is completed with adding this to the end of the goroutine
    * `wg.Done()`
* And finally add `wg.Wait()` to the program where you'd like to wait for the goroutine to complete

## Method Sets Revisited
* You can see that these methods are attached to the `WaitGroup` type, ACTUALLY a pointer to the `WaitGroup` (wg *WaitGroup)
    * https://go.googlesource.com/go/+/go1.16/src/sync/waitgroup.go#53
* Well then how was I able to do `wg.Add(1)` above?
    * wg was just of type `WaitGroup` not actually a pointer
* Pointers to a type work for methods with recievers like (pointer to a type work for either method recievers)
    * `func (t T) testFunction(){...}`
    * `func (t *T) testFunction(){...}`
* BUT a reciever that calls for a pointer should be only be able to recieve a pointer!
    * `func (t *T) testFunction(){...}` can only accept a pointer
* AGAIN, how was I able to do `wg.Add(1)` above??
* __The method set of a type determines the interfaces that the type implements..__
    * https://golang.org/ref/spec#Method_sets

```
type circle struct {
    radius float64
}

func (c *circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

type shape interface {
    area() float64
}

func info(s shape) {
    fmt.Println("area", s.area())
}

func main() {
    c := circle{
        "radius": 5,
    }

    info(c)
}
```
* This will return the following error
    * `cannot use c (type circle) as type shape in argument to info; cicle doesnt implement shape (area method has pointer reciever)`
    * This is because the interface is not implmented correctly
* `c.area()` works just fine because the method to the `circle` type was implemented
* If I were to delete the pointer to cirlce from the `area()` definition it would run
* Honestly didnt fully understand this and it felt like a tangent for Todd. Just be careful when creating methods

## Additional Documentation
* __Concurrent programming in many environments is made difficult by the subleties required to implement correct access to shared variables__
    * Struggle to share access to variables on different threads, etc.
    * I know this all too well
* "Race Condition" is when the program gets wonky because the order of operations gets screwed up.
    * With shared variables, the read and write causes problems because different go routines are accessing variables simultaneously
* Go encourages a differnt approach in which shared values are passed around on channels and in fact, never actively shard by separate threads of execution. Only one goroutine has access to a value at a given time. Data races cannot occur by design
* They're called `goroutines` because the existing terms, threads, coroutines, processes and so on-convey inaccurate connotations. 
    * __A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space.__
    * It is lightweight, costing little more than the allocation of stack space.
* __Goroutines are multiplexed onto multiple OS threads so if one should block, such as while waiting for I/O, others continue to run. Their design hides many of the complexities of thread creation and management.__
* Prefix a function or metho with the `go` keyword to run the call in a new goroutine. When the call completes the goroutine exits silently. 

```
func doSomething(x int) int {
    return x * 5
}

func main() {
    ch : make(chan int)
    go func() {
        ch <- doSomething(5)
    }()
    fmt.Println(<-ch)

}
```
* __If the function has any return values, they are discarded when the function completes__
    * You generally will not call a goroutine on a function or method that returns a variable.
* So you can we called a func literal (or annoymous function, either way you'd like to call it) around `doSomething()` and called a go routine for that wrapper function to avoid calling goroutine on returning function.

## Race Condition Example
```
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
```
* __Counter never gets to 100, this is happening because I have multiple go routines accessing the same variable!__
* `go run -race 10_concurrency/test.go `
    * `-race` arguement will look for race conditions in your program
    * `Found 1 data race(s)`
* BTW // `Gosched` "yields" the processor, allowing other go routines to run. It does not suspend the current goroutine, so execution resumes automatically

## Mutex
* Prevent the race condition, by preventing go routines from accessing a variable when another is still using it.
```
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
```
* Locking and Unlocking code while a routine has access to it
* https://pkg.go.dev/sync#Mutex

## Atomic
* Does the same thing as Mutex, just a different approach
```
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
```