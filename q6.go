package main

import "fmt"

func main() {
    a := []float64{1,2,3,4,5,6,7,8,9,10,11,12}
    fmt.Printf("%g\n", avg(a))
    fmt.Printf("%g\n", avg(a[2:5]))
}

func avg(a[]float64) (float64) {
    s := 0.0
    for _, x := range a {
	s += x
    }
    return s / float64(len(a))
}
