# Multiplayer Card Game
This project is an implementation of a multiplayer card game using Go programming language. It supports multiple players, a deck of cards, and basic game logic.

## Requirements
To run the game, you need to have Go installed on your machine. You can download it from the official website: https://golang.org/dl/

## How to Play
To play the game, run the following command from the project directory:

`go run main.go`

OR

To build the game and run:
```
go build
chmod +x multiplayer-card-game
./multiplayer-card-game
```

The game will start and prompt you to enter the number of players (2 to 4). Once you enter the number of players, the game will shuffle the deck and deal the cards to each player.

Each turn, a player has to choose a card to play from their hand. The player can only play a card if it matches the suit or rank of the top card on the discard pile. If the player cannot play any card, they have to draw a card from the deck.

The game continues until one player runs out of cards, declared as Winner or the draw pile is empty, resulting in Game Draw.

## Project Structure
The project is structured as follows:

```
multiplayer-card-game/
├── card/
│   ├── card.go
│   └── card_test.go
├── deck/
│   ├── deck.go
│   └── deck_test.go
├── player/
│   ├── player.go
│   └── player_test.md
├── main.go
└── README.md
```

## Testing
To run the tests for the project, run the following command from the project directory:

go test -v ./...

This command will run all the tests in the project.