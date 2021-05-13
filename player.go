package main

type Player struct {
	Name string
	Deck *Deck
}

func (p Player) HasLost() bool {
	return len(p.Deck.Cards) == 0
}
