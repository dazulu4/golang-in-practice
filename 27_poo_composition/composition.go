package main

import "fmt"

type liquid struct {
	color string
}

func (l *liquid) isColor() {
	l.color = "white"
	fmt.Println(l.color)
}

func main() {
	l := liquid{color: "yellow"}
	l.isColor()
	fmt.Println(l.color)
}
