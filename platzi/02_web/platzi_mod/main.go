package main

import (
	"fmt"

	mc "github.com/dazulu4/mycalculator"
)

func main() {
	input := mc.ReadInput()
	operation := mc.ReadInput()
	c := mc.Calc{}
	fmt.Println(c.Operate(input, operation))
}
