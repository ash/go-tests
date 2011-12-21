package main

import (
    "fmt"
    "time"
)

var c chan int

func message(msg string, delay int) {
    time.Sleep(int64(delay) * 1E9)
    fmt.Printf("%s\n", msg)
    c <- delay
}

func main() {
    c = make(chan int)
    go message("two", 2)
    go message("one", 1)
    fmt.Printf("sent\n")
    a, b := <- c, <- c
    fmt.Printf("done: %v, %v\n", a, b)
}
