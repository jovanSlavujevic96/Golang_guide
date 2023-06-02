package main

import "fmt"

func main() {
	// initialization statement
	card := newCard()

	fmt.Println(card)
}

// when we want to return some value within the function we mention it right after the paranethesesis
func newCard() string {
	return "Five of Diamonds"
}
