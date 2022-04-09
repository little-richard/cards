package main

import "log"

func main() {

	deck := Deck{}

	deck = deck.newDeck()

	//hand, remainingDeck := deck.deal(deck, 5)
	//
	//fmt.Println("Hand:")
	//hand.print()
	//fmt.Println("\n\n")
	//
	//fmt.Println("Remaining deck:")
	//remainingDeck.print()
	//
	//test := "Hi there!"
	//
	//fmt.Println([]byte(test))

	err := deck.saveToFile("my_deck.txt")

	if err != nil {
		log.Fatal("Error: ", err)
	}

}
