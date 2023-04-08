package card

type Card struct {
	Suit string
	Rank string
}

func NewCard(suit, rank string) Card {
	return Card{Suit: suit, Rank: rank}
}
