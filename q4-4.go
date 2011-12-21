package main

func main() {
    a := []float64{1,2,3,4,5,6,7,8,9,10,11,12}
    
    b := a[1:7]
    s := 0.0
    for _, x := range b {
	s += x
    }
    println(s / float64(len(b)))
}
