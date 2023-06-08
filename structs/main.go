package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo // use one struct within another one
	contactInfo // actually we have field named & type of contactInfo
}

// receiver func
func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
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
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("")
	jim.updateName("jimmy") // this is not going to update the value
	jim.print()
}
