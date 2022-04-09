package main

type Suit string

const (
	Spades   Suit = "Spades"
	Diamonds      = "Diamonds"
	Hearts        = "Hearts"
	Clubs         = "Clubs"
	_             = ""
)

func (s Suit) get() []Suit {
	return []Suit{Spades, Diamonds, Hearts, Clubs}
}

func (s Suit) toString() string {
	return string(s)
}
