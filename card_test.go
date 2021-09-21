package main

import "testing"

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

func TestCreateCards(t *testing.T) {
	testCases := map[string]struct {
		records    [][]string
		deckLength int
		wantError  string
	}{
		"success": {
			records: [][]string{
				{"value", "suit", "1"},
				{"value", "suit", "1"},
			},
			deckLength: 2,
		},
		"too few fields": {
			records: [][]string{
				{"value", "suit"},
				{"value", "suit"},
			},
			wantError: errInvalidRecords,
		},
		"too many fields": {
			records: [][]string{
				{"value", "suit", "1", "extra field"},
				{"value", "suit", "1", "extra field"},
			},
			wantError: errInvalidRecords,
		},
		"different numbers of fields": {
			records: [][]string{
				{"value", "suit"},
				{"value", "suit", "1", "extra field"},
				{"value", "suit", "1"},
			},
			wantError: errInvalidRecords,
		},
		"rank can't be converted to int": {
			records: [][]string{
				{"value", "suit", "17x"},
			},
			wantError: "strconv.Atoi: parsing \"17x\": invalid syntax",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			cards, err := createCards(tc.records)
			got := len(cards)
			if got != tc.deckLength {
				t.Errorf("created %d card but should have created %d", got, tc.deckLength)
			}

			wantErr := len(tc.wantError) > 0
			if wantErr && err.Error() != tc.wantError {
				t.Errorf("got %v want %s", err, tc.wantError)
			}

			if !wantErr && err != nil {
				t.Errorf("got %v want %v", err, nil)
			}
		})
	}
}
