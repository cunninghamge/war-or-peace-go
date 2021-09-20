package main

type Player struct {
	name string
	deck *Deck
	lost bool
}

func (p Player) CardsLeft() int {
	return len(p.deck.cards)
}

func (p Player) HasLost() bool {
	return p.lost || len(p.deck.cards) == 0
}
