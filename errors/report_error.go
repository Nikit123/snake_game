package errors

import "fmt"

func ReportError(err error) {
	fmt.Println(err.Error())
}
