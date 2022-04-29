package main

import "fmt"

// Now we are going to look how to implement a functional approach.
func PlayGame(start, takeTurn func(), haveWinner func() bool, winningPlayer func() int) {
	start()
	for !haveWinner() {
		takeTurn()
	}

	fmt.Printf("Player %d wins. \n", winningPlayer())
}

func main() {
	turn, maxTurns, currentPlayer := 1, 10, 0

	start := func() {
		fmt.Println("Starting a game of chess")
	}

	takeTurn := func() {
		turn++
		fmt.Printf("Turn %d taken by player %d\n", turn, currentPlayer)
		currentPlayer = (currentPlayer + 1) % 2
	}

	haveWinner := func() bool {
		return turn == maxTurns // This is just a simulation. We have a winner once we reach the maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	PlayGame(start, takeTurn, haveWinner, winningPlayer)
}
