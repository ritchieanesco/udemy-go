package main

import "testing"

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
