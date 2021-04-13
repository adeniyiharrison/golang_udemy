# Testing & Benchmarking

## Intro
* Tests must be
    * in a file that ends with `_test.go`
    * put in the file in the __same package__ as the one being tested
    * be in a func with a signature `func TestXXX(*testing.T)`
* Run a test
    * `go test`
    * `go test -v` // more information "verbose"
* Deal with test failure
    * use the Error, Fail or related methods to signnal failure

```
package pkg

// Average calculates average
func Average(v ...float64) float64 {
	var x float64
	for _, i := range v {
		x += i
	}
	x = x / float64(len(v))

	return x
}
```
```
package pkg

import (
	"testing"
)

func Test1(t *testing.T) {
	values := []float64{5, 10}
	x := Average(values[0], values[1])
	if x != 7.5 {
		t.Errorf("Did not finish with 7.5: x = %v", x)
	}
}

```

## Table tests
```
func TestAverage2(t *testing.T) {
	type tableTest struct {
		Data   []float64
		Answer float64
	}

	tests := []tableTest{
		{
			Data:   []float64{10, 20},
			Answer: float64(15)},
		{
			Data:   []float64{5, 5, 5, 5, 5},
			Answer: float64(5)},
		{
			Data:   []float64{1, 2, 3, 4},
			Answer: float64(2.5)},
	}

	for _, test := range tests {
		if x := Average(test.Data...); x != test.Answer {
			t.Errorf("Average should equal %v but got %v", test.Answer, x)
		}
	}
}
```

## Benchmark
* Measure performance of your code
* `go test -bench .`
* `go test -bench Average`

```
func BenchmarkAveage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Average(1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}
```
```
goos: darwin
goarch: amd64
pkg: github.com/adeniyiharrison/golang_udemy/14_testing
BenchmarkAveage-8   	98969569	        11.9 ns/op
PASS
ok  	github.com/adeniyiharrison/golang_udemy/14_testing	1.665s
```

## Coverage
* `go test -cover`