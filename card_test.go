package main

import "testing"

func TestNewCard(t *testing.T) {
	card := Card{Value: King, Suit: Spades}

	if card.Value != King || card.Suit != Spades {
		t.Errorf("Expected card with %q of %q but got %q", King, Spades, card.toString())
	}

}

func TestCardToString(t *testing.T) {
	card := Card{Value: Ace, Suit: Spades}

	if card.toString() != "Ace of Spades" {
		t.Errorf("Expected Ace of Spaces but got %q", card.toString())
	}
}
