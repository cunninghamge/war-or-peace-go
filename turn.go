package main

type Turn struct {
	Player1, Player2 Player
	SpoilsOfWar      []Card
}

func (t Turn) Type() string {
	if t.Player1.Deck.RankofCardAt(0) == t.Player2.Deck.RankofCardAt(0) {
		return "war"
	}
	return "basic"
}

func (t Turn) Winner() Player {
	switch t.Type() {
	case "basic":
		if t.Player1.Deck.RankofCardAt(0) > t.Player2.Deck.RankofCardAt(0) {
			return t.Player1
		} else {
			return t.Player2
		}
	case "war":
		if t.Player1.Deck.RankofCardAt(2) > t.Player2.Deck.RankofCardAt(2) {
			return t.Player1
		} else {
			return t.Player2
		}
	}

	return Player{Name: "No Winner"}
}

func (t *Turn) PileCards() {
	var numCards int
	switch t.Type() {
	case "basic":
		numCards = 0
	case "war":
		numCards = 2
	}

	for i := 0; i <= numCards; i++ {
		t.SpoilsOfWar = append(t.SpoilsOfWar, t.Player1.Deck.RemoveCard())
		t.SpoilsOfWar = append(t.SpoilsOfWar, t.Player2.Deck.RemoveCard())
	}
}

func (t Turn) AwardSpoils(winner Player) {
	for _, card := range t.SpoilsOfWar {
		winner.Deck.AddCard(card)
	}
}
