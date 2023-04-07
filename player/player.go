package player

import (
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
