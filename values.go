package main

type Value string

const (
	Ace   Value = "Ace"
	Two         = "Two"
	Three       = "Three"
	Four        = "Four"
	Five        = "Five"
	Six         = "Six"
	Seven       = "Seven"
	Eight       = "Eight"
	Nine        = "Nine"
	Ten         = "Ten"
	Jack        = "Jack"
	Queen       = "Queen"
	King        = "King"
	_           = ""
)

func (v Value) get() []Value {
	return []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
}
