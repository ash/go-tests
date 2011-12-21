package main

import "fmt"

func main() {
    for c := 0; c != 10; c++ {
	fmt.Printf("%d! = %d\n", c, factorial(c))
    }
}

func factorial(n int) (int) {
    if (n < 2) {
        return 1
    }
    
    return n * factorial(n - 1)
}

