package main

import "fmt"

func main() {
	slice := make([]int, 3)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice1 := make([]int, 3, 5)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := make([]int, 3, 5)
	slice2 = append(slice2, 2)
	fmt.Println(slice2)

	slice3 := make([]int, 0)
	slice3 = append(slice3, 2)
	fmt.Println(slice3)

	//No usar el capacity provoca que cree nuevos objetos en memoria es menos eficiente
	//Usar el capacity produce que optimices, ya que reserva en el mismo puntero
}
