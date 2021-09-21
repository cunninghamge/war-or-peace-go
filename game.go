package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Game struct {
	scanner          *bufio.Scanner
	writer           io.Writer
	player1, player2 *Player
}

func newGame(filepath string, r io.Reader, w io.Writer) Game {
	game := Game{
		scanner: bufio.NewScanner(r),
		writer:  w,
		player1: &Player{deck: &Deck{}},
		player2: &Player{deck: &Deck{}},
	}

	deck, err := NewDeckFromCSV(filepath)
	if err != nil {
		exitWithError(err)
	}
	fmt.Fprintf(game.writer, "Welcome to War!\nThis game will be played with %d cards.\n", len(deck))

	game.getPlayers()
	game.deal(deck)

	return game
}

func (g Game) getPlayers() {
	fmt.Fprintf(g.writer, "Enter player %d: ", 1)
	g.scanner.Scan()
	g.player1.name = g.scanner.Text()

	fmt.Fprintf(g.writer, "Enter player %d: ", 2)
	g.scanner.Scan()
	g.player2.name = g.scanner.Text()

	fmt.Fprintf(g.writer, "The players today are %s and %s.\n",
		g.player1.name, g.player2.name)
}

func (g Game) deal(deck Deck) {
	deck.Shuffle()
	for i := 0; i < len(deck); i++ {
		switch i%2 == 0 {
		case true:
			g.player1.deck.AddCards([]Card{deck[i]})
		case false:
			g.player2.deck.AddCards([]Card{deck[i]})
		}
	}
}

func (g Game) start() {
	var cmd string
	for cmd != "GO" {
		fmt.Fprint(g.writer, "Type 'GO' to start the game!\n"+
			strings.Repeat("-", 40)+"\n")
		g.scanner.Scan()
		cmd = strings.ToUpper(g.scanner.Text())
		if cmd == "Q" {
			return
		}
	}

	g.play()
}

func (g Game) play() {
	for i := 1; !g.player1.HasLost() && !g.player2.HasLost(); i++ {
		turn := Turn{player1: g.player1, player2: g.player2}
		fmt.Fprintf(g.writer, "Turn %d: %s has %d cards left and %s has %d cards left\n", i,
			g.player1.name, g.player1.CardsLeft(),
			g.player2.name, g.player2.CardsLeft(),
		)
		g.showCards(0, "		")

		if turn.Type() != basic {
			g.war()
		}
		g.awardSpoils(turn)

		if i%500 == 0 {
			g.player1.deck.Shuffle()
			g.player2.deck.Shuffle()
		}
	}
	g.displayResult()
}

func (g Game) showCards(i int, prefix string) {
	fmt.Fprintf(g.writer, "%s%s played %s and %s played %s\n",
		prefix,
		g.player1.name, (*g.player1.deck)[i],
		g.player2.name, (*g.player2.deck)[i],
	)
}

func (g Game) canPlayWar() bool {
	var canPlayWar = true
	for _, player := range []*Player{g.player1, g.player2} {
		if player.CardsLeft() <= 3 {
			fmt.Fprintf(g.writer, "		%s does not have enough cards for war!\n", player.name)
			player.lost = true
			canPlayWar = false
		}
	}
	return canPlayWar
}

func (g Game) war() {
	if canPlayWar := g.canPlayWar(); !canPlayWar {
		return
	}
	for i := 3; i > 0; i-- {
		g.showCards(i, "	WAR!	")
		if g.player1.deck.RankofCardAt(i) != g.player2.deck.RankofCardAt(i) {
			return
		}
	}
}

func (g Game) awardSpoils(turn Turn) {
	winner := turn.Winner()
	turnType := turn.Type()
	turn.PileCards()
	turn.AwardSpoils(winner)
	switch turnType {
	case mutuallyAssuredDestruction:
		fmt.Fprintf(g.writer, "		*mutually assured destruction* %d cards removed from play\n\n", len(turn.spoilsOfWar))
	default:
		fmt.Fprintf(g.writer, "		%s won %d cards\n\n", winner.name, len(turn.spoilsOfWar))
	}
}

func (g Game) displayResult() {
	for _, player := range []*Player{g.player1, g.player2} {
		if player.CardsLeft() == 0 {
			fmt.Fprintf(g.writer, "%s is out of cards!\n", player.name)
		}
		if !player.HasLost() {
			defer fmt.Fprintf(g.writer, "*~*~*~* %s won the game! *~*~*~*\n", player.name)
		}
	}

	if g.player1.HasLost() && g.player2.HasLost() {
		fmt.Fprintf(g.writer, "*~*~*~* It's a draw! *~*~*~*\n")
	}
}
