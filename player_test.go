package main

import (
	"reflect"
	"testing"
)

func TestPlayer(t *testing.T) {
	card1 := Card{"diamond", "Queen", 12}
	card2 := Card{"spade", "3", 3}
	card3 := Card{"heart", "Ace", 14}
	deck := &Deck{
		[]Card{card1, card2, card3},
	}
	player := Player{"Clarissa", deck}

	t.Run("attributes", func(t *testing.T) {
		if player.Name != "Clarissa" {
			t.Errorf("got %q want %q for player name", player.Name, "Clarissa")
		}

		if !reflect.DeepEqual(player.Deck, deck) {
			t.Errorf("got %v, want %v for player deck", player.Deck, deck)
		}
	})

	t.Run("player loses when they run out of cards", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			if player.HasLost() {
				t.Errorf("player has lost but still has cards")
			}
			player.Deck.RemoveCard()
		}

		if !player.HasLost() {
			t.Errorf("player has not lost but has no cards")
		}
	})
}
