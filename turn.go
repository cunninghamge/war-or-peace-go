package main

type Turn struct {
	Player1, Player2 Player
	SpoilsOfWar      []Card
}

func (t Turn) Type() string {
	r1, _ := t.Player1.Deck.RankofCardAt(0)
	r2, _ := t.Player2.Deck.RankofCardAt(0)
	if r1 != r2 {
		return "basic"
	}

	r1, err1 := t.Player1.Deck.RankofCardAt(2)
	r2, err2 := t.Player2.Deck.RankofCardAt(2)
	if err1 != nil || err2 != nil || r1 != r2 {
		return "war"
	}

	return "mutually assured destruction"
}

func (t Turn) Winner() Player {
	var index int
	switch t.Type() {
	case "basic":
		index = 0
	case "war":
		index = 2
	case "mutually assured destruction":
		return Player{Name: "No Winner"}
	}

	r1, _ := t.Player1.Deck.RankofCardAt(index)
	r2, err2 := t.Player2.Deck.RankofCardAt(index)
	if err2 != nil || r1 > r2 {
		return t.Player1
	} else {
		return t.Player2
	}
}

func (t *Turn) PileCards() {
	var numCards int = 0
	switch t.Type() {
	case "basic":
		numCards = 1
	case "war":
		numCards = 3
	case "mutually assured destruction":
		for _, deck := range []*Deck{t.Player1.Deck, t.Player2.Deck} {
			for i := 0; i < 3; i++ {
				if len(deck.Cards) > 0 {
					deck.RemoveCard()
				}
			}
		}
	}

	for i := 0; i < numCards; i++ {
		for _, deck := range []*Deck{t.Player1.Deck, t.Player2.Deck} {
			if len(deck.Cards) > 0 {
				t.SpoilsOfWar = append(t.SpoilsOfWar, deck.RemoveCard())
			}
		}
	}
}

func (t Turn) AwardSpoils(winner Player) {
	for _, card := range t.SpoilsOfWar {
		winner.Deck.AddCard(card)
	}
}
