package main

import "fmt"

func main() {
    d(42)
}

func d(x int) {
    defer func(y int) {
        fmt.Printf("Before exit x=%v, y=%v\n", x, y)
    }(x)

    x = 3
    
    fmt.Printf("x=%v\n", x)
    
    return
}
