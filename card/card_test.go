package card

import (
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
