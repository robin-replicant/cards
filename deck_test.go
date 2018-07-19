package main

import (
	"os"
	"testing"
)

// you could test for example:
// check the length of the deck
// make sure the first card is the ace of spades

func TestNewDeck(t *testing.T) { //test handler
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "A ♠" {
		t.Errorf("Expected first card A ♠, but got %v", d[0])
	}

	if d[len(d)-1] != "K ♣" {
		t.Errorf("Expected first card K ♣, but got %v", len(d)-1)
	}
}

// test create deck, save deck and load it
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
