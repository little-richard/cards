package main

func main() {

	deck := Deck{}

	deck = deck.newDeckFromFile("my_deckss.txt")

	deck.print()

}
