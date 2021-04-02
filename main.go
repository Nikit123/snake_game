package main

import (
	"snake_game/actions"
	"snake_game/board"
	"snake_game/errors"
)

func main() {
	l, b := actions.RequestDimensions()
	myBoard, err := board.CreateMyGameBoard(l, b)
	if err != nil {
		errors.ReportError(err)
		return
	}
	myBoard.Visualize()
}
