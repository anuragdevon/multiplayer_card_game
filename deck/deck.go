package deck

import (
	"fmt"
	"math/rand"
	"multiplayer-card-game/card"
	"multiplayer-card-game/player"
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

func (d *Deck) DrawCards(players []*player.Player) {
	for i := 0; i < 5; i++ {
		for _, player := range players {
			card, cards := d.cards[len(d.cards)-1], d.cards[:len(d.cards)-1]
			d.cards = cards
			player.Hand = append(player.Hand, card)
		}
	}
}

// temporary method for showing deck cards
func (d *Deck) String() string {
	var str string
	for _, card := range d.cards {
		str += fmt.Sprintf("%s\n", card)
	}
	return str
}
