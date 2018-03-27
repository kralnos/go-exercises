package main

import (
	"fmt"
	"math"
)

const DIFF = 0.000001

// Sqrt computes the square root of x using Newton's method
func Sqrt(x float64) float64 {
	fmt.Printf("sqrt of %f is %f\n", x, math.Sqrt(x))
    newtonsMethod := func(z, x float64) float64 {
        return z - (z * z - x) / (2 * z)        
    }
    z := 1.0
	for nm := newtonsMethod(1.0, x); math.Abs(nm - z) > DIFF; {        
		nm = newtonsMethod(nm, x)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
