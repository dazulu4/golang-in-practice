package mypackage

import "fmt"

type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

func PrintMessage(message string) {
	fmt.Println(message)
}

func printMessage1(message string) {
	fmt.Println(message)
}
