package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	newDeck := NewDeck()

	if len(newDeck) != 52 {
		t.Errorf("got %d not 52 for deck length", len(newDeck))
	}

	var hearts []Card
	var diamonds []Card
	var spades []Card
	var clubs []Card
	for _, card := range newDeck {
		switch card.Suit {
		case "heart":
			hearts = append(hearts, card)
		case "diamond":
			diamonds = append(diamonds, card)
		case "spade":
			spades = append(spades, card)
		case "club":
			clubs = append(clubs, card)
		}
	}

	for _, suit := range [][]Card{hearts, diamonds, spades, clubs} {
		if len(suit) != 13 {
			t.Errorf("got %d not 13 for cards in suit", len(suit))
		}
	}
}

func TestNewPlayers(t *testing.T) {
	player1, player2 := NewPlayers()

	if player1.Name != "Cacco" {
		t.Errorf("got %q want %q for player1 name", player1.Name, "Cacco")
	}

	if player2.Name != "Pickles" {
		t.Errorf("got %q want %q for player2 name", player2.Name, "Pickles")
	}

	if len(player1.Deck.Cards) != 26 {
		t.Errorf("got %d not 26 for player1's cards", len(player1.Deck.Cards))
	}

	if len(player2.Deck.Cards) != 26 {
		t.Errorf("got %d not 26 for player2's cards", len(player2.Deck.Cards))
	}
}
