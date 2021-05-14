package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Start() {
	player1, player2 := NewPlayers()
	fmt.Println("Welcome to War! (or Peace) This game wil be played with 52 cards.")
	fmt.Printf("The players today are %s and %s.\n", player1.Name, player2.Name)
	fmt.Println("Type 'GO' to start the game!")
	fmt.Println("-----------------------------------------------------------------")

	for i := 1; i < 10_000; i++ {
		PlayTurn(player1, player2, i)

		gameOver := GameOver(player1, player2, i)
		if gameOver {
			break
		}
	}
}

func NewPlayers() (Player, Player) {
	cards := NewDeck()
	return Player{"Cacco", &Deck{cards[:26]}}, Player{"Pickles", &Deck{cards[26:]}}
}

func NewDeck() []Card {
	suits := []string{"heart", "spade", "diamond", "club"}
	values := map[string]int{
		"Ace":   14,
		"King":  13,
		"Queen": 12,
		"Jack":  11,
		"10":    10,
		"9":     9,
		"8":     8,
		"7":     7,
		"6":     6,
		"5":     5,
		"4":     4,
		"3":     3,
		"2":     2,
	}

	var cards []Card
	for _, suit := range suits {
		for s, n := range values {
			cards = append(cards, Card{suit, s, n})
		}
	}

	// shuffle cards
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := len(cards); i > 0; i-- {
		randIndex := r.Intn(i)
		cards[i-1], cards[randIndex] = cards[randIndex], cards[i-1]
	}

	return cards
}

func PlayTurn(player1, player2 Player, i int) {
	t := &Turn{player1, player2, []Card{}}
	p1 := t.Player1.Name
	p2 := t.Player2.Name
	c1 := t.Player1.Deck.Cards[0].String()
	c2 := t.Player2.Deck.Cards[0].String()
	fmt.Printf("Turn %d: %s played %s, and %s played %s\n", i, p1, c1, p2, c2)

	switch t.Type() {
	case "basic":
		BasicTurn(t)
	case "war":
		WarTurn(t)
	case "mutually assured destruction":
		MADTurn(t)
	}
}

func BasicTurn(t *Turn) {
	winner := t.Winner()
	fmt.Printf("   %s won 2 cards\n", winner.Name)
	t.PileCards()
	t.AwardSpoils(winner)
}

func WarTurn(t *Turn) {
	p1 := t.Player1.Name
	p2 := t.Player2.Name
	winner := t.Winner()
	if len(t.Player2.Deck.Cards) < 3 {
		fmt.Printf("   WAR - %s doesn't have enough cards!\n", p2)
		t.PileCards()
	} else if len(t.Player1.Deck.Cards) < 3 {
		fmt.Printf("   WAR - %s doesn't have enough cards!\n", p1)
		t.PileCards()
	} else {
		c1 := t.Player1.Deck.Cards[2].String()
		c2 := t.Player2.Deck.Cards[2].String()
		fmt.Printf("   WAR - %s played %s, and %s played %s\n", p1, c1, p2, c2)
		fmt.Printf("   %s won 6 cards\n", winner.Name)
		t.PileCards()
		t.AwardSpoils(winner)
	}
}

func MADTurn(t *Turn) {
	p1 := t.Player1.Name
	p2 := t.Player2.Name
	c1 := t.Player1.Deck.Cards[2].String()
	c2 := t.Player2.Deck.Cards[2].String()
	t.PileCards()
	fmt.Printf("   WAR - %s played %s, and %s played %s\n", p1, c1, p2, c2)
	fmt.Println("   *mutually assured destruction* 6 cards removed from play")
}

func GameOver(player1, player2 Player, i int) bool {
	if player1.HasLost() {
		fmt.Printf("*~*~*~* %s has won the game! *~*~*~*", player2.Name)
		return true
	}
	if player2.HasLost() {
		fmt.Printf("*~*~*~* %s has won the game! *~*~*~*", player1.Name)
		return true
	}

	l1 := len(player1.Deck.Cards)
	l2 := len(player2.Deck.Cards)
	fmt.Printf("   %s has %d cards and %s has %d cards\n", player1.Name, l1, player2.Name, l2)
	if i == 9_999 {
		fmt.Printf("Maxiumum turns exceeded.\nGame Over!")
	}
	return false
}
