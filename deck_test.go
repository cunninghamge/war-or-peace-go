package main

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	testCases := map[string]Card{
		"the Queen of diamonds": testCards[0],
		"the 3 of spades":       testCards[1],
		"the Ace of hearts":     testCards[2],
	}

	for str, card := range testCases {
		t.Run(str, func(t *testing.T) {
			got := card.String()
			if got != str {
				t.Errorf("got %q str %q", got, str)
			}
		})
	}
}

func TestRankOfCardAt(t *testing.T) {
	t.Run("cards in deck", func(t *testing.T) {
		for i, card := range testCards {
			rank, err := testDeck.RankofCardAt(i)
			if rank != card.Rank {
				t.Errorf("got %d want %d", rank, card.Rank)
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}
	})

	t.Run("index out of range", func(t *testing.T) {
		rank, err := testDeck.RankofCardAt(3)
		if rank != 0 {
			t.Errorf("got %d want %d", rank, 0)
		}
		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func TestHighRankingCards(t *testing.T) {
	got := testDeck.HighRankingCards()
	want := []Card{testCards[0], testCards[2]}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPercentHighRanking(t *testing.T) {
	got := testDeck.PercentHighRanking()
	want := 66.7
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestRemoveCard(t *testing.T) {
	deck := Deck{testCards}
	removed := deck.RemoveCard()
	if removed != testCards[0] {
		t.Errorf("got %v want %v for remove card", removed, testCards[0])
	}

	want := []Card{testCards[1], testCards[2]}
	if !reflect.DeepEqual(deck.Cards, want) {
		t.Errorf("%v should have been removed from deck but was not", testCards[0])
	}

	got := deck.HighRankingCards()
	if !reflect.DeepEqual(got, []Card{testCards[2]}) {
		t.Errorf("got %v want %v for new high ranking cards", got, []Card{testCards[2]})
	}

	newPct := deck.PercentHighRanking()
	if newPct != 50.00 {
		t.Errorf("newPct %f want %f for new percent high ranking", newPct, 50.00)
	}
}

func TestAddCard(t *testing.T) {
	deck := Deck{[]Card{testCards[1], testCards[2]}}
	card4 := Card{"club", "5", 5}

	deck.AddCard(card4)

	want := []Card{testCards[1], testCards[2], card4}
	if !reflect.DeepEqual(deck.Cards, want) {
		t.Errorf("got %v want %v for deck after adding new card", deck.Cards, want)
	}

	got := deck.HighRankingCards()
	if !reflect.DeepEqual(got, []Card{testCards[2]}) {
		t.Errorf("got %v want %v for new high ranking cards", got, []Card{testCards[3]})
	}

	newPct := deck.PercentHighRanking()
	if newPct != 33.3 {
		t.Errorf("newPct %f want %f for new percent high ranking", newPct, 33.3)
	}
}
