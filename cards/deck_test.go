package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck of length of 20 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Spades, but got %v", d[0])

	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Clubs, but got %v", d[len(d)-1])

	}
}
