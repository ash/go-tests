package main

import "fmt"

func main() {
   fmt.Println(1)
   panicker()
   fmt.Println(2)
}

func panicker() {
    defer func(){
	fmt.Println("Defer")
	recover()
    }()
    
    fmt.Println(3)
    panic("Hi!")
    fmt.Println(4)
}
