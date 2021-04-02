package board

import (
	"fmt"
	"snake_game/board/cell"
)

type GameBoard struct {
	Length  int
	Breadth int
	Data    [][]cell.GameCell
}

func CreateMyGameBoard(l int, b int) GameBoard {
	var g GameBoard
	if l < 1 || b < 2 {
		fmt.Println("Length and Breadth of the play area cannot be less than 1.")
		return g
	}
	g.Length = l
	g.Breadth = b
	grid := make([][]cell.GameCell, l)
	for i := range grid {
		grid[i] = make([]cell.GameCell, b)
	}
	g.Data = grid
	return g
}

func (g GameBoard) Visualize() {
	data := g.Data
	for i := 0; i < g.Length; i++ {
		for j := 0; j < g.Breadth; j++ {
			fmt.Print(data[i][j])
		}
		fmt.Println()
	}
}
