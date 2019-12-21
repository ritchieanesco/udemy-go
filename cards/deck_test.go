package main

import "testing"

import "os"

// t is the test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()
	l := len(d)
	if l != 16 {
		t.Errorf("Expected deck length of 16, but got %v", l)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[l-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[l-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	l := len(loadedDeck)
	if l != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", l)
	}
	os.Remove("_decktesting")
}
