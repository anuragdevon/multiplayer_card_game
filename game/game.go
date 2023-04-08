package game

import (
	"errors"
	"fmt"
	"multiplayer-card-game/card"
	"multiplayer-card-game/deck"
	"multiplayer-card-game/player"
	"os"
)

type Game struct {
	Deck          *deck.Deck
	topCard       *card.Card
	Players       []*player.Player
	CurrentIndex  int
	SkipCount     int
	ReverseOrder  bool
	PlusTwoCount  int
	PlusFourCount int
}

func NewGame(names []string) *Game {
	g := &Game{}
	g.Deck = deck.NewDeck()
	g.Deck.Shuffle()

	g.Players = make([]*player.Player, len(names))
	for i, name := range names {
		g.Players[i] = player.NewPlayer(name)
	}
	g.Deck.DrawCards(g.Players)
	// Draw a top card for new game, which will be replaced later by lastplayed card
	g.topCard, _ = g.Deck.DrawCard()
	return g
}

func (g *Game) ReturnLastCard() card.Card {
	return *g.topCard
}

func (g *Game) CurrentPlayer() *player.Player {
	return g.Players[g.CurrentIndex]
}

func (g *Game) PlayCard(index int) error {
	p := g.CurrentPlayer()
	topCard := g.ReturnLastCard()
	card, err := p.PlayCard(index, topCard)
	if err != nil {
		if len(g.Deck.Cards) == 0 {
			fmt.Print("draw pile is empty, game ends in a draw")
			os.Exit(0)
		}

		card, cards := g.Deck.Cards[len(g.Deck.Cards)-1], g.Deck.Cards[:len(g.Deck.Cards)-1]
		g.Deck.Cards = cards
		p.Hand = append(p.Hand, card)
	}

	p.LastPlayed = card // Store the last played card in the player's LastPlayed field
	g.UpdateGameStatus(card)
	g.UpdatePlayerOrder(card)
	// g.Deck.Cards = append(g.Deck.Cards, p.LastPlayed)
	g.topCard = &p.LastPlayed
	return err
}

func (g *Game) UpdateGameStatus(card card.Card) {
	switch card.Rank {
	case "Ace":
		g.SkipCount++
	case "King":
		g.ReverseOrder = !g.ReverseOrder
	case "Queen":
		g.PlusTwoCount++
	case "Jack":
		g.PlusFourCount++
	}
}

func (g *Game) UpdatePlayerOrder(card card.Card) {
	if g.ReverseOrder {
		g.CurrentIndex--
		if g.CurrentIndex < 0 {
			g.CurrentIndex = len(g.Players) - 1
		}
	} else {
		g.CurrentIndex++
		if g.CurrentIndex >= len(g.Players) {
			g.CurrentIndex = 0
		}
	}

	if g.SkipCount > 0 {
		g.SkipCount--
		g.UpdatePlayerOrder(card)
	}

	if g.PlusTwoCount > 0 {
		g.PlusTwoCount--
		g.DrawCards(2)
	}

	if g.PlusFourCount > 0 {
		g.PlusFourCount--
		g.DrawCards(4)
	}
}

func (g *Game) TopCard() card.Card {
	return g.Deck.Cards[len(g.Deck.Cards)-1]
}

func (g *Game) DrawCards(count int) {
	p := g.CurrentPlayer()
	for i := 0; i < count; i++ {
		if len(g.Deck.Cards) == 0 {
			return
		}

		card, cards := g.Deck.Cards[len(g.Deck.Cards)-1], g.Deck.Cards[:len(g.Deck.Cards)-1]
		g.Deck.Cards = cards
		p.Hand = append(p.Hand, card)
	}
}

func (g *Game) IsGameOver() bool {
	for _, player := range g.Players {
		if len(player.Hand) == 0 {
			return true
		}
	}
	return false
}

func (g *Game) Winner() (*player.Player, error) {
	if !g.IsGameOver() {
		return nil, errors.New("Game is not over yet")
	}

	winner := g.Players[0]
	for _, player := range g.Players {
		if len(player.Hand) < len(winner.Hand) {
			winner = player
		}
	}
	return winner, nil
}

func (g *Game) String() string {
	var str string
	for _, player := range g.Players {
		str += fmt.Sprintf("%s: %s\n", player, player.Hand)
	}
	str += fmt.Sprintf("Top card: %s\n", g.TopCard())
	str += fmt.Sprintf("Current player: %s\n", g.CurrentPlayer().Name)
	return str
}
