package main

import "fmt"

func main() {
    var p *int
    fmt.Printf("%v\n", p)

    var i int = 42
    p = &i
    fmt.Printf("%v\n", p)
    
    var value int = *p
    fmt.Printf("%v\n", value)
    
    y := *p
    fmt.Printf("%v\n", y)
}
