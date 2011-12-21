package main

import "fmt"

type person struct {
    name string
    age int
}

func (who *person) inc() {
    who.age++
}

func main() {
    p := person{name: "Someone", age: 100}
    fmt.Printf("%v\n", p)
    
    p.inc()
    fmt.Printf("%v\n", p)
}
