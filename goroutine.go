package main

import (
    "fmt"
    "time"
)

func gr(message string, delay int64) {
    time.Sleep(delay * 1E9)
    fmt.Printf("%s\n", message)
}

func main() {
    go gr("two", 2)
    go gr("one", 1)
    time.Sleep(5 * 1E9)
}
