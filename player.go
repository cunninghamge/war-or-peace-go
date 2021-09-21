package main

type Player struct {
	name string
	deck *Deck
	lost bool
}

func (p Player) CardsLeft() int {
	return len(*p.deck)
}

func (p Player) HasLost() bool {
	return p.lost || len(*p.deck) == 0
}
