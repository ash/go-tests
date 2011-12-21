package main

func main() {
    for c := 1; c != 100; c++ {
	p := false
	if (0 == c % 5) {
	    print("Fizz")
	    p = true
	}
	if (0 == c % 3) {
	    print("Buzz")
	    p = true
	}
	if (!p) {
	    print(c)
	}
	println()
    }
}
