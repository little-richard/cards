package main

import "fmt"

func main() {

	deck := Deck{}

	deck = deck.newDeck()

	deck.print()

	fmt.Println("\n")

	deck.shuffle()

	deck.print()
}
