package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	welcome   = "Welcome to War!\n"
	lineBreak = strings.Repeat("-", 40) + "\n"
)

type Game struct {
	scanner          *bufio.Scanner
	writer           io.Writer
	deck             Deck
	player1, player2 *Player
}

func newGame(filepath string, r io.Reader, w io.Writer) Game {
	scanner := bufio.NewScanner(r)
	var deck Deck
	deck.NewFromCSV(filepath)
	return Game{
		scanner: scanner,
		writer:  w,
		deck:    deck,
		player1: &Player{Deck: &Deck{[]Card{}}},
		player2: &Player{Deck: &Deck{[]Card{}}},
	}
}

func (g Game) play() {
	fmt.Fprintf(g.writer, "%sThis game will be played with %d cards.\n", welcome, len(g.deck.Cards))

	g.getPlayers()
	g.deal()

	var cmd string
	for cmd != "GO" {
		fmt.Fprint(g.writer, "Type 'GO' to start the game!\n"+lineBreak)
		g.scanner.Scan()
		cmd = strings.ToUpper(g.scanner.Text())
		if cmd == "Q" {
			os.Exit(0)
		}
	}

	g.playTurns()
	g.declareWinner()
}

func (g Game) getPlayers() {
	fmt.Fprint(g.writer, "Enter player 1: ")
	g.scanner.Scan()
	g.player1.Name = g.scanner.Text()

	fmt.Fprint(g.writer, "Enter player 2: ")
	g.scanner.Scan()
	g.player2.Name = g.scanner.Text()

	fmt.Fprintf(g.writer, "The players today are %s and %s.\n",
		g.player1.Name, g.player2.Name)
}

func (g Game) deal() {
	g.deck.Shuffle()
	for i := 0; i < len(g.deck.Cards); i++ {
		switch i%2 == 0 {
		case true:
			g.player1.Deck.AddCards([]Card{g.deck.Cards[i]})
		case false:
			g.player2.Deck.AddCards([]Card{g.deck.Cards[i]})
		}
	}
}

func (g Game) playTurns() {
	for i := 1; !g.player1.HasLost() && !g.player2.HasLost(); i++ {
		turn := Turn{Player1: g.player1, Player2: g.player2}
		winner := turn.Winner()

		g.cardsPlayed(i)
		switch turn.Type() {
		case basic:
			fmt.Fprintf(g.writer, "		%s won 2 cards\n", winner.Name)
		case war:
			fmt.Fprintf(g.writer, "		%s won 8 cards\n", winner.Name)
		case mutuallyAssuredDestruction:
			fmt.Fprint(g.writer, "		*mutually assured destruction* 8 cards removed from play\n")
		}

		turn.PileCards()
		turn.AwardSpoils(winner)
	}
}

func (g Game) declareWinner() {
	var winner, loser *Player
	switch g.player1.HasLost() {
	case true:
		winner, loser = g.player2, g.player1
	default:
		winner, loser = g.player1, g.player2
	}
	fmt.Fprintf(g.writer, "%s is out of cards!\n", loser.Name)
	fmt.Fprintf(g.writer, "*~*~*~* %s won the game! *~*~*~*\n", winner.Name)
}

func (g Game) cardsPlayed(turnNumber int) {
	for _, i := range []int{0, 3, 2, 1} {
		prefix := "		WAR -"
		if i == 0 {
			prefix = fmt.Sprintf("Turn %d:", turnNumber)
		}
		player1Card := g.player1.Deck.Cards[i]
		player2Card := g.player2.Deck.Cards[i]
		fmt.Fprintf(g.writer, "%s %s played %s and %s played %s\n",
			prefix,
			g.player1.Name, player1Card,
			g.player2.Name, player2Card,
		)
		if player1Card.Rank != player2Card.Rank {
			break
		}
	}
}
