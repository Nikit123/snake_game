package main

import (
	"snake_game/pkg/actions"
	"snake_game/pkg/board"
	"snake_game/pkg/errorcatch"
	"snake_game/pkg/snake"
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
