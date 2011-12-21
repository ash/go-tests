package main

import "fmt"

func main() {
    d(42)
}

func deferred(x int) {
    fmt.Printf("Value = %v\n", x);
}

func d(x int) {
    defer deferred(2);
    defer deferred(3);

    return
}
