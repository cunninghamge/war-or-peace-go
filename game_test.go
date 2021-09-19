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
	reader.Write([]byte("Player2\nPlayer1\ngo\n"))
	writer := &bytes.Buffer{}
	cardSource := "./fixtures/two_cards.csv"
	game := newGame(cardSource, reader, writer)

	game.play()

	got := writer.String()
	want := `Welcome to War!
This game will be played with 2 cards.
Enter player 1: Enter player 2: The players today are Player2 and Player1.
Type 'GO' to start the game!
----------------------------------------
`
	if !strings.HasPrefix(got, want) {
		t.Errorf("got\n%s\n  want\n%s\n", got, want)
	}
}

func TestGetPlayers(t *testing.T) {
	reader := &bytes.Buffer{}
	reader.Write([]byte("Player2\nPlayer1\n"))
	writer := &bytes.Buffer{}
	cardSource := "./fixtures/two_cards.csv"
	game := newGame(cardSource, reader, writer)

	game.getPlayers()

	got := writer.String()
	want := `Enter player 1: Enter player 2: The players today are Player2 and Player1.
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
			Name: "Player1",
			Deck: &Deck{deck.Cards[:9]},
		},
		player2: &Player{
			Name: "Player2",
			Deck: &Deck{deck.Cards[9:]},
		},
	}

	game.playTurns()

	got := writer.String()
	want := `Turn 1: Player1 played the 2 of spades and Player2 played the 3 of clubs
		Player2 won 2 cards
Turn 2: Player1 played the 9 of clubs and Player2 played the 9 of spades
		WAR - Player1 played the 5 of diamonds and Player2 played the 4 of clubs
		Player1 won 8 cards
Turn 3: Player1 played the 10 of diamonds and Player2 played the 10 of hearts
		WAR - Player1 played the 7 of clubs and Player2 played the 7 of hearts
		WAR - Player1 played the 6 of hearts and Player2 played the 6 of spades
		WAR - Player1 played the 5 of spades and Player2 played the 5 of hearts
		*mutually assured destruction* 8 cards removed from play
Turn 4: Player1 played the 9 of clubs and Player2 played the 2 of spades
		Player1 won 2 cards
Turn 5: Player1 played the 9 of spades and Player2 played the 3 of clubs
		Player1 won 2 cards
`
	if got != want {
		t.Errorf("got\n%s\nwant\n%s", got, want)
	}
}

func TestCardsPlayed(t *testing.T) {
	testCases := map[string]struct {
		player1Cards []Card
		player2Cards []Card
		want         string
	}{
		"basic turn": {
			player1Cards: []Card{
				{"Queen", "diamond", 12},
			},
			player2Cards: []Card{
				{"10", "heart", 10},
			},
			want: "Turn 1: Player1 played the Queen of diamonds and Player2 played the 10 of hearts\n",
		},
		"war: first card wins": {
			player1Cards: []Card{
				{"Queen", "diamond", 12},
				{"2", "club", 2},
				{"7", "spade", 7},
				{"4", "heart", 4},
			},
			player2Cards: []Card{
				{"Queen", "heart", 12},
				{"10", "diamond", 10},
				{"5", "spade", 5},
				{"3", "club", 3},
			},
			want: "Turn 1: Player1 played the Queen of diamonds and Player2 played the Queen of hearts\n" +
				"		WAR - Player1 played the 4 of hearts and Player2 played the 3 of clubs\n",
		},
		"war: second card wins": {
			player1Cards: []Card{
				{"Queen", "diamond", 12},
				{"2", "club", 2},
				{"7", "spade", 7},
				{"3", "heart", 3},
			},
			player2Cards: []Card{
				{"Queen", "heart", 12},
				{"10", "diamond", 10},
				{"5", "spade", 5},
				{"3", "club", 3},
			},
			want: "Turn 1: Player1 played the Queen of diamonds and Player2 played the Queen of hearts\n" +
				"		WAR - Player1 played the 3 of hearts and Player2 played the 3 of clubs\n" +
				"		WAR - Player1 played the 7 of spades and Player2 played the 5 of spades\n",
		},
		"war: last card wins": {
			player1Cards: []Card{
				{"Queen", "diamond", 12},
				{"10", "diamond", 10},
				{"7", "spade", 7},
				{"3", "heart", 3},
			},
			player2Cards: []Card{
				{"Queen", "heart", 12},
				{"2", "club", 2},
				{"7", "club", 7},
				{"3", "club", 3},
			},
			want: "Turn 1: Player1 played the Queen of diamonds and Player2 played the Queen of hearts\n" +
				"		WAR - Player1 played the 3 of hearts and Player2 played the 3 of clubs\n" +
				"		WAR - Player1 played the 7 of spades and Player2 played the 7 of clubs\n" +
				"		WAR - Player1 played the 10 of diamonds and Player2 played the 2 of clubs\n",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			game := Game{
				scanner: bufio.NewScanner(&bytes.Buffer{}),
				writer:  writer,
				player1: &Player{
					Name: "Player1",
					Deck: &Deck{tc.player1Cards},
				},
				player2: &Player{
					Name: "Player2",
					Deck: &Deck{tc.player2Cards},
				},
			}

			game.cardsPlayed(1)

			got := writer.String()
			want := tc.want
			if got != want {
				t.Errorf("got\n%s\nwant\n%s", got, want)
			}
		})
	}
}

func TestDeclareWinner(t *testing.T) {
	writer := &bytes.Buffer{}
	game := Game{
		scanner: bufio.NewScanner(&bytes.Buffer{}),
		writer:  writer,
		player1: &Player{
			Name: "Player1",
			Deck: &Deck{testCards[:1]},
		},
		player2: &Player{
			Name: "Player2",
			Deck: &Deck{[]Card{}},
		},
	}

	game.declareWinner()

	got := writer.String()
	want := `Player2 is out of cards!
*~*~*~* Player1 won the game! *~*~*~*
`
	if got != want {
		t.Errorf("got\n%s\nwant\n%s", got, want)
	}
}
