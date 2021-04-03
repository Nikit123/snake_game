package actions

import (
	"errors"
	"fmt"
	"snake_game/pkg/board"
	"snake_game/pkg/errorcatch"
	"snake_game/pkg/food"
	"snake_game/pkg/snake"
)

func RequestDimensions() (int, int) {
	var l int
	var b int
	fmt.Println("Please Enter lenght of the play area:")
	fmt.Scanf("%d", &l)
	fmt.Println("Please Enter breadth of the play area:")
	fmt.Scanf("%d", &b)
	return l, b
}

func RequestAction() (string, error) {
	moveDirection := ""
	fmt.Println("Please choose your move for next round (w/a/s/d)")
	fmt.Scanf("%s", &moveDirection)
	if moveDirection == "w" || moveDirection == "a" || moveDirection == "s" || moveDirection == "d" {
		return moveDirection, nil
	} else {
		return "", errors.New("invalid move")
	}
}

func InitializeGame(myBoard *board.GameBoard, mySnake *snake.Snake) {
	myBoard.UpdateSnakeOnBoard(*mySnake)
	myFood := food.CreateFood(myBoard.Data)
	myBoard.UpdateFoodOnBoard(myFood)
	myBoard.Round = 0
	myBoard.Score = CalculateScore(*mySnake)
	myBoard.Visualize()
}

func CalculateScore(mySnake snake.Snake) int {
	score := mySnake.Tail.Status - 2
	return score
}

func StartGame(myBoard *board.GameBoard, mySnake *snake.Snake) (bool, bool) {
	fmt.Println("-------------------GAME START-------------------")
	gameOver := false
	gameWon := false
	roundNo := 1

	for !gameOver && !gameWon {
		moveDirection, err := RequestAction()
		if err != nil {
			errorcatch.ReportError(err)
			continue
		}
		gameWon, gameOver = ProcessMoveAction(moveDirection, mySnake, myBoard)
		myBoard.Round = roundNo
		myBoard.Score = CalculateScore(*mySnake)
		if !gameOver {
			myBoard.Visualize()
			roundNo++
		}
	}
	return gameWon, gameOver
}

func FinalizeGame(gameWon bool, gameOver bool) {
	if gameWon {
		fmt.Println("------------------- GAME WON -------------------")
	}
	if gameOver {
		fmt.Println("------------------- GAME OVER -------------------")
	}
}
