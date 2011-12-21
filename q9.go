package main

import "fmt"

var stack [10]int
var top int = 0

func main() {
    push(2)
    push(4)
    fmt.Printf("%v", pop())
    push(5)
    fmt.Printf("%v", pop())
    fmt.Printf("%v", pop())
}

func push(v int) {
    stack[top] = v
    top++
}

func pop() (int) {
    top--
    return stack[top]
}
