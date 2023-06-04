package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// make function for "receiver" which is deck
func (d deck) print() {
	// iterate over slice & print one per one
	for i, card := range d {
		fmt.Println(i, card)
	}
}
