package main

import(
    "fmt"
    "math"
)

type Number float64

type Complex struct {
    Real float64
    Imag float64
}

type Vector interface {
    Length() float64
}

func (n Number) Length() float64 {
    if (n < 0) {
	return -float64(n)
    }
    return float64(n)
}

func (c Complex) Length() float64 {
    return math.Sqrt(c.Real * c.Real + c.Imag * c.Imag)
}

func main() {
    var x Vector
    
    x = Number(42)
    fmt.Println(x.Length())
    
    x = Complex{3, 4}
    fmt.Println(x.Length())
}
