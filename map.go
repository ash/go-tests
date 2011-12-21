package main

func main() {
    greek := map[string]string {
        "alpha": "a",
        "beta" : "b",
        "gamma": "g",
    }
    println(greek["alpha"])
    
    for k, v := range greek {
	println(k, v)
    }
    
    chi, exists := greek["chi"]
    println(chi, exists)
}
