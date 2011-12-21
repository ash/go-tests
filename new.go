package main

import "fmt"

func main() {
    p := new([]int)
    fmt.Printf("%t\n", p)
    fmt.Printf("%v\n", *p)
}
