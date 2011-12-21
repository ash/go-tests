package main

import "fmt"

func main() {
    fib(20)
}

func fib(n int) {
    a, b := 1, 1
    fmt.Printf("%v\n%v\n", a, b)
    
    for c := 2; c != n; c++ {
       c := a + b
       a, b = b, c
       fmt.Printf("%v\n", c)
    }
}
