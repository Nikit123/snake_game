package snake

import "snake_game/pkg/board/cell"

type Snake struct {
	Head         cell.GameCell
	Body         []cell.GameCell
	Tail         cell.GameCell
	PreviousTail cell.GameCell
}

func CreateMySnake() Snake {
	var snake Snake

	snake.Head.X = 0
	snake.Head.Y = 1
	snake.Head.Status = 1

	snake.Tail.X = 0
	snake.Tail.Y = 0
	snake.Tail.Status = 2

	snake.PreviousTail.X = -1
	snake.PreviousTail.Y = -1

	return snake
}

func (mySnake *Snake) UpdateSnake(targetLocation cell.GameCell, hasEaten bool) {
	mySnake.Body = append(mySnake.Body, mySnake.Head)
	mySnake.Head.X = targetLocation.X
	mySnake.Head.Y = targetLocation.Y
	for i := range mySnake.Body {
		mySnake.Body[i].Status += 1
	}
	if hasEaten {
		mySnake.Tail.Status += 1
	} else {
		mySnake.PreviousTail = mySnake.Tail
		if len(mySnake.Body) > 0 {
			mySnake.Tail = mySnake.Body[0]
			mySnake.Body = mySnake.Body[1:]
		}
	}
}
