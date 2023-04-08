package card

import (
	// "multiplayer-card-game/card"
	"testing"
)

func TestNewCard(t *testing.T) {
	suit := "Spades"
	rank := "Ace"
	c := NewCard(suit, rank)
	if c.Suit != suit {
		t.Errorf("Expected suit to be %s, but got %s", suit, c.Suit)
	}
	if c.Rank != rank {
		t.Errorf("Expected rank to be %s, but got %s", rank, c.Rank)
	}
}

func TestCardString(t *testing.T) {
	c := Card{Suit: "Diamonds", Rank: "10"}
	expected := "10 of Diamonds"
	if c.String() != expected {
		t.Errorf("Expected card string to be %s, but got %s", expected, c.String())
	}
}
