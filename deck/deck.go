package deck

import (
	"fmt"
	"math/rand"
	"multiplayer-card-game/card"
	"multiplayer-card-game/player"
	"time"
)

type Deck struct {
	Cards []card.Card
}

func NewDeck() *Deck {
	d := &Deck{}
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "King", "Queen"}

	for _, suit := range suits {
		for _, rank := range ranks {
			card := card.NewCard(suit, rank)
			d.Cards = append(d.Cards, card)
		}
	}
	return d
}

func (d *Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := len(d.Cards) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *Deck) DrawCards(players []*player.Player) {
	for i := 0; i < 5; i++ {
		for _, player := range players {
			card, cards := d.Cards[len(d.Cards)-1], d.Cards[:len(d.Cards)-1]
			d.Cards = cards
			player.Hand = append(player.Hand, card)
		}
	}
}

func (d *Deck) DrawCard() (*card.Card, error) {
	if len(d.Cards) == 0 {
		return nil, fmt.Errorf("deck is empty")
	}

	card, cards := d.Cards[len(d.Cards)-1], d.Cards[:len(d.Cards)-1]
	d.Cards = cards

	return &card, nil
}
