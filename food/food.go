package food

import (
	"math/rand"
	"snake_game/board/cell"
	"time"
)

type Food struct {
	Location cell.GameCell
}

func CreateFood(allCells [][]cell.GameCell) Food {
	var emptyCells []cell.GameCell

	for i := range allCells {
		for j := range allCells[i] {
			if allCells[i][j].Status == 0 {
				emptyCells = append(emptyCells, allCells[i][j])
			}
		}
	}

	var food Food
	rand.Seed(time.Now().Unix())
	food.Location = emptyCells[rand.Intn(len(emptyCells))]
	food.Location.Status = -1
	return food
}
