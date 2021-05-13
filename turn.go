package main

type Turn struct {
	Player1, Player2 Player
	SpoilsOfWar      []Card
}

func (t Turn) Type() string {

	return "basic"
}

func (t Turn) Winner() Player {
	if t.Player1.Deck.RankofCardAt(0) > t.Player2.Deck.RankofCardAt(0) {
		return t.Player1
	}
	return t.Player2
}
