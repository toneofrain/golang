package main

import (
	"fmt"
	"math"
)

func eqaul(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f + %0.18f = %v\n", c, a+b, eqaul(a+b, c))
}
