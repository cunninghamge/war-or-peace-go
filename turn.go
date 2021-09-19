package main

type Turn struct {
	Player1, Player2 *Player
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

	for i := 3; i > 0; i-- {
		if t.Player1.Deck.RankofCardAt(i) != t.Player2.Deck.RankofCardAt(i) {
			return war
		}
	}

	return mutuallyAssuredDestruction
}

func (t Turn) Winner() *Player {
	for _, i := range []int{0, 3, 2, 1} {
		player1Card := t.Player1.Deck.RankofCardAt(i)
		player2Card := t.Player2.Deck.RankofCardAt(i)
		if player1Card > player2Card {
			return t.Player1
		}
		if player2Card > player1Card {
			return t.Player2
		}
	}
	return nil
}

func (t *Turn) PileCards() {
	turnType := t.Type()
	for i := 0; i < 4; i++ {
		for _, deck := range []*Deck{t.Player1.Deck, t.Player2.Deck} {
			if len(deck.Cards) > 0 {
				t.SpoilsOfWar = append(t.SpoilsOfWar, deck.RemoveCard())
			}
		}
		if turnType == basic {
			break
		}
	}
}

func (t *Turn) AwardSpoils(winner *Player) {
	if winner != nil {
		winner.Deck.AddCards(t.SpoilsOfWar)
	}
}
