package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	sqrtOrig := math.Sqrt(2)
	sqrt := Sqrt(2)
	fmt.Println("math.Sqrt: ", sqrtOrig)
	fmt.Println("exercise function: ", sqrt)
	fmt.Println("diff: ", sqrtOrig-sqrt)
}
