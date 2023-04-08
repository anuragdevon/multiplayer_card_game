package game

import (
	"multiplayer-card-game/card"
	"multiplayer-card-game/deck"
	"testing"
)

func TestNewGame(t *testing.T) {
	names := []string{"Alice", "Bob", "Charlie"}
	game := NewGame(names)

	// Check if deck is shuffled
	if game.Deck.Cards[0] == deck.NewDeck().Cards[0] {
		t.Errorf("Deck is not shuffled")
	}

	for _, player := range game.Players {
		if len(player.Hand) != 5 {
			t.Errorf("Player does not have 5 cards in hand")
		}
	}

	// Check if top card is drawn
	if game.topCard == nil {
		t.Errorf("Top card is not drawn")
	}

	// Check if current player is first player in list
	if game.CurrentPlayer().Name != names[0] {
		t.Errorf("Current player is not first player in list")
	}
}

func TestPlayCard(t *testing.T) {
	// Create a new game with two players
	g := NewGame([]string{"Alice", "Bob"})

	// New Specific top card for testing purposes
	newCard := card.NewCard("Diamonds", "4")
	g.topCard = &newCard

	// Draw some cards for the current player
	card1 := card.NewCard("Diamonds", "9")
	card2 := card.NewCard("Spades", "5")
	card3 := card.NewCard("Hearts", "2")

	p := g.CurrentPlayer()
	p.Hand = []card.Card{
		card1, card2, card3,
	}

	// Play a valid card
	err := g.PlayCard(0)
	if err != nil {
		t.Errorf("unexpected error for playing valid card: %v", err)
	}

	// Attempt to play a card that doesn't match the top card
	err = g.PlayCard(1)
	if err != nil {
		t.Errorf("player played an invalid card")
	}

	// Attempt to play a card that isn't in the player's hand
	err = g.PlayCard(4)
	if err == nil {
		t.Errorf("expected error for playing card that isn't in player's hand")
	}
}
