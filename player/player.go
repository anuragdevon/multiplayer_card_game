package player

import (
	"errors"
	"multiplayer-card-game/card"
)

type Player struct {
	Name       string
	Hand       []card.Card
	LastPlayed card.Card
}

func NewPlayer(name string) *Player {
	p := &Player{}
	p.Name = name
	p.Hand = make([]card.Card, 0)
	return p
}

func (p *Player) PlayCard(index int, topCard card.Card) (card.Card, error) {
	if index < 0 || index >= len(p.Hand) {
		return card.Card{}, errors.New("invalid index")
	}
	cards := p.Hand[index]
	if cards.Suit != topCard.Suit && cards.Rank != topCard.Rank {
		return card.Card{}, errors.New("not a valid move, has to pick a card")
	}

	p.Hand = append(p.Hand[:index], p.Hand[index+1:]...)
	return cards, nil
}
