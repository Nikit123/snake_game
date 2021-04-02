package actions

import (
	"fmt"
)

func RequestDimensions() (int, int) {
	var l int
	var b int
	fmt.Println("------------------GAME START-------------------")
	fmt.Println("Please Enter lenght of the play area:")
	fmt.Scanf("%d", &l)
	fmt.Println("Please Enter breadth of the play area:")
	fmt.Scanf("%d", &b)
	return l, b
}
