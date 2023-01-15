package main

import "fmt"

/*
	== igual a
	!= diferente
	&& and
	|| or
*/

func main() {
	if true {
		fmt.Println("Hola mundo")
	}

	x := 10
	y := 10
	if x > y {
		fmt.Printf("%d es mayor que %d\n", x, y)
	} else if x < y {
		fmt.Printf("%d es mayor que %d\n", y, x)
	} else {
		fmt.Printf("Son iguales\n")
	}
}
