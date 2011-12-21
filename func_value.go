package main

import "fmt"

func main() {
    myfunc := func(x int) int {
        return 2 * x
    }
    fmt.Println(myfunc(21))
}
