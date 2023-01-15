package main

import "fmt"

func main() {
	/*
		1. Es una dirección en memoria
		2. En lugar del valor tenemos al dirección donde esta el valor en posición 0
		3. X,Y apuntan al mismo lugar en memoria => AS123D => 5
		4. X => AS123D => 6
		5. Y ¿? => 6
		*T => *int, *string, *Struct
		Valor zero => nil
	*/

	var x, y *int
	entero := 5

	x = &entero
	y = &entero

	fmt.Println(*x)
	fmt.Println(*y)

	// Modifiquemos el valor en el puntero de x

	*x = 6

	fmt.Println(*x)
	fmt.Println(*y)
}
