package card

type Card struct {
	suit string
	rank string
}

func NewCard(suit, rank string) Card {
	return Card{suit: suit, rank: rank}
}

func (c Card) String() string {
	return c.rank + " of " + c.suit
}
