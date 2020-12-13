package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	length := 16
	firstCard := "Ace of Spades"
	lastCard := "Four of Clubs"

	if len(d) != length {
		t.Errorf("Error Test: Expected length of %v, but got %v", length, len(d))
	}

	if d[0] != firstCard {
		t.Errorf("Error Test: Expected first card of %v, but got %v", firstCard, d[0])
	}
	if d[length-1] != lastCard {
		t.Errorf("Error Test: Expected last card of %v, but got %v", lastCard, d[length-1])
	}
}

func TestSaveToFilekAndNewDeckFromFile(t *testing.T) {
	length := 16
	testFileName := "_decktesting.txt"
	os.Remove(testFileName)
	deck := newDeck()
	deck.saveToFile(testFileName)

	loadedDeck := newDeckFromFile(testFileName)

	if len(loadedDeck) != length {
		t.Errorf("Error Test: Expected length of loaded should be %v, but got %v", length, len(deck))
	}

	os.Remove(testFileName)
}
