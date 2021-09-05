package main

import (
	"errors"
	"fmt"
	"math"
)

var ErrOutOfRange = errors.New("not enough cards")

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

func (d Deck) RankofCardAt(index int) (int, error) {
	if len(d.Cards) <= index {
		return 0, ErrOutOfRange
	}
	return d.Cards[index].Rank, nil
}

func (d Deck) HighRankingCards() []Card {
	var cards []Card
	for _, card := range d.Cards {
		if card.Rank > 10 {
			cards = append(cards, card)
		}
	}
	return cards
}

func (d Deck) PercentHighRanking() float64 {
	numHighRanking := len(d.HighRankingCards())
	numCards := len(d.Cards)
	pct := float64(numHighRanking) / float64(numCards)
	return math.Round(pct*1000) / 10
}

func (d *Deck) RemoveCard() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

func (d *Deck) AddCard(card Card) {
	d.Cards = append(d.Cards, card)
}
