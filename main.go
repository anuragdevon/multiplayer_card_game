package main

import (
	"fmt"
	"multiplayer-card-game/game"
)

func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	g := game.NewGame(names)

	for !g.IsGameOver() {
		p := g.CurrentPlayer()
		fmt.Println(len(g.Deck.Cards))
		fmt.Printf(">>>>>>>>%s's turn:\n", p.Name)
		fmt.Printf("^^^^^^Top card^^^^^: %s\n", g.ReturnLastCard())
		for i := 0; i < len(p.Hand); i++ {
			c := p.Hand[i]
			fmt.Printf("%d. %s of %s-------\n", i+1, c.Rank, c.Suit)
		}

		var index int
		fmt.Print("Enter the index of the card you want to play: ")
		fmt.Scanln(&index)

		if err := g.PlayCard(index - 1); err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {

			fmt.Printf("%s played: %s\n", p.Name, p.LastPlayed)
		}
	}

	winner, err := g.Winner()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Game over! %s wins!\n", winner.Name)
	}
}
