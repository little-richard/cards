package main

type Card struct {
	Suit  Suit
	Value Value
}

func (d Card) toString() string {
	return d.Value.toString() + " of " + d.Suit.toString()
}
