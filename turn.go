package main

type Turn struct {
	Player1, Player2 Player
	SpoilsOfWar      []Card
}

func (t Turn) Type() string {
	if t.Player1.Deck.RankofCardAt(0) != t.Player2.Deck.RankofCardAt(0) {
		return "basic"
	}

	if t.Player1.Deck.RankofCardAt(2) != t.Player2.Deck.RankofCardAt(2) {
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

	if t.Player1.Deck.RankofCardAt(index) > t.Player2.Deck.RankofCardAt(index) {
		return t.Player1
	} else {
		return t.Player2
	}
}

func (t *Turn) PileCards() {
	var numCards int
	switch t.Type() {
	case "basic":
		numCards = 1
	case "war":
		numCards = 3
	case "mutually assured destruction":
		for i := 0; i < 3; i++ {
			t.Player1.Deck.RemoveCard()
			t.Player2.Deck.RemoveCard()
		}
	}

	for i := 0; i < numCards; i++ {
		t.SpoilsOfWar = append(t.SpoilsOfWar, t.Player1.Deck.RemoveCard())
		t.SpoilsOfWar = append(t.SpoilsOfWar, t.Player2.Deck.RemoveCard())
	}
}

func (t Turn) AwardSpoils(winner Player) {
	for _, card := range t.SpoilsOfWar {
		winner.Deck.AddCard(card)
	}
}
