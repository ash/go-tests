package main

import "fmt"

func main() {
    c := 0
    Loop:
        fmt.Printf("%d\n", c)
        c++
        if (c != 10) {
            goto Loop
        }
}
