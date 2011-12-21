package main

import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5, -1, -2, -3, -4, -5}
    
    b := mapf(a, f_sqr)
    fmt.Printf("%v\n", b)
    
    b = mapf(a, f_next)
    fmt.Printf("%v\n", b)
    
    b = mapf(a, f_mul2)
    fmt.Printf("%v\n", b)
}

func f_sqr(i int) (int) {
    return i * i
}

func f_next(i int) (int) {
    return i + 1
}

func f_mul2(i int) (int) {
    return i * 2
}

func mapf(a[]int, f func(int) int) (m[]int) {
    m = make([]int, len(a))
    for i, v := range a {
	m[i] = f(v)
    }
    return
}
