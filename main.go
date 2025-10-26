package main

import "fmt"

type Game struct {
	board    [3][3]string
	player   string
	gameOver bool
	winner   string
}

func NewGame() *Game {
	game := &Game{
		player:   "X",
		gameOver: false,
		winner:   "",
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game.board[i][j] = " "
		}
	}
	return game
}
func (game *Game) printBoard() {
	fmt.Println("\n  0   1   2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d %s | %s | %s\n", i, game.board[i][0], game.board[i][1], game.board[i][2])
		if i < 2 {
			fmt.Println("  ---------")
		}
	}
	fmt.Println()
}

func main() {
	game := NewGame()
	game.printBoard()
}
