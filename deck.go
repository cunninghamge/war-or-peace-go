package main

import (
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

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func NewDeckFromCSV(filepath string) (*Deck, error) {
	var source = "fixtures/cards.csv"
	if len(filepath) > 0 {
		source = filepath
	}
	records, err := reader.ReadFile(source)
	if err != nil {
		return nil, err
	}

	cards, err := createCards(records)
	if err != nil {
		return nil, err
	}

	return &Deck{cards}, nil
}
