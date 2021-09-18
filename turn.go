package main

type Turn struct {
	Player1, Player2 Player
	SpoilsOfWar      []Card
}

type TurnType int

const (
	basic TurnType = iota
	war
	mutuallyAssuredDestruction
)

func (t Turn) Type() TurnType {
	if t.Player1.Deck.RankofCardAt(0) != t.Player2.Deck.RankofCardAt(0) {
		return basic
	}

	if t.Player1.Deck.RankofCardAt(2) != t.Player2.Deck.RankofCardAt(2) {
		return war
	}

	return mutuallyAssuredDestruction
}

func (t Turn) Winner() *Player {
	var index int
	switch t.Type() {
	case basic:
		index = 0
	case war:
		index = 2
	case mutuallyAssuredDestruction:
		return nil
	}

	if t.Player1.Deck.RankofCardAt(index) > t.Player2.Deck.RankofCardAt(index) {
		return &t.Player1
	} else {
		return &t.Player2
	}
}

func (t *Turn) PileCards() {
	turnType := t.Type()
	for i := 0; i < 3; i++ {
		if turnType == basic && i > 0 {
			return
		}
		for _, deck := range []*Deck{t.Player1.Deck, t.Player2.Deck} {
			if len(deck.Cards) > 0 {
				t.SpoilsOfWar = append(t.SpoilsOfWar, deck.RemoveCard())
			}
		}
	}
}

func (t *Turn) AwardSpoils(winner *Player) {
	if winner != nil {
		for _, card := range t.SpoilsOfWar {
			winner.Deck.AddCard(card)
		}
	}
}
