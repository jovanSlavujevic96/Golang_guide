package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // use one struct within another one
}

func main() {
	// var alex person // => empty strings by default
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// alex := person{"Alex", "Anderson"}
	alex := person{lastName: "Anderson", firstName: "Alex"}
	fmt.Println(alex)
	fmt.Printf("%+v", alex) // prints field names and their values

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Printf("%+v", jim)
}
