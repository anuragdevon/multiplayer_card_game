package card

type Card struct {
	Suit string
	Rank string
}

func NewCard(suit, rank string) Card {
	return Card{Suit: suit, Rank: rank}
}

func (c Card) String() string {
	return c.Rank + " of " + c.Suit
}
