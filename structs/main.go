package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{"Alex", "Anderson"}
	alex := person{lastName: "Anderson", firstName: "Alex"}
	fmt.Println(alex)
}
