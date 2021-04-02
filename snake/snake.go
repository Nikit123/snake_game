package snake

import (
	"snake_game/board/cell"
)

type Snake struct {
	Head cell.GameCell
	Body []cell.GameCell
	Tail cell.GameCell
}

func CreateMySnake() Snake {
	var snake Snake

	snake.Head.X = 0
	snake.Head.Y = 1
	snake.Head.Status = 1

	snake.Tail.X = 0
	snake.Tail.Y = 0
	snake.Tail.Status = 2

	return snake
}
