package main

import (
	"snake_game/actions"
	"snake_game/board"
	"snake_game/errorcatch"
	"snake_game/snake"
)

func main() {

	l, b := actions.RequestDimensions()

	myBoard, err := board.CreateMyGameBoard(l, b)
	if err != nil {
		errorcatch.ReportError(err)
		return
	}

	mySnake := snake.CreateMySnake()

	actions.InitializeGame(&myBoard, &mySnake)
	gameWon, gameOver := actions.StartGame(&myBoard, &mySnake)
	actions.FinalizeGame(gameWon, gameOver)
}
