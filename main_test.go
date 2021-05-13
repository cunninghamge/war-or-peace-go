package main

import (
	"reflect"
	"testing"
)

func TestCard(t *testing.T) {
	card := Card{
		"heart",
		"Jack",
		11,
	}

	if card.Suit != "heart" {
		t.Errorf("got %s want %s for card suit", card.Suit, "heart")
	}

	if card.Value != "Jack" {
		t.Errorf("got %s want %s for card value", card.Value, "Jack")
	}

	if card.Rank != 11 {
		t.Errorf("got %d want %d for card rank", card.Rank, 11)
	}
}

func TestDeckCards(t *testing.T) {
	card1 := Card{"diamond", "Queen", 12}
	card2 := Card{"spade", "3", 3}
	card3 := Card{"heart", "Ace", 14}
	deck := Deck{
		[]Card{card1, card2, card3},
	}

	t.Run("deck.Cards", func(t *testing.T) {
		want := []Card{card1, card2, card3}
		if !reflect.DeepEqual(deck.Cards, want) {
			t.Errorf("got %v want %v for deck cards", deck.Cards, want)
		}
	})

	t.Run("deck.RankofCardAt", func(t *testing.T) {
		if deck.RankofCardAt(0) != card1 {
			t.Errorf("got %v want %v for rank of card at 0", deck.RankofCardAt(0), card1)
		}

		if deck.RankofCardAt(2) != card3 {
			t.Errorf("got %v want %v for rank of card at 2", deck.RankofCardAt(2), card3)
		}
	})

	t.Run("deck.HighRankingCards", func(t *testing.T) {
		got := deck.HighRankingCards()
		want := []Card{card1, card3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v for high ranking cards", got, want)
		}
	})

	t.Run("deck.PercentHighRanking", func(t *testing.T) {
		got := deck.PercentHighRanking()
		want := 66.67
		if got != want {
			t.Errorf("got %f want %f for percent high ranking", got, want)
		}
	})

	t.Run("deck.RemoveCard", func(t *testing.T) {
		removed := deck.RemoveCard()
		if removed != card1 {
			t.Errorf("got %v want %v for remove card", removed, card1)
		}
	})
}
