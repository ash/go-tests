package main

import(
    "fmt"
    "regexp"
)

func main() {
    str := "abCDef"
    re := regexp.MustCompile("([a-z])")

    res := re.Find(str)
    fmt.Println(res)
}

