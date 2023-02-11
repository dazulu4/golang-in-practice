package main

import "fmt"

func returnValue(a int) int {
	return a * 2
}

func doubleReturnValue(a int) (b, c int) {
	return a, a * 2
}

func main() {
	fmt.Println("Single Return:", returnValue(2))

	b, c := doubleReturnValue(2)
	fmt.Println("Double Return:", b, c)

	// Para no tomar todos los valores retornados
	d, _ := doubleReturnValue(2)
	fmt.Println("Double Return:", d)
}
