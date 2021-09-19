package main

type Player struct {
	name string
	deck *Deck
}

func (p Player) HasLost() bool {
	return len(p.deck.cards) == 0
}
