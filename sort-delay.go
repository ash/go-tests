package main

import(
    "fmt"
    "time"
)

var get_value chan int

func send_value(x int) {
    time.Sleep(int64(x) * 1E8)
    fmt.Printf("%v\n", x)
    get_value <- x
}

func main() {
    values := []int{3, 1, 9, 7, 2, 6, 4, 8, 5, 10}
    
    get_value = make(chan int)
    
    for _, x := range values {
	go send_value(x)
    }
    for _, _ = range values {
	<- get_value
    }
}
