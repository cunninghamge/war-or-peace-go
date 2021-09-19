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
		player1: &Player{deck: &Deck{[]Card{}}},
		player2: &Player{deck: &Deck{[]Card{}}},
	}
}

func (g Game) play() {
	fmt.Fprintf(g.writer, "%sThis game will be played with %d cards.\n", welcome, len(g.deck.cards))

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
	g.player1.name = g.scanner.Text()

	fmt.Fprint(g.writer, "Enter player 2: ")
	g.scanner.Scan()
	g.player2.name = g.scanner.Text()

	fmt.Fprintf(g.writer, "The players today are %s and %s.\n",
		g.player1.name, g.player2.name)
}

func (g Game) deal() {
	g.deck.Shuffle()
	for i := 0; i < len(g.deck.cards); i++ {
		switch i%2 == 0 {
		case true:
			g.player1.deck.AddCards([]Card{g.deck.cards[i]})
		case false:
			g.player2.deck.AddCards([]Card{g.deck.cards[i]})
		}
	}
}

func (g Game) playTurns() {
	for i := 1; !g.player1.HasLost() && !g.player2.HasLost(); i++ {
		turn := Turn{player1: g.player1, player2: g.player2}
		winner := turn.Winner()

		g.cardsPlayed(i)
		switch turn.Type() {
		case basic:
			fmt.Fprintf(g.writer, "		%s won 2 cards\n", winner.name)
		case war:
			fmt.Fprintf(g.writer, "		%s won 8 cards\n", winner.name)
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
	fmt.Fprintf(g.writer, "%s is out of cards!\n", loser.name)
	fmt.Fprintf(g.writer, "*~*~*~* %s won the game! *~*~*~*\n", winner.name)
}

func (g Game) cardsPlayed(turnNumber int) {
	for _, i := range []int{0, 3, 2, 1} {
		prefix := "		WAR -"
		if i == 0 {
			prefix = fmt.Sprintf("Turn %d:", turnNumber)
		}
		player1Card := g.player1.deck.cards[i]
		player2Card := g.player2.deck.cards[i]
		fmt.Fprintf(g.writer, "%s %s played %s and %s played %s\n",
			prefix,
			g.player1.name, player1Card,
			g.player2.name, player2Card,
		)
		if player1Card.rank != player2Card.rank {
			break
		}
	}
}
