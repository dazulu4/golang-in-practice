package main

import "fmt"

func main() {
	var arreglo [10]int
	fmt.Println(arreglo)

	arreglo2 := [3]int{1, 2, 3}
	fmt.Println(arreglo2)

	arreglo3 := [3]int{1, 2}
	fmt.Println(arreglo3)

	arreglo3[2] = 20

	for i := 0; i < len(arreglo3); i++ {
		fmt.Println(arreglo3[i])
	}

	var matriz [3][2]int

	matriz[0][1] = 10

	fmt.Println(matriz)

}
