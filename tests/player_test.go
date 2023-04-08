package test

import (
	"multiplayer-card-game/card"
	"multiplayer-card-game/player"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	p := player.NewPlayer("Alice")
	if p.Name != "Alice" {
		t.Errorf("Expected name to be 'Alice', but got '%s'", p.Name)
	}

	if len(p.Hand) != 0 {
		t.Errorf("Expected empty hand, but got %v", p.Hand)
	}
}

func TestPlayerString(t *testing.T) {
	p := player.NewPlayer("Bob")
	p.Hand = append(p.Hand, card.NewCard("Hearts", "Ace"))
	p.Hand = append(p.Hand, card.NewCard("Diamonds", "King"))

	expectedStr := "Ace of Hearts\nKing of Diamonds\n"
	if p.String() != expectedStr {
		t.Errorf("Expected string representation to be '%s', but got '%s'", expectedStr, p.String())
	}
}

func TestPlayCard(t *testing.T) {
	p := player.NewPlayer("Charlie")
	p.Hand = append(p.Hand, card.NewCard("Spades", "10"))
	p.Hand = append(p.Hand, card.NewCard("Hearts", "Jack"))

	// valid move
	card_, err := p.PlayCard(1, card.NewCard("Hearts", "9"))
	if err != nil {
		t.Errorf("Expected no error, but got '%v'", err)
	}

	if card_.Suit != "Hearts" || card_.Rank != "Jack" {
		t.Errorf("Expected played card to be Jack of Hearts, but got %v", card_)
	}

	if len(p.Hand) != 1 || p.Hand[0].Suit != "Spades" || p.Hand[0].Rank != "10" {
		t.Errorf("Expected hand to have one card, but got %v", p.Hand)
	}

	// invalid move
	_, err = p.PlayCard(0, card.NewCard("Hearts", "9"))
	if err == nil {
		t.Errorf("Expected error, but got none")
	}
}
