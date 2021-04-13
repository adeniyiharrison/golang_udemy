package pkg

import (
	"testing"
)

func TestAverage(t *testing.T) {
	x := Average(10, 5)
	if x != 7.5 {
		t.Errorf("Did not finish with 7.5: x = %v", x)
	}
}

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

func BenchmarkAveage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Average(1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}
