package main

func main() {
    var a [5]int
    var b [5]int
    a[3] = 42
    a[4] = 43

    b = a
    b[3] = 100
    println(a[3])
    println(b[3])

    a[3] = 43
    println(a[3])
    println(b[3])
}

