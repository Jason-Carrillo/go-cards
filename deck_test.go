package main

import (
	"os"
	"testing"
)

// TEST function are all first character UPPER case NOT standard camelCase
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got ", d[0])
	}

	if d[d.len-1] != "Four of Cpades" {
		t.Errorf("Expected last card to be Four of Clubs but got ", d[d.len-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in a deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
