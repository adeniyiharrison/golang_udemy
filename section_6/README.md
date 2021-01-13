# Section 6 Notes

### How computers work
* Computer storage
    * 1 bit (store on or off state)
    * 8 bits = byte
    * 1000 bytes = kb
    * 1000 kb = mb
    * 1000 mb = gb
    * 1000 gb = tb

* Vacuum Tubes (ran hot and attracted moths) -> Transistors (smaller and cool) -> Silicon Wafers (integrated circuits)
* Moore's Law basically states that the number of bits (I/O state) stored on a chip will double every two years
    * 4000 back in 1971
    * 2.6 Billion now

### Strings
* Strings ultimely are slices of bytes
    * https://www.ascii-code.com
```
func main () {
    s := fmt.Sprintln("Hello, world")
    fmt.Printf("%T\n", s)

    // Convert to s to slice of bytes
    bs := []byte(s)
    
    fmt.Println(bs)
    // [72 101 108 108 111 44 32 119 111 114 108 100 10]
    
    fmt.Printf("%T\n", bs)
    // []uint8
}
```

### Constants
* `const` declares a constant value
```
// without specifying type
const a = 42
const b = 42.78
const c = "James Bond"

// OR

const (
    a int = 42
    b float64 = 42.78
    c string = "James Bond"
)
```

### Iota
* Auto incrementing constant

```
const (
    a = iota
    b
    c
)
// this will return 1,2,3

const (
    d = iota
    e = iota
    f = iota
    g
)
// this will return 1,2,3,4.
// Resets on declaration of new const
```