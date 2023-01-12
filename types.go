package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 22
	edad_str := strconv.Itoa(edad)
	fmt.Println("Mi edad es " + edad_str)

	edad2 := "42"
	edad2_int, _ := strconv.Atoi(edad2)
	fmt.Println(edad2_int + 10)
}
