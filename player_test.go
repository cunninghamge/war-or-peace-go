package main

import (
	"testing"
)

func TestCardsLeft(t *testing.T) {
	deck := &Deck{testCards}
	player := Player{name: "Clarissa", deck: deck}

	for i := len(deck.cards); i > 0; i-- {
		if player.CardsLeft() != i {
			t.Errorf("got %d want %d", player.CardsLeft(), i)
		}
		player.deck.RemoveCard()
	}
}

func TestHasLost(t *testing.T) {
	t.Run("with no cards", func(t *testing.T) {
		deck := &Deck{testCards}
		player := Player{name: "Clarissa", deck: deck}

		for i := 0; i < 4; i++ {
			if player.HasLost() {
				t.Errorf("player has lost but still has cards")
			}
			player.deck.RemoveCard()
		}

		if !player.HasLost() {
			t.Errorf("player has not lost but has no cards")
		}
	})

	t.Run("with cards", func(t *testing.T) {
		deck := &Deck{testCards}
		player := Player{name: "Clarissa", deck: deck, lost: true}

		if !player.HasLost() {
			t.Errorf("player has not lost but has no cards")
		}
	})
}
