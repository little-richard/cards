package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Deck struct {
	Cards []Card
}

func (d Deck) newDeck() Deck {
	values := Value("")
	suits := Suit("")

	for _, suit := range suits.get() {
		for _, value := range values.get() {
			d = d.addCard(Card{Suit: suit, Value: value})
		}
	}

	return d
}

func (d Deck) addCard(newCard Card) Deck {
	d.Cards = append(d.Cards, newCard)
	return d
}

func (d Deck) print() {

	for _, card := range d.Cards {
		fmt.Printf("%s of %s\n", card.Value, card.Suit)
	}

}

func (d Deck) deal(deck Deck, handSize int) (Deck, Deck) {
	return Deck{deck.Cards[:handSize]}, Deck{deck.Cards[handSize:]}
}

func (d Deck) toString() string {
	var cardsStr []string

	for _, card := range d.Cards {
		cardsStr = append(cardsStr, card.toString())
	}

	return strings.Join(cardsStr, ",")
}

func (d Deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d Deck) newDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal("Error:", err)
	}

	fileStrData := string(bs)

	cardsStrSplit := strings.Split(fileStrData, ",")

	for _, cardStr := range cardsStrSplit {
		cardSplit := strings.Split(cardStr, " of ")
		value := Value(cardSplit[0])
		suit := Suit(cardSplit[1])
		d = d.addCard(Card{Value: value, Suit: suit})
	}

	return d

}
