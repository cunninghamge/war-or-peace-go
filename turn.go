package main

type Turn struct {
	player1, player2 *Player
	spoilsOfWar      []Card
}

type TurnType int

const (
	basic TurnType = iota
	war
	mutuallyAssuredDestruction
)

func (t Turn) Type() TurnType {
	if t.player1.deck.RankofCardAt(0) != t.player2.deck.RankofCardAt(0) {
		return basic
	}

	for i := 3; i > 0; i-- {
		if t.player1.deck.RankofCardAt(i) != t.player2.deck.RankofCardAt(i) {
			return war
		}
	}

	return mutuallyAssuredDestruction
}

func (t Turn) Winner() *Player {
	for _, i := range []int{0, 3, 2, 1} {
		player1Card := t.player1.deck.RankofCardAt(i)
		player2Card := t.player2.deck.RankofCardAt(i)
		if player1Card > player2Card {
			return t.player1
		}
		if player2Card > player1Card {
			return t.player2
		}
	}
	return nil
}

func (t *Turn) PileCards() {
	turnType := t.Type()
	for i := 0; i < 4; i++ {
		for _, deck := range []*Deck{t.player1.deck, t.player2.deck} {
			if len(deck.cards) > 0 {
				t.spoilsOfWar = append(t.spoilsOfWar, deck.RemoveCard())
			}
		}
		if turnType == basic {
			break
		}
	}
}

func (t *Turn) AwardSpoils(winner *Player) {
	if winner != nil {
		winner.deck.AddCards(t.spoilsOfWar)
	}
}
