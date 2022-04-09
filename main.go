package main

import "fmt"

func main() {

	deck := Deck{}

	deck = deck.newDeck()

	hand, remainingDeck := deal(deck, 5)

	fmt.Println("Hand:")
	hand.print()
	fmt.Println("\n\n")

	fmt.Println("Remaining deck:")
	remainingDeck.print()

}
