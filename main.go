package main

import (
	"snake_game/pkg/actions"
	"snake_game/pkg/board"
	"snake_game/pkg/errorcatch"
	"snake_game/pkg/snake"
)

func main() {

	/*Requesting play area dimensions(length and breadth) from user*/
	l, b := actions.RequestDimensions()

	/*Creating initial empty gameboard*/
	myBoard, err := board.CreateMyGameBoard(l, b)
	if err != nil {
		errorcatch.ReportError(err)
		return
	}

	/*Creating a snake of initial length 2 at top-left corner of grid*/
	mySnake := snake.CreateMySnake()

	/*Prepare final game board with snake and food to start the game*/
	actions.InitializeGame(&myBoard, &mySnake)
	gameWon, gameOver := actions.StartGame(&myBoard, &mySnake)
	actions.FinalizeGame(gameWon, gameOver)
}
