package main

import "fmt"

func main() {
	// string variable
	card := newCard()
	fmt.Println(card)

	// slice of strings variable
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	// iterate over slice & print one per one
	for i, card := range cards {
		fmt.Println(i, card)
	}

	// print the whole slice
	fmt.Println(cards)
}

// when we want to return some value within the function we mention it right after the paranethesesis
func newCard() string {
	return "Five of Diamonds"
}
