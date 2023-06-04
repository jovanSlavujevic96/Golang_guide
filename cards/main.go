package main

import "fmt"

func main() {
	// string variable
	card := newCard()
	fmt.Println(card)

	// slice of strings variable
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

// when we want to return some value within the function we mention it right after the paranethesesis
func newCard() string {
	return "Five of Diamonds"
}
