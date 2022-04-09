package main

import "testing"

func TestValueGet(t *testing.T) {
	value := Value("")

	valueArray := value.get()

	if len(valueArray) != 13 {
		t.Errorf("Expected suit length of 4, but got %v", len(valueArray))
	}

}

func TestValueToString(t *testing.T) {
	value := Value(King)

	expected := King

	if value.toString() != expected {
		t.Errorf("Expected suit with %q, but got %q", expected, value)
	}

}
