package main

type Player struct {
	name    string
	deck    *Deck
	hasLost bool
}

func (p Player) CardsLeft() int {
	return len(p.deck.cards)
}

// TODO get rid of this method
func (p Player) HasLost() bool {
	return p.hasLost || len(p.deck.cards) == 0
}
