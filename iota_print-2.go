package main

import "fmt"

type Day int

const(
    Monday Day = iota
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday
)

var dayname = []string{
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
    "Sunday",
}

func (d Day) String() string {
    return dayname[d]
}

func main() {
    fmt.Println(Wednesday)
}
