package main

import "fmt"

func main() {
    c, d := f(7, 2)
    fmt.Printf("%v, %v\n", c, d)
    
    c, d = f(2, 7)
    fmt.Printf("%v, %v\n", c, d)
}

func f(a, b int) (c, d int) {
    if (a < b) {
	c = a
	d = b
    } else {
	c = b
	d = a
    }
    
    return
}
