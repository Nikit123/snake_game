package errorcatch

import "fmt"

func ReportError(err error) {
	switch err.Error() {
	case "invalid dimensions":
		fmt.Println("Length and Breadth of play area cannot be less than 2.")
	case "invalid move":
		fmt.Println("Please enter a valid move from w/a/s/d.")
	case "wall collision":
		fmt.Println("Looks like your snake ran into a Wall.")
	case "self collision":
		fmt.Println("Looks like your snake was hungry enough to eat itself.")
	case "backward move":
		fmt.Println("Your snake cannot move backwards.")
	default:
		fmt.Println("Some error occured.")
	}
}
