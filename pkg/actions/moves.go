package actions

import (
	"errors"
	"snake_game/pkg/board"
	"snake_game/pkg/errorcatch"
	"snake_game/pkg/food"
	"snake_game/pkg/snake"
)

func ProcessMoveAction(moveDirection string, mySnake *snake.Snake, myBoard *board.GameBoard) (bool, bool) {
	gameWon := false
	gameOver := false
	switch moveDirection {
	case "w":
		gameWon, gameOver = ProcessUpMove(mySnake, myBoard)
	case "a":
		gameWon, gameOver = ProcessLeftMove(mySnake, myBoard)
	case "s":
		gameWon, gameOver = ProcessDownMove(mySnake, myBoard)
	case "d":
		gameWon, gameOver = ProcessRightMove(mySnake, myBoard)
	}
	return gameWon, gameOver
}

func ProcessUpMove(mySnake *snake.Snake, myBoard *board.GameBoard) (bool, bool) {
	if mySnake.Head.X-1 < 0 {
		errorcatch.ReportError(errors.New("wall collision"))
		return false, true
	}
	targetLocation := myBoard.Data[mySnake.Head.X-1][mySnake.Head.Y]
	if targetLocation.Status > 0 {
		if targetLocation.Status == 2 {
			errorcatch.ReportError(errors.New("backward move"))
			return false, true
		} else {
			errorcatch.ReportError(errors.New("self collision"))
			return false, true
		}
	}
	if targetLocation.Status == -1 {
		mySnake.UpdateSnake(targetLocation, true)
	} else {
		mySnake.UpdateSnake(targetLocation, false)
	}
	myBoard.UpdateSnakeOnBoard(*mySnake)
	if mySnake.Tail.Status == (myBoard.Breadth * myBoard.Length) {
		return true, false
	}
	if targetLocation.Status == -1 {
		food := food.CreateFood(myBoard.Data)
		myBoard.UpdateFoodOnBoard(food)
	}
	return false, false
}

func ProcessDownMove(mySnake *snake.Snake, myBoard *board.GameBoard) (bool, bool) {
	if mySnake.Head.X+1 > myBoard.Length-1 {
		errorcatch.ReportError(errors.New("wall collision"))
		return false, true
	}
	targetLocation := myBoard.Data[mySnake.Head.X+1][mySnake.Head.Y]
	if targetLocation.Status > 0 {
		if targetLocation.Status == 2 {
			errorcatch.ReportError(errors.New("backward move"))
			return false, true
		} else {
			errorcatch.ReportError(errors.New("self collision"))
			return false, true
		}
	}
	if targetLocation.Status == -1 {
		mySnake.UpdateSnake(targetLocation, true)
	} else {
		mySnake.UpdateSnake(targetLocation, false)
	}
	myBoard.UpdateSnakeOnBoard(*mySnake)
	if mySnake.Tail.Status == (myBoard.Breadth * myBoard.Length) {
		return true, false
	}
	if targetLocation.Status == -1 {
		food := food.CreateFood(myBoard.Data)
		myBoard.UpdateFoodOnBoard(food)
	}
	return false, false
}

func ProcessLeftMove(mySnake *snake.Snake, myBoard *board.GameBoard) (bool, bool) {
	if mySnake.Head.Y-1 < 0 {
		errorcatch.ReportError(errors.New("wall collision"))
		return false, true
	}
	targetLocation := myBoard.Data[mySnake.Head.X][mySnake.Head.Y-1]
	if targetLocation.Status > 0 {
		if targetLocation.Status == 2 {
			errorcatch.ReportError(errors.New("backward move"))
			return false, true
		} else {
			errorcatch.ReportError(errors.New("self collision"))
			return false, true
		}
	}
	if targetLocation.Status == -1 {
		mySnake.UpdateSnake(targetLocation, true)
	} else {
		mySnake.UpdateSnake(targetLocation, false)
	}
	myBoard.UpdateSnakeOnBoard(*mySnake)
	if mySnake.Tail.Status == (myBoard.Breadth * myBoard.Length) {
		return true, false
	}
	if targetLocation.Status == -1 {
		food := food.CreateFood(myBoard.Data)
		myBoard.UpdateFoodOnBoard(food)
	}
	return false, false
}

func ProcessRightMove(mySnake *snake.Snake, myBoard *board.GameBoard) (bool, bool) {
	if mySnake.Head.Y+1 > myBoard.Breadth-1 {
		errorcatch.ReportError(errors.New("wall collision"))
		return false, true
	}
	targetLocation := myBoard.Data[mySnake.Head.X][mySnake.Head.Y+1]
	if targetLocation.Status > 0 {
		if targetLocation.Status == 2 {
			errorcatch.ReportError(errors.New("backward move"))
			return false, false
		} else {
			errorcatch.ReportError(errors.New("self collision"))
			return false, true
		}
	}
	if targetLocation.Status == -1 {
		mySnake.UpdateSnake(targetLocation, true)
	} else {
		mySnake.UpdateSnake(targetLocation, false)
	}
	myBoard.UpdateSnakeOnBoard(*mySnake)
	if mySnake.Tail.Status == (myBoard.Breadth * myBoard.Length) {
		return true, false
	}
	if targetLocation.Status == -1 {
		food := food.CreateFood(myBoard.Data)
		myBoard.UpdateFoodOnBoard(food)
	}
	return false, false
}
