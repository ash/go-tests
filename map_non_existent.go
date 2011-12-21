package main

import "fmt"

func main() {
    m := map[string]int{"alpha": 1, "beta": 2}
    fmt.Println(m)
    
    n := m["non existent"]
    fmt.Println(n)
    fmt.Println(m)
}
