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
			if rank != card.rank {
				t.Errorf("got %d want %d", rank, card.rank)
			}
		}
	})

	t.Run("with insufficient cards", func(t *testing.T) {
		rank := testDeck.RankofCardAt(4)
		if rank != 0 {
			t.Errorf("got %d want %d", rank, 0)
		}
	})
}

func TestRemoveCard(t *testing.T) {
	deck := Deck(testCards)
	removed := deck.RemoveCard()
	if removed != testCards[0] {
		t.Errorf("got %v want %v for remove card", removed, testCards[0])
	}

	want := Deck(testCards[1:])
	if !reflect.DeepEqual(deck, want) {
		t.Errorf("%v should have been removed from deck but was not", testCards[0])
	}
}

func TestAddCards(t *testing.T) {
	deck := Deck{testCards[1], testCards[2]}
	card4 := Card{"5", "club", 5}

	deck.AddCards([]Card{card4})

	want := Deck{testCards[1], testCards[2], card4}
	if !reflect.DeepEqual(deck, want) {
		t.Errorf("got %v want %v for deck after adding new card", deck, want)
	}
}

func TestShuffle(t *testing.T) {
	deck := Deck{testCards[0], testCards[1], testCards[2]}
	for i := 0; i < 30; i++ {
		shuffledDeck := make(Deck, len(deck))
		copy(shuffledDeck, deck)
		shuffledDeck.Shuffle()
		if !reflect.DeepEqual(deck, shuffledDeck) {
			return
		}
	}

	t.Error("failed to randomize card order")
}

func TestNewDeckFromCSV(t *testing.T) {
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
			deck, err := NewDeckFromCSV(tc.filepath)

			switch tc.wantErr {
			case true:
				if err == nil {
					t.Errorf("expected an error but didn't get one")
				}
			default:
				got := len(deck)
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
