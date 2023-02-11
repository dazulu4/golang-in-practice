package main

import "fmt"

func main() {
	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println(result)

	//Resta
	result = y - x
	fmt.Println(result)

	//Mult
	result = x * y
	fmt.Println(result)

	//Division
	result = y / x
	fmt.Println(result)

	//Modulo
	result = y % x
	fmt.Println(result)

	//Incremental
	x++
	fmt.Println(x)

	//Decremental
	x--
	fmt.Println(x)
}
