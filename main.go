// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("-------------------------")
// 	fmt.Println("- Multiplayer Card Game -")
// 	fmt.Println("-------------------------")
// }

package main

import (
	"multiplayer-card-game/deck"
	"multiplayer-card-game/game"
	"multiplayer-card-game/player"
)

func main() {
	// Create deck and shuffle cards
	d := deck.NewDeck()
	d.Shuffle()

	// Create players
	p1 := player.NewPlayer("Player 1")
	p2 := player.NewPlayer("Player 2")
	p3 := player.NewPlayer("Player 3")
	p4 := player.NewPlayer("Player 4")
	players := []*player.Player{p1, p2, p3, p4}

	// Draw initial hands
	d.DrawCards(players)

	// Start game
	g := game.NewGame(players, d)
	g.Play()
}
