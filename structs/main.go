package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// var alex person // => empty strings by default
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// alex := person{"Alex", "Anderson"}
	alex := person{lastName: "Anderson", firstName: "Alex"}
	fmt.Println(alex)
	fmt.Printf("%+v", alex) // prints field names and their values
}
