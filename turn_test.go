package main

import (
	"reflect"
	"testing"
)

var (
	player1   = &Player{Deck: &Deck{}}
	player2   = &Player{Deck: &Deck{}}
	emptyDeck = Deck{[]Card{}}
)

func TestTurn(t *testing.T) {
	testCases := map[string]struct {
		turnType     TurnType
		winner       *Player
		player1Cards []Card
		player2Cards []Card
	}{
		"basic": {
			turnType: basic,
			player1Cards: []Card{
				{"heart", "Jack", 11},
			},
			player2Cards: []Card{
				{"heart", "9", 9},
			},
			winner: player1,
		},
		"war": {
			turnType: war,
			player1Cards: []Card{
				{"heart", "10", 10},
				{"heart", "9", 9},
				{"diamond", "2", 2},
			},
			player2Cards: []Card{
				{"diamond", "10", 10},
				{"diamond", "8", 8},
				{"heart", "3", 3},
			},
			winner: player2,
		},
		"mutually assured destruction": {
			turnType: mutuallyAssuredDestruction,
			player1Cards: []Card{
				{"heart", "10", 10},
				{"heart", "9", 9},
				{"diamond", "2", 2},
			},
			player2Cards: []Card{
				{"diamond", "10", 10},
				{"diamond", "8", 8},
				{"heart", "2", 2},
			},
			winner: nil,
		},
	}

	for name, tc := range testCases {
		player1.Deck.Cards = tc.player1Cards
		player2.Deck.Cards = tc.player2Cards
		turn := &Turn{
			Player1: *player1,
			Player2: *player2,
		}
		var allCards []Card
		for i := 0; i < len(tc.player1Cards); i++ {
			allCards = append(allCards, tc.player1Cards[i])
			allCards = append(allCards, tc.player2Cards[i])
		}

		t.Run(name+" turn type", func(t *testing.T) {
			if turn.Type() != tc.turnType {
				t.Errorf("got %d want %d", turn.Type(), tc.turnType)
			}
		})

		t.Run(name+" winner", func(t *testing.T) {
			var (
				got  Player
				want Player
			)
			if tc.winner != nil {
				got = *turn.Winner()
				want = *tc.winner
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})

		t.Run(name+" pile cards", func(t *testing.T) {
			turn.PileCards()

			if !reflect.DeepEqual(turn.SpoilsOfWar, allCards) {
				t.Errorf("got %v want %v", turn.SpoilsOfWar, allCards)
			}
			if !reflect.DeepEqual(*player1.Deck, emptyDeck) {
				t.Errorf("got %v want %v", player1.Deck, emptyDeck)
			}
			if !reflect.DeepEqual(*player2.Deck, emptyDeck) {
				t.Errorf("got %v want %v", player2.Deck, emptyDeck)
			}
		})

		t.Run(name+" award spoils", func(t *testing.T) {
			turn.AwardSpoils(tc.winner)

			if tc.winner != nil && !reflect.DeepEqual(tc.winner.Deck.Cards, allCards) {
				t.Errorf("got %v want %v", tc.winner.Deck.Cards, allCards)
			}
			if tc.winner != player2 && !reflect.DeepEqual(*player2.Deck, emptyDeck) {
				t.Errorf("got %v want %v", player2.Deck, emptyDeck)
			}
			if tc.winner != player1 && !reflect.DeepEqual(*player1.Deck, emptyDeck) {
				t.Errorf("got %v want %v", player1.Deck, emptyDeck)
			}
		})
	}
}
