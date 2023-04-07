package player

import (
	"fmt"
	"multiplayer-card-game/card"
)

type Player struct {
	name string
	Hand []card.Card
}

func NewPlayer(name string) *Player {
	p := &Player{}
	p.name = name
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
