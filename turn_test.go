package main

import (
	"reflect"
	"testing"
)

func TestTurn(t *testing.T) {
	card1 := Card{"heart", "Jack", 11}
	card2 := Card{"heart", "10", 10}
	card3 := Card{"heart", "9", 9}
	card4 := Card{"diamond", "Jack", 11}
	card5 := Card{"heart", "8", 8}
	card6 := Card{"diamond", "Queen", 12}
	card7 := Card{"heart", "3", 3}
	card8 := Card{"diamond", "2", 2}

	deck1 := &Deck{
		[]Card{card1, card2, card5, card8},
	}
	deck2 := &Deck{
		[]Card{card3, card4, card6, card7},
	}

	player1 := Player{"Megan", deck1}
	player2 := Player{"Aurora", deck2}

	t.Run("attributes", func(t *testing.T) {
		turn := Turn{player1, player2, []Card{}}

		if turn.Player1 != player1 {
			t.Errorf("got %v want %v for player1", turn.Player1, player1)
		}

		if turn.Player2 != player2 {
			t.Errorf("got %v want %v for player2", turn.Player2, player2)
		}

		if !reflect.DeepEqual(turn.SpoilsOfWar, []Card{}) {
			t.Errorf("got %v want %v for spoils of war", turn.SpoilsOfWar, []Card{})
		}
	})
	// player1, player2, spoilsOfWar
	// type: basic, war, or mutuallyAssuredDestruction
	// winner
	// pile cards
	// award spoils
}
