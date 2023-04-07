package game

import (
	"fmt"
	"math/rand"
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
	currentPlayerIndex := rand.Intn(len(g.Players))

	// Setting initial top card
	// topCard, g.Deck = g.Deck.Cards[len(g.Deck.Cards)-1], g.Deck.Cards[:len(g.Deck.Cards)-1]

	// main decision control flow for selection of cards
	for {
		currentPlayer := g.Players[currentPLayerIndex]

		// Display top card and current player's hand
		fmt.Println("Top card: ", topCard)
		fmt.Println("Current Player: ", currentPlayer.Name)
		fmt.Println("Your hand: ")
		fmt.Println(currentPlayer)

		// Check if player has valid move
		if currentPlayer.HasValidMove(topCard) {
			// Prompt Player to play the card
			index := promptForIndex(currentPlayer)
			card, err := currentPlayer.PlayCard(index, topCard)
			if err != nil {
				fmt.Println(err)
				continue
			}

			// Case for action cards and apply effects
			switch card.Rank {
			case "Ace":
				increment = 1
			case "King":
				turnDirection *= -1
			case "Queen":
				increment = 2
			case "Jack":
				increment = 4
			default:
				increment = 1
			}

			// Direction check
			if turnDirection == 1 {
				currentPlayerIndex = (currentPlayerIndex + increment) % len(g.Players)
			} else {
				currentPlayerIndex = (currentPlayerIndex - increment + len(g.Players)) % len(g.Players)
			}

			// Set the top card as player card
			topCard = card

			// Check if player has won
			if len(currentPlayer.Hand) == 0 {
				fmt.Println("Game over! Winner:", currentPlayer.Name)
				return
			}
		} else {
			// Player has no valid move, draw a card
			if len(g.Deck.Cards) == 0 {
				fmt.Println("Draw pile is empty! Game Over.")
				return
			}

			card, cards := g.Deck.Cards[len(g.Deck.Cards)-1], g.Deck.Cards[:len(g.Deck.Cards)-1]
			g.Deck.Cards = cards
			currentPlayer.Hand = append(currentPlayer.Hand, card)

			// Skip next player of ace is drawn
			if card.Rank == "Ace" {
				if turnDirection == 1 {
					currentPlayerIndex = (currentPlayerIndex + 1) % len(g.Players)
				} else {
					currentPlayerIndex = (currentPlayerIndex - 1 + len(g.Players)) % len(g.Players)
				}
			}

			// Switch turn direction if King is drawn
			if card.Rank == "King" {
				turnDirection *= -1
			}

			// Switch to next plater in forward direction
			if turnDirection == 1 {
				currentPlayerIndex = (currentPlayerIndex + 1) % len(g.Players)
			} else {
				// Switch to next player in backward direction
				currentPlayerIndex = (currentPlayerIndex - 1 + len(g.Players)) % len(g.Players)
			}
		}
	}
	// checking the gamestatus
}

func promptForIndex(p *player.Player) int {
	for {
		fmt.Println("Enter the index of the card to play: ")
		var index int
		_, err := fmt.Scanln(&index)
		if err != nil || index < 0 || index >= len(p.Hand) {
			fmt.Println("Invalid input! Please enter a valid index. ")
			continue
		}
		return index
	}
}
