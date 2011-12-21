package main

import "fmt"

type Day int

var daynames = []string{
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
    "Sunday",
}

func (this Day) String() string {
    return daynames[this - 1]
}

func main() {
    day := Day(3)
    fmt.Println(day.String())
}
