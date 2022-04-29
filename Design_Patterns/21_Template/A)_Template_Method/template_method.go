package main

import "fmt"

/*
	Imagine that we are simulating different kinds of games.
	These games all have very similar structures.
	If you think of games like chess, checkers or card cames, then they all are pretty
	much the same, in the sense that you have a bunch of players and each of the players takes
	their turns.

	We can formalize this turn process using a template method.
	In order to do this we need an interface for what parts of the game we are
	interested in.
*/
type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
}

/*
	Template method which makes use of this game interface.
	Essentially, the template method is a skeleton algorithm.
	We are usign the abstract members of an interface which don't have a
	concrete implementation, until we actually implement the game interface
	in a particular struct.
*/
func PlayGame(g Game) {
	g.Start()
	for !g.HaveWinner() {
		g.TakeTurn()
	}
	fmt.Printf("Player %d wins.\n", g.WinningPlayer())
}

// Structural approach
type chess struct { // Hiding struct and its internals, This will be operated externally by the interface
	turn, maxTurns, currentPlayer int
}

func NewGameOfChess() Game {
	return &chess{0, 10, 0}
}

func (c *chess) Start() {
	fmt.Println("Starting a new game of chess.")
}

func (c *chess) TakeTurn() {
	c.turn++
	fmt.Printf("Turn %d taken by player %d\n", c.turn, c.currentPlayer)
	c.currentPlayer = 1 - c.currentPlayer
}

func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns // This is just a simulation. We have a winner once we reach the maxTurns
}

func (c *chess) WinningPlayer() int {
	return c.currentPlayer
}

func main() {
	chess := NewGameOfChess()
	PlayGame(chess)
}
