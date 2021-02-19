package main

import "testing"

// TEST function are all first character UPPER case NOT standard camelCase
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got " + len(d))
	}
}
