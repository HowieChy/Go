package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("预期长度是52，但现在长度是 %v", len(d))
	}
}
func TestSvaeToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesing")
	deck := newDeck()
	deck.saveToFile("_decktesing")

	loadedDeck := newDeckFromFile("_decktesing")
	if len(loadedDeck) != 520 {
		t.Errorf("预期长度是520，但现在长度是 %v", len(loadedDeck))
	}
	os.Remove("_decktesing")
}
