package main

import "fmt"

func main() {
    a := []int{3, 5, 1, 75, 34}
    one_by_one(a)
}

func one_by_one(a[]int) {
    for _, v := range a {
        fmt.Printf("%v\n", v)
    }
}

