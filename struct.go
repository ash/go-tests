package main

import "fmt"

func main() {
    type person struct {
	name string
	age int
    }
    
    john := person{"John", 30}
    fmt.Println(john)
    
    var peter person
    peter.name = "Peter"
    peter.age = 20
    fmt.Println(peter)
    
    var matt *person = new(person)
    (*matt).name = "Matt"
    matt.age = 25
    fmt.Println(matt)
}
