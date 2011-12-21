package main

import "fmt"

func main() {
    x := int64(1E9)
    var y int = 2
    var z int64 = x * y
    fmt.Printf("%v %t\n", z, z)
}
