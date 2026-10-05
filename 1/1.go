// for и if
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	for i := 0; i < 100; i++ {
		prev := z
		z -= (z*z - x) / (2*z)
		if math.Abs(z-prev) < 1e-15*math.Abs(z) {
			break
		}
		
	} 
	return z
} 

func main() {
	fmt.Println(Sqrt(4))
}