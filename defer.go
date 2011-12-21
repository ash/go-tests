package main

import "fmt"

func main() {
    d()
}

func d() {
    defer fmt.Println("Before exit")
    defer fmt.Println("Before before exit ;-)")
    
    fmt.Printf("a\n")
    
    return
}
