package main

import (
	"bytes"
	"reflect"
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
	if got != want {
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
