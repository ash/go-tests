package main

import "fmt"

type ints    []int
type strings []string

type Sortable interface {
    Length() int
    Less(int, int) bool
    Swap(int, int)
}

func (arr ints) Length() int {
    return len(arr)
}

func (arr ints) Less(a int, b int) bool {
    return arr[a] < arr[b]
}

func (arr ints) Swap(a int, b int) {
    arr[a], arr[b] = arr[b], arr[a]
}

func (arr strings) Length() int {
    return len(arr)
}

func (arr strings) Less(a int, b int) bool {
    return arr[a] < arr[b]
}

func (arr strings) Swap(a int, b int) {
    arr[a], arr[b] = arr[b], arr[a]
}

func sort(arr Sortable) {
    for i := 0; i != arr.Length(); i++ {
	for j := i + 1; j != arr.Length(); j++ {
	    if arr.Less(j, i) {
		arr.Swap(i, j)
	    }
	}
    }
}

func main() {
    iarr := ints{10, 20, 15, 3, 10, 50, 24}
    sort(iarr)
    fmt.Printf("%v\n", iarr)

    sarr := strings{"alpha", "beta", "gamma", "delta", "epsilon"}
    sort(sarr)
    fmt.Printf("%v\n", sarr)
}
