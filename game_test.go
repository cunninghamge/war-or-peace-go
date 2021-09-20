package main

import (
	"bufio"
	"bytes"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestNewGame(t *testing.T) {
	reader := &bytes.Buffer{}
	writer := &bytes.Buffer{}
	reader.Write([]byte("Player1\nPlayer2\n"))

	game := newGame("fixtures/two_cards.csv", reader, writer)

	got := writer.String()
	want := `Welcome to War!
This game will be played with 2 cards.
Enter player 1: Enter player 2: The players today are Player1 and Player2.
`
	if got != want {
		t.Errorf("got\n%s\n  want\n%s\n", got, want)
	}
	card1, card2 := Card{"2", "diamond", 2}, Card{"8", "spade", 8}
	p1Card, p2Card := game.player1.deck.cards[0], game.player2.deck.cards[0]
	if !reflect.DeepEqual(p1Card, card1) && !reflect.DeepEqual(p1Card, card2) {
		t.Errorf("got %v want %v or %v", p1Card, card1, card2)
	}
	if !reflect.DeepEqual(p2Card, card1) && !reflect.DeepEqual(p2Card, card2) {
		t.Errorf("got %v want %v or %v", p2Card, card1, card2)
	}
}

func TestGetPlayers(t *testing.T) {
	reader := &bytes.Buffer{}
	reader.Write([]byte("Player1\nPlayer2\n"))
	writer := &bytes.Buffer{}
	game := Game{
		scanner: bufio.NewScanner(reader),
		writer:  writer,
		player1: &Player{deck: &Deck{[]Card{}}},
		player2: &Player{deck: &Deck{[]Card{}}},
	}

	game.getPlayers()

	got := writer.String()
	want := `Enter player 1: Enter player 2: The players today are Player1 and Player2.
`
	if got != want {
		t.Errorf("got\n%s\n  want\n%s\n", got, want)
	}
	if game.player1.name != "Player1" {
		t.Errorf("got %s want %s", game.player1.name, "Player1")
	}
	if game.player2.name != "Player2" {
		t.Errorf("got %s want %s", game.player2.name, "Player2")
	}
}

func TestDeal(t *testing.T) {
	game := Game{
		player1: &Player{deck: &Deck{[]Card{}}},
		player2: &Player{deck: &Deck{[]Card{}}},
	}
	deck := &Deck{testCards}

	game.deal(deck)

	if game.player1.CardsLeft() != 2 {
		t.Errorf("player should have 2 cards but has %d", game.player1.CardsLeft())
	}
	if game.player2.CardsLeft() != 2 {
		t.Errorf("player should have 2 cards but has %d", game.player2.CardsLeft())
	}
}

func TestStart(t *testing.T) {
	testCases := map[string]string{
		"GO\n":   startMessage + "Turn 1: A has 2 cards left and B has 2 cards left\n",
		"go\n":   startMessage + "Turn 1: A has 2 cards left and B has 2 cards left\n",
		"q\n":    startMessage,
		"f\nq\n": startMessage + startMessage,
	}

	for cmd, want := range testCases {
		t.Run(cmd, func(t *testing.T) {
			reader := &bytes.Buffer{}
			reader.Write([]byte(cmd))
			writer := &bytes.Buffer{}
			game := Game{
				scanner: bufio.NewScanner(reader),
				writer:  writer,
				player1: &Player{name: "A", deck: &Deck{testCards[:2]}},
				player2: &Player{name: "B", deck: &Deck{testCards[2:]}},
			}

			game.start()

			got := writer.String()
			if !strings.HasPrefix(got, want) {
				t.Errorf("got %s want %s", got, want)
			}
		})
	}
}

func TestPlay(t *testing.T) {
	t.Run("game with basic and war turns", func(t *testing.T) {
		writer := &bytes.Buffer{}
		game := Game{
			scanner: bufio.NewScanner(&bytes.Buffer{}),
			writer:  writer,
			player1: &Player{
				name: "Player1",
				deck: &Deck{[]Card{
					{"2", "spade", 2},
					{"9", "club", 9},
					{"Ace", "heart", 14},
					{"3", "diamond", 3},
					{"4", "club", 4},
				}},
			},
			player2: &Player{
				name: "Player2",
				deck: &Deck{[]Card{
					{"3", "club", 3},
					{"9", "spade", 9},
					{"Ace", "spade", 14},
					{"3", "heart", 3},
					{"5", "diamond", 5},
				}},
			},
		}

		game.play()

		got := writer.String()
		want := `Turn 1: Player1 has 5 cards left and Player2 has 5 cards left
		Player1 played the 2 of spades and Player2 played the 3 of clubs
		Player2 won 2 cards

Turn 2: Player1 has 4 cards left and Player2 has 6 cards left
		Player1 played the 9 of clubs and Player2 played the 9 of spades
	WAR!	Player1 played the 4 of clubs and Player2 played the 5 of diamonds
		Player2 won 8 cards

Player1 is out of cards!
*~*~*~* Player2 won the game! *~*~*~*
`
		if got != want {
			t.Errorf("got\n%s\nwant\n%s", got, want)
		}
	})

	t.Run("shuffles after 500 turns", func(t *testing.T) {
		writer := &bytes.Buffer{}
		game := Game{
			scanner: bufio.NewScanner(&bytes.Buffer{}),
			writer:  writer,
			player1: &Player{
				name: "Player1",
				deck: &Deck{[]Card{
					{"2", "spade", 2},
					{"4", "club", 4},
				}},
			},
			player2: &Player{
				name: "Player2",
				deck: &Deck{[]Card{
					{"5", "diamond", 5},
					{"3", "club", 3},
				}},
			},
		}

		game.play()

		s := writer.String()
		begin := strings.LastIndex(s, "Turn")
		end := strings.LastIndex(s, ":")
		lastTurn, _ := strconv.Atoi(s[begin+5 : end])
		if (lastTurn-2)%2 != 0 {
			t.Errorf("expected game to end after 500n + 2 turns but ended in %d", lastTurn)
		}
	})
}

func TestShowCards(t *testing.T) {
	writer := &bytes.Buffer{}
	game := Game{
		scanner: bufio.NewScanner(&bytes.Buffer{}),
		writer:  writer,
		player1: &Player{
			name: "Player1",
			deck: &Deck{[]Card{
				{"Queen", "diamond", 12},
			}},
		},
		player2: &Player{
			name: "Player2",
			deck: &Deck{[]Card{
				{"10", "heart", 10},
			}},
		},
	}

	game.showCards(0, "	prefix	")

	got := writer.String()
	want := "	prefix	Player1 played the Queen of diamonds and Player2 played the 10 of hearts\n"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestCanPlayWar(t *testing.T) {
	testCases := map[string]struct {
		player1, player2 *Player
		result           bool
		want             string
	}{
		"both players have enough cards": {
			player1: &Player{
				name: "Player1",
				deck: &Deck{[]Card{
					{"Queen", "diamond", 12},
					{"Queen", "diamond", 12},
					{"Queen", "diamond", 12},
					{"Queen", "diamond", 12},
				}},
			},
			player2: &Player{
				name: "Player2",
				deck: &Deck{[]Card{
					{"10", "heart", 10},
					{"10", "heart", 10},
					{"10", "heart", 10},
					{"10", "heart", 10},
				}},
			},
			result: true,
		},
		"one players has enough cards": {
			player1: &Player{
				name: "Player1",
				deck: &Deck{[]Card{
					{"Queen", "diamond", 12},
					{"Queen", "diamond", 12},
					{"Queen", "diamond", 12},
					{"Queen", "diamond", 12},
				}},
			},
			player2: &Player{
				name: "Player2",
				deck: &Deck{[]Card{
					{"10", "heart", 10},
				}},
			},
			result: false,
			want: "		Player2 does not have enough cards for war!\n",
		},
		"neither player has enough cards": {
			player1: &Player{
				name: "Player1",
				deck: &Deck{[]Card{
					{"Queen", "diamond", 12},
				}},
			},
			player2: &Player{
				name: "Player2",
				deck: &Deck{[]Card{
					{"10", "heart", 10},
				}},
			},
			result: false,
			want: "		Player1 does not have enough cards for war!\n" +
				"		Player2 does not have enough cards for war!\n",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			game := Game{
				writer:  writer,
				player1: tc.player1,
				player2: tc.player2,
			}

			canPlayWar := game.canPlayWar()

			if canPlayWar != tc.result {
				t.Errorf("got %t want %t", canPlayWar, tc.result)
			}
			got := writer.String()
			if got != tc.want {
				t.Errorf("got %s want %s", got, tc.want)
			}
		})
	}
}

func TestWar(t *testing.T) {
	testCases := map[string]struct {
		player1Cards []Card
		player2Cards []Card
		want         string
	}{
		"war: first card wins": {
			player1Cards: []Card{
				{"Queen", "diamond", 12},
				{"2", "club", 2},
				{"7", "spade", 7},
				{"4", "heart", 4},
			},
			player2Cards: []Card{
				{"Queen", "heart", 12},
				{"10", "diamond", 10},
				{"5", "spade", 5},
				{"3", "club", 3},
			},
			want: "	WAR!	Player1 played the 4 of hearts and Player2 played the 3 of clubs\n",
		},
		"war: second card wins": {
			player1Cards: []Card{
				{"Queen", "diamond", 12},
				{"2", "club", 2},
				{"7", "spade", 7},
				{"3", "heart", 3},
			},
			player2Cards: []Card{
				{"Queen", "heart", 12},
				{"10", "diamond", 10},
				{"5", "spade", 5},
				{"3", "club", 3},
			},
			want: "	WAR!	Player1 played the 3 of hearts and Player2 played the 3 of clubs\n" +
				"	WAR!	Player1 played the 7 of spades and Player2 played the 5 of spades\n",
		},
		"war: last card wins": {
			player1Cards: []Card{
				{"Queen", "diamond", 12},
				{"10", "diamond", 10},
				{"7", "spade", 7},
				{"3", "heart", 3},
			},
			player2Cards: []Card{
				{"Queen", "heart", 12},
				{"2", "club", 2},
				{"7", "club", 7},
				{"3", "club", 3},
			},
			want: "	WAR!	Player1 played the 3 of hearts and Player2 played the 3 of clubs\n" +
				"	WAR!	Player1 played the 7 of spades and Player2 played the 7 of clubs\n" +
				"	WAR!	Player1 played the 10 of diamonds and Player2 played the 2 of clubs\n",
		},
		"mutually assured destruction": {
			player1Cards: []Card{
				{"Queen", "diamond", 12},
				{"10", "diamond", 10},
				{"7", "spade", 7},
				{"3", "heart", 3},
			},
			player2Cards: []Card{
				{"Queen", "heart", 12},
				{"10", "club", 10},
				{"7", "club", 7},
				{"3", "club", 3},
			},
			want: "	WAR!	Player1 played the 3 of hearts and Player2 played the 3 of clubs\n" +
				"	WAR!	Player1 played the 7 of spades and Player2 played the 7 of clubs\n" +
				"	WAR!	Player1 played the 10 of diamonds and Player2 played the 10 of clubs\n",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			game := Game{
				scanner: bufio.NewScanner(&bytes.Buffer{}),
				writer:  writer,
				player1: &Player{
					name: "Player1",
					deck: &Deck{tc.player1Cards},
				},
				player2: &Player{
					name: "Player2",
					deck: &Deck{tc.player2Cards},
				},
			}

			game.war()

			got := writer.String()
			want := tc.want
			if got != want {
				t.Errorf("got\n%s\nwant\n%s", got, want)
			}
		})
	}
}

