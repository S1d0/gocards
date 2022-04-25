package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected lenght of deck of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got: %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got: %v", d[0])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	d := newDeck()
	d.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)

	loadedDeck.print()
	if len(loadedDeck) != len(d) {
		t.Errorf("Expected %v cards in deck, got %v cards", len(d), len(loadedDeck))
	}

	t.Cleanup(func() {
		os.Remove(filename)
	})
}
