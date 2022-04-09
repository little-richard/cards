package main

import (
	"log"
)

func main() {

	deck := Deck{}

	deck, err := deck.newDeckFromFile("my_deck.txt")

	if err != nil {
		log.Fatal("Error:", err)
	}

	deck.print()

}
