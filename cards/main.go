package main

func main() {
	cards := newDeck()

	// to capture both of the values
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}
