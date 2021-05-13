package main

import (
	"math"
)

type Card struct {
	Suit  string
	Value string
	Rank  int
}

type Deck struct {
	Cards []Card
}

func (d Deck) RankofCardAt(index int) Card {
	return d.Cards[index]
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
