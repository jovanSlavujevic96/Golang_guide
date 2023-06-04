package main

func main() {
	cards := newDeck()

	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	// to capture both of the values
	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()
}
