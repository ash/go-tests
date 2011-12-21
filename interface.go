package main

import "fmt"

type s struct {
    i int
}

func (x *s) Get() int {
    return x.i
}

func (x *s) Set(i int) {
    x.i = i
}

type I interface {
    Get() int
    Set(int)
}

func assign_and_print(x I, v int) {
    x.Set(v)
    fmt.Printf("%v\n", x.Get())
}

func main() {
    var x s
    x.Set(15)
    fmt.Printf("%v\n", x.Get())
    
    var y s
    assign_and_print(&y, 42)
}
