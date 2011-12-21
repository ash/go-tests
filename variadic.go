package main

import "fmt"

func main() {
    f(42)
    f(43, 45)
    f(100, 200, 300)
}

func f(x... int) {
    for n, v := range x {
        fmt.Printf("%v: %v\n", n, v)
    }
}
