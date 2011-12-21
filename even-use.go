package main

import "./even"
import "fmt"

func main() {
    fmt.Printf("6? %v\n", even.Even(6))
    fmt.Printf("7? %v\n", even.Even(7))
}
