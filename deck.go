package main

import "fmt"

type Deck struct {
	Cards []Card
}

func (d Deck) newDeck() Deck {
	deck := Deck{}
	values := Value("")
	suits := Suit("")

	for _, suit := range suits.get() {
		for _, value := range values.get() {
			deck = deck.addCard(Card{Suit: suit, Value: value})
		}
	}

	return deck
}

func (d Deck) addCard(newCard Card) Deck {
	d.Cards = append(d.Cards, newCard)
	return d
}

func (d Deck) print() {

	fmt.Println("Quantidade Cartas Deck: ", len(d.Cards))

	for _, card := range d.Cards {
		fmt.Printf("%s of %s\n", card.Value, card.Suit)
	}
}
