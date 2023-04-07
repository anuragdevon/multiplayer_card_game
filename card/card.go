package main

import "fmt"

type Card struct {
	suit string
	rank string
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.rank, c.suit)
}
