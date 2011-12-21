package main

import "fmt"

func main() {
    s := "абв"
    
    for k, v := range s {
	fmt.Println(k, v, string(v))
    }
}
