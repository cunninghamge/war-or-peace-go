package main

import (
	"bufio"
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestNewGame(t *testing.T) {
	reader := &bytes.Buffer{}
	writer := &bytes.Buffer{}
	reader.Write([]byte("some text"))

	game := newGame("", reader, writer)

	if !reflect.DeepEqual(game.writer, writer) {
		t.Errorf("got %v want %v", game.writer, writer)
	}
}

func TestStart(t *testing.T) {
	reader := &bytes.Buffer{}
	reader.Write([]byte("Cacco\nPickles\ngo\n"))
	writer := &bytes.Buffer{}
	cardSource := "./fixtures/two_cards.csv"
	game := newGame(cardSource, reader, writer)

	game.start()

	got := writer.String()
	want := `Welcome to War!
This game will be played with 2 cards.
Enter player 1: Enter player 2: The players today are Cacco and Pickles.
Type 'GO' to start the game!
----------------------------------------
`
	if !strings.HasPrefix(got, want) {
		t.Errorf("got\n%s\n  want\n%s\n", got, want)
	}
}

func TestGetPlayers(t *testing.T) {
	reader := &bytes.Buffer{}
	reader.Write([]byte("Cacco\nPickles\n"))
	writer := &bytes.Buffer{}
	cardSource := "./fixtures/two_cards.csv"
	game := newGame(cardSource, reader, writer)

	game.getPlayers()

	got := writer.String()
	want := `Enter player 1: Enter player 2: The players today are Cacco and Pickles.
`
	if got != want {
		t.Errorf("got\n%s\n  want\n%s\n", got, want)
	}
}

func TestPlay(t *testing.T) {
	writer := &bytes.Buffer{}
	var deck Deck
	deck.NewFromCSV("fixtures/test_cards.csv")
	game := Game{
		scanner: bufio.NewScanner(&bytes.Buffer{}),
		writer:  writer,
		player1: &Player{
			Name: "Pickles",
			Deck: &Deck{deck.Cards[:7]},
		},
		player2: &Player{
			Name: "Cacco",
			Deck: &Deck{deck.Cards[7:]},
		},
	}

	game.play()

	got := writer.String()
	want := `Turn 1: Pickles played the 2 of spades and Cacco played the 3 of clubs
		Cacco won 2 cards
Turn 2: Pickles played the Ace of hearts and Cacco played the Ace of spades
		WAR - Pickles played the 5 of diamonds and Cacco played the 4 of clubs
		Pickles won 6 cards
Turn 3: Pickles played the 5 of spades and Cacco played the 5 of hearts
		WAR - Pickles played the 7 of clubs and Cacco played the 7 of hearts
		*mutually assured destruction* 6 cards removed from play
Turn 4: Pickles played the Ace of hearts and Cacco played the 2 of spades
		Pickles won 2 cards
Turn 5: Pickles played the Ace of spades and Cacco played the 3 of clubs
		Pickles won 2 cards
Cacco is out of cards!
*~*~*~* Pickles won the game! *~*~*~*
`
	if got != want {
		t.Errorf("got\n%s\nwant\n%s", got, want)
	}
}
