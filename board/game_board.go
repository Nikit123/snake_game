package board

import (
	"errors"
	"fmt"
	"snake_game/board/cell"
	"snake_game/food"
	"snake_game/snake"
)

type GameBoard struct {
	Length  int
	Breadth int
	Data    [][]cell.GameCell
}

func CreateMyGameBoard(l int, b int) (GameBoard, error) {
	var g GameBoard
	if l < 2 || b < 2 {
		return g, errors.New("length and breadth of the play area cannot be less than 2")
	}
	g.Length = l
	g.Breadth = b
	grid := make([][]cell.GameCell, l)
	for i := range grid {
		grid[i] = make([]cell.GameCell, b)
		for j := range grid[i] {
			grid[i][j].X = i
			grid[i][j].Y = j
			grid[i][j].Status = 0
		}
	}
	g.Data = grid
	g.InitializeGameBoard()
	return g, nil
}

func (g GameBoard) Visualize() {
	data := g.Data
	for i := 0; i < g.Length; i++ {
		for j := 0; j < g.Breadth; j++ {
			fmt.Printf(" %d ", data[i][j].Status)
		}
		fmt.Println()
	}
}

func (g *GameBoard) UpdateSnakeOnBoard(s snake.Snake) {
	g.Data[s.Head.X][s.Head.Y].Status = s.Head.Status
	for i := range s.Body {
		g.Data[s.Body[i].X][s.Body[i].Y].Status = s.Body[i].Status
	}
	g.Data[s.Tail.X][s.Tail.Y].Status = s.Tail.Status
}

func (g *GameBoard) InitializeGameBoard() {
	mySnake := snake.CreateMySnake()
	g.UpdateSnakeOnBoard(mySnake)
	food := food.CreateFood(g.Data)
	g.UpdateFoodOnBoard(food)
}

func (g *GameBoard) UpdateFoodOnBoard(f food.Food) {
	g.Data[f.Location.X][f.Location.Y].Status = f.Location.Status
}
