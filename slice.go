package main

import "fmt"

func main() {
    n := 3
    s := make([]int, n, 2*n)
    fmt.Println(s)

    fmt.Println(len(s), cap(s))

    s = s[:n+1]
    fmt.Println(len(s), cap(s))
}
