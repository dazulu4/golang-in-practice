package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	copia := make([]int, 4)
	copy(copia, slice)
	fmt.Println(slice)
	fmt.Println(copia)

	copia2 := make([]int, 1)
	copy(copia2, slice)
	fmt.Println(slice)
	fmt.Println(copia2)
	//No copia nada, ya que no tiene la longitud correcta
	//Copy solo copia el mínimo de elementos, según el que tenga menor longitud

	copia3 := make([]int, len(slice))
	copy(copia3, slice)
	fmt.Println(slice)
	fmt.Println(copia3)

	copia4 := make([]int, cap(slice)*2)
	copy(copia4, slice)
	fmt.Println(slice)
	fmt.Println(copia4)

}
