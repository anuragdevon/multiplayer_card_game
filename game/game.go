package game

import (
	"multiplayer-card-game/card"
	"multiplayer-card-game/deck"
	"multiplayer-card-game/player"
)

type Game struct {
	Players []*player.Player
	Deck    *deck.Deck
}

func NewGame(playerNames []string) *Game {
	g := &Game{}
	players := make([]*player.Player, len(playerNames))
	for i, name := range playerNames {
		players[i] = player.NewPlayer(name)
	}
	g.Players = players
	g.Deck = deck.NewDeck()
	g.Deck.Shuffle()
	g.Deck.DrawCards(g.Players)
	return g
}

func (g *Game) Start() {
	var topCard card.Card
	var currentPLayerIndex int
	var turnDirection int // 1: Forward. -1: backwards
	var increment int     // skip and draw

	// Randomly select first player
	// Setting initial top card
	// main decision control flow for selection of cards
	// checking the gamestatus
}
