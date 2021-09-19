package main

import (
	"errors"
	"fmt"
	"strconv"
)

const errInvalidRecords = "invalid record: must contain exactly 3 fields"

type Card struct {
	Value string
	Suit  string
	Rank  int
}

func (c Card) String() string {
	return fmt.Sprintf("the %s of %ss", c.Value, c.Suit)
}

func createCards(records [][]string) ([]Card, error) {
	var cards []Card
	for _, record := range records {
		if len(record) != 3 {
			return nil, errors.New(errInvalidRecords)
		}

		rank, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}
		cards = append(cards, Card{record[0], record[1], rank})
	}
	return cards, nil
}
