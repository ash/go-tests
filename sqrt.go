package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) (z float64) {
    z = x / 2

    z0 := z
    for {
    	z = z - (z * z - x) / (2 * z)
    	delta := math.Fabs(z - z0)
	if delta < 1e-10 {
	    break
	}
	fmt.Println(delta)
	z0 = z
    }
    
    return
}

func main() {
    fmt.Println(Sqrt(2))
}
