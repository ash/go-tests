package main

import "fmt"

func main() {
    x := 11
    
    var message string
    switch x {
	case -1: message = "Minus one"
	case 1 : message = "Plus one"
	default: message = "Something else"
    }
    
    fmt.Printf("%s\n", message)
}
