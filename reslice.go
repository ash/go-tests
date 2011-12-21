package main

import "fmt"

func main() {
    slice := make([]int, 5, 10)
    fmt.Println(len(slice), cap(slice))

    slice = slice[0:7]
    fmt.Println(len(slice), cap(slice))

//    slice = slice[0:12]
//    fmt.Println(len(slice), cap(slice))
}
