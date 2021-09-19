package main

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

var (
	testDeck  Deck
	testCards []Card
)

func TestMain(m *testing.M) {
	testCards = []Card{
		{"Queen", "diamond", 12},
		{"3", "spade", 3},
		{"Ace", "heart", 14},
	}
	testDeck = Deck{testCards[:3]}

	code := m.Run()
	os.Exit(code)
}

func TestArg(t *testing.T) {
	testCases := map[string]struct {
		args   []string
		result string
	}{
		"with no filepath": {
			args:   []string{"exec"},
			result: "",
		},
		"with a filepath": {
			args:   []string{"exec", "file.txt"},
			result: "file.txt",
		},
		"with an extra arg": {
			args:   []string{"exec", "file.txt", "extra"},
			result: "file.txt",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := arg(tc.args)
			if got != tc.result {
				t.Errorf("got %s want %s", got, tc.result)
			}
		})
	}
}

func TestExitWithError(t *testing.T) {
	if os.Getenv("OS_EXIT_CALLED") == "1" {
		exitWithError(errors.New("error"))
		return
	}
	subTest := exec.Command(os.Args[0], "-test.run=TestExitWithError")
	subTest.Env = append(os.Environ(), "OS_EXIT_CALLED=1")
	err := subTest.Run()
	if exitError, ok := err.(*exec.ExitError); !ok || exitError.Success() {
		t.Error("process exited with no error, wanted exit status 1")
	}
}

// func TestNewDeck(t *testing.T) {
// 	newDeck := NewDeck()

// 	if len(newDeck) != 52 {
// 		t.Errorf("got %d not 52 for deck length", len(newDeck))
// 	}

// 	var hearts []Card
// 	var diamonds []Card
// 	var spades []Card
// 	var clubs []Card
// 	for _, card := range newDeck {
// 		switch card.Suit {
// 		case "heart":
// 			hearts = append(hearts, card)
// 		case "diamond":
// 			diamonds = append(diamonds, card)
// 		case "spade":
// 			spades = append(spades, card)
// 		case "club":
// 			clubs = append(clubs, card)
// 		}
// 	}

// 	for _, suit := range [][]Card{hearts, diamonds, spades, clubs} {
// 		if len(suit) != 13 {
// 			t.Errorf("got %d not 13 for cards in suit", len(suit))
// 		}
// 	}
// }

// func TestNewPlayers(t *testing.T) {
// 	cards := NewDeck()
// 	player1, player2 := NewPlayers(cards)

// 	if player1.Name != "Cacco" {
// 		t.Errorf("got %q want %q for player1 name", player1.Name, "Cacco")
// 	}

// 	if player2.Name != "Pickles" {
// 		t.Errorf("got %q want %q for player2 name", player2.Name, "Pickles")
// 	}

// 	if len(player1.Deck.Cards) != 26 {
// 		t.Errorf("got %d not 26 for player1's cards", len(player1.Deck.Cards))
// 	}

// 	if len(player2.Deck.Cards) != 26 {
// 		t.Errorf("got %d not 26 for player2's cards", len(player2.Deck.Cards))
// 	}
// }
