package main

import (
	"os"
	"reflect"
	"testing"
)

func TestRankOfCardAt(t *testing.T) {
	t.Run("with cards", func(t *testing.T) {
		for i, card := range testCards {
			rank := testDeck.RankofCardAt(i)
			if rank != card.Rank {
				t.Errorf("got %d want %d", rank, card.Rank)
			}
		}
	})

	t.Run("with insufficient cards", func(t *testing.T) {
		rank := testDeck.RankofCardAt(3)
		if rank != 0 {
			t.Errorf("got %d want %d", rank, 0)
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
	want := 66.67
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
	if newPct != 33.33 {
		t.Errorf("newPct %f want %f for new percent high ranking", newPct, 33.3)
	}
}

func TestNewFromCSV(t *testing.T) {
	os.WriteFile("reader_err.csv", []byte("a,b,c\nd,e"), 0644)
	defer os.Remove("reader_err.csv")
	os.WriteFile("card_err.csv", []byte("a,b\nd,e"), 0644)
	defer os.Remove("card_err.csv")

	testCases := map[string]struct {
		length   int
		filepath string
		wantErr  bool
	}{
		"with a file": {
			filepath: "fixtures/two_cards.csv",
			length:   2,
		},
		"using default file": {
			filepath: "",
			length:   52,
		},
		"with file reading error": {
			filepath: "reader_err.csv",
			wantErr:  true,
		},
		"card creation error": {
			filepath: "card_err.csv",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.filepath, func(t *testing.T) {
			var deck = &Deck{}
			err := deck.NewFromCSV(tc.filepath)

			switch tc.wantErr {
			case true:
				if err == nil {
					t.Errorf("expected an error but didn't get one")
				}
			default:
				got := len(deck.Cards)
				want := tc.length
				if got != want {
					t.Errorf("got %d want %d", got, want)
				}
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}
