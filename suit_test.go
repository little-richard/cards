package main

import "testing"

func TestSuitGet(t *testing.T) {
	suit := Suit("")

	suitArray := suit.get()

	if len(suitArray) != 4 {
		t.Errorf("Expected suit length of 4, but got %v", len(suitArray))
	}

}

func TestSuitToString(t *testing.T) {
	suit := Suit(Clubs)

	expected := Clubs

	if suit.toString() != expected {
		t.Errorf("Expected suit with %q, but got %q", expected, suit)
	}

}
