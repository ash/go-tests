package main

import "fmt"

type person struct {
    name string
    age int
}

func main() {
    someone := person{"Anders", 33}
    fmt.Printf("%v\n", someone)
    
    type pair [2]int
    p := new(pair)
    p[0] = 3
    p[1] = 4
    fmt.Printf("%v\n", p)
    
    var t pair
    t[0] = 7
    t[1] = 8
    fmt.Printf("%v\n", t)
}
