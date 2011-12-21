package main

import "fmt"

func main() {
    m := make(map[string]int)
    fmt.Println(m)
    
    m["one"] = 1
    fmt.Println(m)
}
