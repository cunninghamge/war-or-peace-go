package main

import (
	"math"
	"math/rand"
	"time"
	"war-or-peace/reader"
)

// TODO change deck type from struct to []Card
type Deck struct {
	cards []Card
}

func (d Deck) RankofCardAt(index int) int {
	if len(d.cards) <= index {
		return 0
	}
	return d.cards[index].rank
}

func (d Deck) HighRankingCards() []Card {
	var cards []Card
	for _, card := range d.cards {
		if card.rank > 10 {
			cards = append(cards, card)
		}
	}
	return cards
}

func (d Deck) PercentHighRanking() float64 {
	numHighRanking := len(d.HighRankingCards())
	numCards := len(d.cards)
	pct := float64(numHighRanking) / float64(numCards)
	return math.Round(pct*10000) / 100
}

func (d *Deck) RemoveCard() Card {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

func (d *Deck) AddCards(newCards []Card) {
	tempCards := make([]Card, 0, len(d.cards)+len(newCards))
	tempCards = append(tempCards, d.cards...)
	tempCards = append(tempCards, newCards...)
	d.cards = tempCards
}

func (d Deck) Shuffle() Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
	return d
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

	d.cards, err = createCards(records)
	if err != nil {
		return err
	}

	return nil
}
