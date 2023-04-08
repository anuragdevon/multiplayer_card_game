package main

import (
	"fmt"
	"multiplayer-card-game/game"
)

func main() {
	fmt.Println("-----------------------Welcome to Multiplayer Card Game!-----------------------")
	fmt.Println("Here are the rules:")
	fmt.Println("- Each player starts with a hand of 5 cards.")
	fmt.Println("- The game starts with a deck of 52 cards (a standard deck of playing cards).")
	fmt.Println("- Players take turns playing cards from their hand, following a set of rules that define what cards can be played when.")
	fmt.Println("- A player can only play a card if it matches either the suit or the rank of the top card on the discard pile.")
	fmt.Println("- If a player cannot play a card, they must draw a card from the draw pile. If the draw pile is empty, the game ends in a draw and no player is declared a winner.")
	fmt.Println("- The game ends when one player runs out of cards who is declared the winner.")
	fmt.Println("- Aces, Kings, Queens and Jack are action cards. When one of these is played the following actions occur:")
	fmt.Println("  - Ace(A): Skip the next player in turn.")
	fmt.Println("  - Kings(K): Reverse the sequence of who plays next.")
	fmt.Println("  - Queens(Q): +2.")
	fmt.Println("  - Jacks(J): +4.")
	fmt.Println("NOTE: actions are not stackable i.e. if Q is played by player 1 then player two draws two cards and cannot play a Q from his hand on that turn even if available.")
	fmt.Println("-----------------------LET'S GO-----------------------")
	var numPlayers int
	for {
		fmt.Print("Enter the number of players (2-4): ")
		fmt.Scanln(&numPlayers)
		if numPlayers >= 2 && numPlayers <= 4 {
			break
		}
		fmt.Println("Invalid number of players. Please enter a number between 2 and 4.")
	}

	names := make([]string, numPlayers)
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Enter name for player %d(leave blank to take default): ", i+1)
		fmt.Scanln(&names[i])
		if names[i] == "" {
			names[i] = fmt.Sprintf("Player%d", i+1)
		}
	}

	g := game.NewGame(names)

	for !g.IsGameOver() {
		p := g.CurrentPlayer()

		fmt.Printf("---- %s's turn ----\n", p.Name)
		fmt.Printf("Top card: %s\n", g.ReturnLastCard())
		fmt.Printf("Your hand:\n")
		for i, c := range p.Hand {
			fmt.Printf("%d. %s of %s\n", i+1, c.Rank, c.Suit)
		}

		var index int
		fmt.Print("Enter the index of the card you want to play: ")
		fmt.Scanln(&index)

		if err := g.PlayCard(index - 1); err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("You played: %s\n", p.LastPlayed)
		}

		fmt.Printf("Cards remaining in deck: %d\n\n", len(g.Deck.Cards))
	}

	winner, err := g.Winner()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Game over! %s wins!\n", winner.Name)
	}
}
