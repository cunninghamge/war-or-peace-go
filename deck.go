package main

import (
	"math"
	"war-or-peace/reader"
)

type Deck struct {
	Cards []Card
}

func (d Deck) RankofCardAt(index int) int {
	if len(d.Cards) <= index {
		return 0
	}
	return d.Cards[index].Rank
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
	return math.Round(pct*10000) / 100
}

func (d *Deck) RemoveCard() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

func (d *Deck) AddCard(card Card) {
	d.Cards = append(d.Cards, card)
}

func (d *Deck) NewFromCSV(filepath string) error {
	var source = "fixtures/cards.csv"
	if len(filepath) > 0 {
		source = filepath
	}
	records, err := reader.ReadFile(source)
	if err != nil {
		return err
	}

	d.Cards, err = createCards(records)
	if err != nil {
		return err
	}

	return nil
}
