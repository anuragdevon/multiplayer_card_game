package main

import "multiplayer-card-game/card"

type Deck struct {
	cards []card.Card
}

func NewDeck() *Deck {
	d := &Deck{}
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "King", "Queen"}

	for _, suit := range suits {
		for _, rank := range ranks {
			card := card.NewCard(suit, rank)
			d.cards = append(d.cards, card)
		}
	}
	return d
}
