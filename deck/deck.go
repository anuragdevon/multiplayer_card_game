package main

import (
	"math/rand"
	"multiplayer-card-game/card"
	"time"
)

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

func (d *Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := len(d.cards) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// TODO: Method for adding a card to players collection
