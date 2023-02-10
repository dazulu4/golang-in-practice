package main

import "fmt"

func main() {
	matriz := []int{1, 2, 3, 4}
	fmt.Println(matriz)

	var matriz1 []int
	if matriz1 == nil {
		fmt.Println("Esta vacio")
	} else {
		fmt.Println(matriz1)
	}

	matriz2 := []int{1, 2}
	if matriz2 == nil {
		fmt.Println("Esta vacio")
	} else {
		fmt.Println(matriz2)
	}

	// Caracteristicas de un slice:
	// 1. Puntero al arreglo
	// 2. Longitud del arreglo al que apunta
	// 3. Capacidad

	arreglo := [3]int{1, 2, 3}
	fmt.Println(arreglo)

	slice := arreglo[:2]
	fmt.Println(slice) //Slicing
}
