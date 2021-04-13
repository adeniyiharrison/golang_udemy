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
