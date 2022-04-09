package main

import (
	"os"
	"testing"
)

const (
	sizeDeck        = 52
	fileNameTesting = "_decktesting"
)

func TestNewDeck(t *testing.T) {
	deck := Deck{}

	deck = deck.newDeck()

	length := len(deck.Cards)

	if length != sizeDeck {
		t.Errorf("Expected deck length of %v, but got %v", sizeDeck, length)
	}

	firtCard := deck.Cards[0]

	if firtCard.Value != Ace || firtCard.Suit != Spades {
		expected := Card{Value: Ace, Suit: Spades}
		t.Errorf("Expected first card of %q, but got %q", expected.toString(), firtCard.toString())
	}

	lastCard := deck.Cards[length-1]

	if lastCard.Value != King || lastCard.Suit != Clubs {
		expected := Card{Value: King, Suit: Clubs}
		t.Errorf("Expected last card of %q, but got %q", expected.toString(), lastCard.toString())
	}
}

func TestClearDeckCards(t *testing.T) {
	deck := Deck{}
	deck = deck.newDeck()

	deck = deck.clear()

	length := len(deck.Cards)

	expected := 0

	if length != 0 {
		t.Errorf("Expected deck length of %v, but got %v", expected, length)
	}
}

func TestSaveDeckFile(t *testing.T) {

	deck := Deck{}
	deck = deck.newDeck()

	length := len(deck.Cards)

	if length != sizeDeck {
		t.Errorf("Expected %v cards in deck, but got %v", sizeDeck, length)
	}

	err := deck.saveToFile(fileNameTesting)

	if err != nil {
		t.Errorf("Expect save deck to file, but got %q", err)
	}

	err = os.Remove(fileNameTesting)

	if err != nil {
		t.Errorf("Expect remove %q file, but %q", fileNameTesting, err)
	}

}

func TestNewDeckFromFile(t *testing.T) {

	deck := Deck{}
	deck = deck.newDeck()
	err := deck.saveToFile(fileNameTesting)

	deck = deck.clear()

	loadedDeck := deck.newDeckFromFile(fileNameTesting)

	length := len(loadedDeck.Cards)

	if length != sizeDeck {
		t.Errorf("Expected %v cards in deck, but got %v", sizeDeck, length)
	}

	err = os.Remove(fileNameTesting)

	if err != nil {
		t.Errorf("Expect remove %q file, but %q", fileNameTesting, err)
	}

}