func TestAwardSpoils(t *testing.T) {
	testCases := map[string]struct {
		player1Cards, player2Cards         []Card
		player1CardsLeft, player2CardsLeft int
		want                               string
	}{
		"basic turn": {
			player1Cards: []Card{
				{"Queen", "diamond", 12},
			},
			player2Cards: []Card{
				{"10", "diamond", 10},
			},
			want: "		Player1 won 2 cards\n\n",
			player1CardsLeft: 2,
			player2CardsLeft: 0,
		},
		"war turn": {
			player1Cards: []Card{
				{"Queen", "diamond", 12},
				{"2", "club", 2},
				{"7", "spade", 7},
				{"3", "club", 3},
			},
			player2Cards: []Card{
				{"Queen", "heart", 12},
				{"10", "diamond", 10},
				{"5", "spade", 5},
				{"4", "heart", 4},
			},
			want: "		Player2 won 8 cards\n\n",
			player1CardsLeft: 0,
			player2CardsLeft: 8,
		},
		"mutually assured destruction": {
			player1Cards: []Card{
				{"Queen", "diamond", 12},
				{"10", "diamond", 10},
				{"7", "spade", 7},
				{"3", "heart", 3},
			},
			player2Cards: []Card{
				{"Queen", "heart", 12},
				{"10", "club", 10},
				{"7", "club", 7},
				{"3", "club", 3},
			},
			want: "		*mutually assured destruction* 8 cards removed from play\n\n",
			player1CardsLeft: 0,
			player2CardsLeft: 0,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			game := Game{
				scanner: bufio.NewScanner(&bytes.Buffer{}),
				writer:  writer,
				player1: &Player{
					name: "Player1",
					deck: &Deck{tc.player1Cards},
				},
				player2: &Player{
					name: "Player2",
					deck: &Deck{tc.player2Cards},
				},
			}
			turn := Turn{player1: game.player1, player2: game.player2}

			game.awardSpoils(turn)

			got := writer.String()
			want := tc.want
			if got != want {
				t.Errorf("got\n%s\nwant\n%s", got, want)
			}
			if game.player1.CardsLeft() != tc.player1CardsLeft {
				t.Errorf("got %d want %d", game.player1.CardsLeft(), tc.player1CardsLeft)
			}
			if game.player2.CardsLeft() != tc.player2CardsLeft {
				t.Errorf("got %d want %d", game.player2.CardsLeft(), tc.player2CardsLeft)
			}
		})
	}
}

