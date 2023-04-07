package player

import (
	"errors"
	"fmt"
	"multiplayer-card-game/card"
)

type Player struct {
	Name       string
	Hand       []card.Card
	LastPlayed card.Card // New field added to store the last card played by the player

}

func NewPlayer(name string) *Player {
	p := &Player{}
	p.Name = name
	p.Hand = make([]card.Card, 0)
	return p
}

func (p *Player) String() string {
	var str string
	for _, card := range p.Hand {
		str += fmt.Sprintf("%s\n", card)
	}
	return str
}

// TODO: Handle addtion checks for playcard and invalid moves
func (p *Player) PlayCard(index int, topCard card.Card) (card.Card, error) {
	card := p.Hand[index]
	if card.Suit != topCard.Suit && card.Rank != topCard.Rank {
		return card, errors.New("Invalid Move") // TODO: change card to empty ~ 1
	}

	p.Hand = append(p.Hand[:index], p.Hand[index+1:]...)
	return card, nil
}

func (p *Player) HasValidMove(topCard card.Card) bool {
	for _, card := range p.Hand {
		if card.Suit == topCard.Suit || card.Rank == topCard.Rank {
			return true
		}
	}
	return false
}
