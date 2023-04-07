// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("-------------------------")
// 	fmt.Println("- Multiplayer Card Game -")
// 	fmt.Println("-------------------------")
// }

// package main

// import (
// 	"multiplayer-card-game/deck"
// 	"multiplayer-card-game/game"
// 	"multiplayer-card-game/player"
// )

// func main() {
// 	// Create deck and shuffle cards
// 	d := deck.NewDeck()
// 	d.Shuffle()

// 	// Create players
// 	p1 := player.NewPlayer("Player 1")
// 	p2 := player.NewPlayer("Player 2")
// 	p3 := player.NewPlayer("Player 3")
// 	p4 := player.NewPlayer("Player 4")
// 	players := []*player.Player{p1, p2, p3, p4}

// 	// Draw initial hands
// 	d.DrawCards(players)

// 	// Start game
// 	g := game.NewGame(players, d)
// 	g.Play()
// }

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
		fmt.Printf("%s's turn:\n", p.Name)
		fmt.Printf("Top card: %s\n", g.TopCard())
		fmt.Printf("Your hand: %v\n", p.Hand)

		var index int
		fmt.Print("Enter the index of the card you want to play: ")
		fmt.Scanln(&index)

		if err := g.PlayCard(index); err != nil {
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
