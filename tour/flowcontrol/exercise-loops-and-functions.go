//go:build OMIT

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	for z*z > x {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	tests := []float64{
		-1,
		0,
		1,
		2,
		4,
		10,
		144,
		10000,
		math.MaxFloat64,
	}

	for _, x := range tests {
		fmt.Printf("MySqrt = %g, math.Sqrt = %g\n", Sqrt(x), math.Sqrt(x))
	}

}
