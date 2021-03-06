# Ninja Level 10

1. Get the following code to function
```
    func p1(){
        c := make(chan int)
        c <- 42
        fmt.Println(<-c)
    }
```

2. Get the following code to function
```
    func p2(){
        cs := make(chan<- int)
        go func(){
            cs<- 1
        }()
        fmt.Println(<-cs)
        ftm.Printf("%T\n", cs)
    }
```

3. Fix this
```
    func p3() {
        c := make(chan int)

        go func() {
            for i := 0; i < 10; i++ {
                c <- i
            }
        }()

        for {
            fmt.Println(<-c)
        }
    }
```

4. Pull values down using select statement
```
    func p4(){
        q := make(chan int)
        c := gen(q)
        recieve(c, q)
        fmt.Println("exit")
    }

    func gen(c <-chan int) <-chan int{
        for i := 0; i < 10; i++ {
            c <- i
        }

        return c
    }
```
5. Write a program that writes 100 numbers to a channel then print them
6. Write a program that launches 10 goroutines that add 10 numbers to a channel
    * Pull off from the channel and print them