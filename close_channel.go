package main

import "fmt"
import "time"

func sendthem(ch chan int, value int) {
    ch <- value
}

func main() {
    ch := make(chan int)
    
    go sendthem(ch, 5)
    go sendthem(ch, 7)
    go sendthem(ch, 9)

time.Sleep(2*1e9);
    for {
        value, ok := <- ch
        if !ok {
    	    break
        }
        fmt.Println(value)
    }
    
}
