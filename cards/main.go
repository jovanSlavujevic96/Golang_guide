package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

	cardsA := newDeckFromFile("my_cards") // no error (file exists)
	fmt.Println(cardsA.toString())

	cardsB := newDeckFromFile("my_") // error occures
	fmt.Println(cardsB.toString())

	// to capture both of the values
	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()
}
