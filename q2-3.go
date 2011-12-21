package main

import "fmt"

func main() {
    var a[10]int
    for c := 0; c != 10; c++ {
	a[c] = c
    }
    for c := 0; c != 10; c++ {
        fmt.Printf("%d\n", a[c])
    }
    fmt.Printf("%v", a)
}
