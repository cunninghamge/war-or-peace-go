package main

import (
	"math/rand"
	"time"
	"war-or-peace/reader"
)

type Deck []Card

func (d Deck) RankofCardAt(index int) int {
	if len(d) <= index {
		return 0
	}
	return d[index].rank
}

func (d *Deck) RemoveCard() Card {
	cards := *d
	card, cards := cards[0], cards[1:]
	*d = cards
	return card
}

func (d *Deck) AddCards(newCards []Card) {
	cards := []Card{}
	cards = append(cards, *d...)
	cards = append(cards, newCards...)
	*d = cards
}

func (d *Deck) Shuffle() {
	cards := *d
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	*d = cards
}

func NewDeckFromCSV(filepath string) (Deck, error) {
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

	return Deck(cards), nil
}
