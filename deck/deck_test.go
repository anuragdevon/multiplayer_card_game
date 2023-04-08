package deck

import (
	"multiplayer-card-game/card"
	"multiplayer-card-game/player"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	expectedCardCount := 52
	if len(d.Cards) != expectedCardCount {
		t.Errorf("Expected %d cards, but got %d", expectedCardCount, len(d.Cards))
	}
}

func TestShuffle(t *testing.T) {
	d := NewDeck()
	originalCards := make([]card.Card, len(d.Cards))
	copy(originalCards, d.Cards)

	d.Shuffle()

	if len(d.Cards) != len(originalCards) {
		t.Errorf("Expected %d cards, but got %d", len(originalCards), len(d.Cards))
	}

	if d.Cards[0] == originalCards[0] && d.Cards[len(d.Cards)-1] == originalCards[len(originalCards)-1] {
		t.Errorf("Deck did not shuffle, first and last cards are the same")
	}
}

func TestDrawCards(t *testing.T) {
	d := NewDeck()
	p1 := player.NewPlayer("Player 1")
	p2 := player.NewPlayer("Player 2")
	players := []*player.Player{p1, p2}

	expectedCardCount := len(players) * 5
	d.DrawCards(players)

	if len(d.Cards) != 42 {
		t.Errorf("Expected %d cards left in the deck, but got %d", 42, len(d.Cards))
	}

	if len(p1.Hand) != 5 || len(p2.Hand) != 5 {
		t.Errorf("Expected each player to have 5 cards, but player 1 has %d and player 2 has %d", len(p1.Hand), len(p2.Hand))
	}

	actualCardCount := len(p1.Hand) + len(p2.Hand)
	if actualCardCount != expectedCardCount {
		t.Errorf("Expected %d cards to be drawn, but got %d", expectedCardCount, actualCardCount)
	}
}

func TestDrawCard(t *testing.T) {
	d := NewDeck()

	// Draw 52 cards from the deck
	for i := 0; i < 52; i++ {
		card, err := d.DrawCard()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(d.Cards) != 52-i-1 {
			t.Errorf("Expected %d cards left in the deck, but got %d", 52-i-1, len(d.Cards))
		}
		if card == nil {
			t.Errorf("Expected a card, but got nil")
		}
	}

	// Draw one more card from the deck, should return an error
	card, err := d.DrawCard()
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	if card != nil {
		t.Errorf("Expected nil card, but got %v", card)
	}
}
