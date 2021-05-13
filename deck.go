package main

import (
	"fmt"
	"math"
)

type Card struct {
	Suit  string
	Value string
	Rank  int
}

func (c Card) String() string {
	return fmt.Sprintf("the %s of %ss", c.Value, c.Suit)
}

type Deck struct {
	Cards []Card
}

func (d Deck) RankofCardAt(index int) int {
	return d.Cards[index].Rank
}

func (d Deck) HighRankingCards() []Card {
	var cards []Card
	for _, card := range d.Cards {
		if card.Rank >= 11 {
			cards = append(cards, card)
		}
	}
	return cards
}

func (d Deck) PercentHighRanking() float64 {
	numHighRanking := float64(len(d.HighRankingCards()))
	numCards := float64(len(d.Cards))
	return math.Round((numHighRanking/numCards*100.00)/0.01) * 0.01
}

func (d *Deck) RemoveCard() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

func (d *Deck) AddCard(card Card) {
	d.Cards = append(d.Cards, card)
}
