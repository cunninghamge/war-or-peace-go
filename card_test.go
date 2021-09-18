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
		wantError  bool
	}{
		"success": {
			records: [][]string{
				{"suit", "value", "1"},
				{"suit", "value", "1"},
			},
			deckLength: 2,
		},
		"too few fields": {
			records: [][]string{
				{"suit", "value"},
				{"suit", "value"},
			},
			wantError: true,
		},
		"too many fields": {
			records: [][]string{
				{"suit", "value", "1", "extra field"},
				{"suit", "value", "1", "extra field"},
			},
			wantError: true,
		},
		"different numbers of fields": {
			records: [][]string{
				{"suit", "value"},
				{"suit", "value", "1", "extra field"},
				{"suit", "value", "1"},
			},
			wantError: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			cards, err := createCards(tc.records)
			got := len(cards)
			if got != tc.deckLength {
				t.Errorf("created %d card but should have created %d", got, tc.deckLength)
			}

			if tc.wantError && err.Error() != errInvalidRecords {
				t.Errorf("got %v want %s", err, errInvalidRecords)
			}

			if !tc.wantError && err != nil {
				t.Errorf("got %v want %v", err, nil)
			}
		})
	}
}