func TestDisplayResult(t *testing.T) {
	testCases := map[string]struct {
		player1, player2 *Player
		want             string
	}{
		"one player is out of cards": {
			player1: &Player{
				name: "Player1",
				deck: &Deck{[]Card{}},
			},
			player2: &Player{
				name: "Player2",
				deck: &Deck{[]Card{
					{"10", "diamond", 10},
				}},
			},
			want: "Player1 is out of cards!\n" +
				"*~*~*~* Player2 won the game! *~*~*~*\n",
		},
		"one player has cards but has lost": {
			player1: &Player{
				name: "Player1",
				deck: &Deck{[]Card{
					{"Queen", "diamond", 12},
				}},
				lost: true,
			},
			player2: &Player{
				name: "Player2",
				deck: &Deck{[]Card{
					{"10", "diamond", 10},
				}},
			},
			want: "*~*~*~* Player2 won the game! *~*~*~*\n",
		},
		"both players are out of cards": {
			player1: &Player{
				name: "Player1",
				deck: &Deck{[]Card{}},
			},
			player2: &Player{
				name: "Player2",
				deck: &Deck{[]Card{}},
			},
			want: "Player1 is out of cards!\n" +
				"Player2 is out of cards!\n" +
				"*~*~*~* It's a draw! *~*~*~*\n",
		},
		"both players have cards but have lost": {
			player1: &Player{
				name: "Player1",
				deck: &Deck{[]Card{
					{"Queen", "diamond", 12},
				}},
				lost: true,
			},
			player2: &Player{
				name: "Player2",
				deck: &Deck{[]Card{
					{"10", "diamond", 10},
				}},
				lost: true,
			},
			want: "*~*~*~* It's a draw! *~*~*~*\n",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			game := Game{
				scanner: bufio.NewScanner(&bytes.Buffer{}),
				writer:  writer,
				player1: tc.player1,
				player2: tc.player2,
			}

			game.displayResult()

			got := writer.String()
			if got != tc.want {
				t.Errorf("got\n%s\nwant\n%s", got, tc.want)
			}
		})
	}
}
