package game

import (
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

	// Check if all players are created and have 7 cards
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