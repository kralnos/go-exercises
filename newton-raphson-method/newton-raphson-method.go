// https://tour.golang.org/flowcontrol/8
package main

import (
	"fmt"
	"math"
)

// Sqrt computes the square root of x2 using Newton Raphson method
// Reference - https://en.wikipedia.org/wiki/Newton%27s_method
func Sqrt(x2 float64) float64 {
	// function
	f := func(x float64) float64 {
		return x*x - x2
	}

	// derivative
	fd := func(x float64) float64 {
		return x * 2
	}

	newtonRaphson := func(x float64) float64 {
		return x - f(x)/fd(x)
	}

	xn := newtonRaphson(1.0)
	i := 0
	for xn1 := newtonRaphson(xn); math.Abs(xn-xn1) > 0.000001; {
		if i == 10 {
			fmt.Println("Loop has reached its limit and will no longer evaluate.")
			break
		}
		xn = xn1
		xn1 = newtonRaphson(xn)
		i++
	}

	return xn
}

func main() {
	x := 3.0
	fmt.Printf("Sqrt of %f using math.Sqrt: %f\n", x, math.Sqrt(x))
	fmt.Printf("Sqrt of %f using Newton Raphson: %f\n", x, Sqrt(x))
}
