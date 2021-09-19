package main

import (
	"testing"
)

func TestHasLost(t *testing.T) {
	deck := &Deck{testCards}
	player := Player{Name: "Clarissa", Deck: deck}

	for i := 0; i < 3; i++ {
		if player.HasLost() {
			t.Errorf("player has lost but still has cards")
		}
		player.Deck.RemoveCard()
	}

	if !player.HasLost() {
		t.Errorf("player has not lost but has no cards")
	}
}
