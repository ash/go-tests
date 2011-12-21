package main 

import "utf8"

func main() {
    s := "абв"
    println(len(s))
    println(utf8.RuneCount([]byte(s)))
}
